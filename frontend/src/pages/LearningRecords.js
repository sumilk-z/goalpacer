import React, { useState, useEffect, useMemo } from 'react';
import {
  Card, Button, Dialog, Form, Input, Select, Space, MessagePlugin,
  Row, Col, Pagination, Radio, List, Loading
} from 'tdesign-react';
import { AddIcon, FilterIcon, SwapIcon } from 'tdesign-icons-react';
import dayjs from 'dayjs';
import { logAPI, goalAPI } from '../services/api';

const { FormItem } = Form;
const { ListItem } = List;

const LearningRecords = () => {
  const [records, setRecords] = useState([]);
  const [goals, setGoals] = useState([]);
  const [loading, setLoading] = useState(false);
  const [dialogVisible, setDialogVisible] = useState(false);
  const [formData, setFormData] = useState({
    goal_id: null,
    content: '',
    duration: null,
    record_date: dayjs().format('YYYY-MM-DD')
  });

  // 筛选、排序和分页的状态
  const [filterGoal, setFilterGoal] = useState('all');
  const [sortOrder, setSortOrder] = useState('desc');
  const [currentPage, setCurrentPage] = useState(1);
  const pageSize = 3;

  // 初始化：加载数据
  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    setLoading(true);
    try {
      const [recordsData, goalsData] = await Promise.all([
        logAPI.getAll(),
        goalAPI.getAll()
      ]);
      setRecords(recordsData || []);
      setGoals(goalsData || []);
    } catch (error) {
      MessagePlugin.error('加载数据失败: ' + error.message);
    } finally {
      setLoading(false);
    }
  };

  const handleAdd = () => {
    setFormData({
      goal_id: null,
      content: '',
      duration: null,
      record_date: dayjs().format('YYYY-MM-DD')
    });
    setDialogVisible(true);
  };

  const handleSubmit = async () => {
    if (!formData.goal_id || !formData.content) {
      MessagePlugin.warning('请填写完整信息');
      return;
    }

    try {
      await logAPI.create(formData);
      MessagePlugin.success('记录添加成功');
      setDialogVisible(false);
      loadData();
    } catch (error) {
      MessagePlugin.error('添加失败: ' + error.message);
    }
  };

  // 使用 useMemo 进行筛选和排序
  const filteredAndSortedRecords = useMemo(() => {
    let result = [...records];

    // 筛选
    if (filterGoal !== 'all') {
      result = result.filter(r => r.goal_id === parseInt(filterGoal));
    }

    // 排序
    result.sort((a, b) => {
      const dateA = new Date(a.record_date);
      const dateB = new Date(b.record_date);
      return sortOrder === 'asc' ? dateA - dateB : dateB - dateA;
    });

    return result;
  }, [records, filterGoal, sortOrder]);

  // 分页数据
  const paginatedRecords = filteredAndSortedRecords.slice(
    (currentPage - 1) * pageSize,
    currentPage * pageSize
  );

  // 统计数据
  const totalRecords = records.length;
  const totalMinutes = records.reduce((sum, r) => sum + (r.duration || 0), 0);
  const todayRecords = records.filter(r => r.record_date === dayjs().format('YYYY-MM-DD')).length;

  // 构建目标选项
  const goalOptions = [
    { label: '📋 所有目标', value: 'all' },
    ...goals.map(g => ({ label: `🎯 ${g.name}`, value: g.id }))
  ];

  return (
    <div>
      <div className="page-header">
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <div>
            <h3>📚 学习记录</h3>
            <p>记录您的每日学习内容</p>
          </div>
          <Button
            theme="primary"
            icon={<AddIcon />}
            onClick={handleAdd}
          >
            添加记录
          </Button>
        </div>
      </div>

      {/* 统计卡片 */}
      <Row gutter={16} style={{ marginBottom: '24px' }}>
        <Col span={8}>
          <Card>
            <div style={{ textAlign: 'center' }}>
              <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#0052d9' }}>
                {totalRecords}
              </div>
              <div style={{ fontSize: '14px', color: '#999', marginTop: '8px' }}>
                总记录数
              </div>
            </div>
          </Card>
        </Col>
        <Col span={8}>
          <Card>
            <div style={{ textAlign: 'center' }}>
              <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#ff7a45' }}>
                {totalMinutes}
              </div>
              <div style={{ fontSize: '14px', color: '#999', marginTop: '8px' }}>
                总学习时长（分钟）
              </div>
            </div>
          </Card>
        </Col>
        <Col span={8}>
          <Card>
            <div style={{ textAlign: 'center' }}>
              <div style={{ fontSize: '24px', fontWeight: 'bold', color: '#00a870' }}>
                {todayRecords}
              </div>
              <div style={{ fontSize: '14px', color: '#999', marginTop: '8px' }}>
                今日记录数
              </div>
            </div>
          </Card>
        </Col>
      </Row>

      {/* 筛选和排序 */}
      <Card style={{ marginBottom: '24px' }}>
        <Space>
          <div style={{
            display: 'flex',
            alignItems: 'center',
            gap: '8px',
            padding: '8px 16px',
            background: '#fff',
            borderRadius: '6px',
            border: '1px solid #e7e7e7'
          }}>
            <FilterIcon style={{ color: '#0052d9' }} />
            <Select
              value={filterGoal}
              onChange={setFilterGoal}
              style={{ width: '180px' }}
              options={goalOptions}
            />
          </div>

          <div style={{
            display: 'flex',
            alignItems: 'center',
            gap: '8px',
            padding: '8px 16px',
            background: '#fff',
            borderRadius: '6px',
            border: '1px solid #e7e7e7'
          }}>
            <SwapIcon style={{ color: '#0052d9' }} />
            <Radio.Group
              value={sortOrder}
              onChange={(val) => setSortOrder(val)}
              variant="default-filled"
            >
              <Radio.Button value="desc">最新优先</Radio.Button>
              <Radio.Button value="asc">最早优先</Radio.Button>
            </Radio.Group>
          </div>
        </Space>
      </Card>

      {/* 记录列表 */}
      <Loading loading={loading}>
        <List>
          {paginatedRecords.map(record => (
            <ListItem key={record.id} style={{ padding: '20px 0', borderBottom: '1px solid #f0f0f0' }}>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', width: '100%', gap: '24px' }}>
                <div style={{ flex: 1 }}>
                  <div style={{ marginBottom: '12px' }}>
                    <span style={{
                      fontSize: '16px',
                      fontWeight: '600',
                      color: '#0052d9',
                      display: 'inline-block',
                      padding: '4px 12px',
                      background: '#f2f3ff',
                      borderRadius: '4px'
                    }}>
                      🎯 {goals.find(g => g.id === record.goal_id)?.name || '未知目标'}
                    </span>
                  </div>
                  <p style={{
                    fontSize: '14px',
                    color: '#333',
                    lineHeight: '1.6',
                    margin: '8px 0'
                  }}>
                    {record.content}
                  </p>
                </div>
                <div style={{
                  display: 'flex',
                  gap: '24px',
                  alignItems: 'center',
                  whiteSpace: 'nowrap'
                }}>
                  <div style={{ textAlign: 'right' }}>
                    <div style={{ fontSize: '12px', color: '#999', marginBottom: '4px' }}>📅 日期</div>
                    <div style={{ fontSize: '16px', fontWeight: '600', color: '#0052d9' }}>
                      {dayjs(record.record_date).format('MM-DD')}
                    </div>
                  </div>
                  {record.duration && (
                    <div style={{ textAlign: 'right' }}>
                      <div style={{ fontSize: '12px', color: '#999', marginBottom: '4px' }}>⏱️ 时长</div>
                      <div style={{ fontSize: '16px', fontWeight: '600', color: '#ff7a45' }}>
                        {record.duration} 分钟
                      </div>
                    </div>
                  )}
                </div>
              </div>
            </ListItem>
          ))}
        </List>
      </Loading>

      {/* 分页 */}
      {filteredAndSortedRecords.length > 0 && (
        <div style={{ marginTop: '24px', textAlign: 'center' }}>
          <Pagination
            current={currentPage}
            pageSize={pageSize}
            total={filteredAndSortedRecords.length}
            onChange={(pageInfo) => setCurrentPage(pageInfo.current)}
          />
        </div>
      )}

      {/* 添加记录对话框 */}
      <Dialog
        header="添加学习记录"
        visible={dialogVisible}
        onClose={() => setDialogVisible(false)}
        onConfirm={handleSubmit}
        width={600}
      >
        <Form labelWidth={100}>
          <FormItem label="学习目标" required>
            <Select
              value={formData.goal_id}
              onChange={(value) => setFormData({ ...formData, goal_id: value })}
              options={goals.map(g => ({ label: g.name, value: g.id }))}
              placeholder="选择学习目标"
            />
          </FormItem>
          <FormItem label="学习日期" required>
            <Input
              type="date"
              value={formData.record_date}
              onChange={(value) => setFormData({ ...formData, record_date: value })}
            />
          </FormItem>
          <FormItem label="学习时长（分钟）">
            <Input
              type="number"
              value={formData.duration}
              onChange={(value) => setFormData({ ...formData, duration: value ? parseInt(value) : null })}
              placeholder="例如：90"
            />
          </FormItem>
          <FormItem label="学习内容" required>
            <Input
              value={formData.content}
              onChange={(value) => setFormData({ ...formData, content: value })}
              placeholder="描述您的学习内容"
              autosize={{ minRows: 4 }}
            />
          </FormItem>
        </Form>
      </Dialog>
    </div>
  );
};

export default LearningRecords;
