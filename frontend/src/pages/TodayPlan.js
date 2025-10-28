import React, { useState, useEffect } from 'react';
import { Card, Button, Checkbox, Tag, Space, Divider, Row, Col, MessagePlugin, Loading } from 'tdesign-react';
import { CheckCircleIcon, TimeIcon, RefreshIcon } from 'tdesign-icons-react';
import dayjs from 'dayjs';
import { planAPI } from '../services/api';

const TodayPlan = () => {
  const [plans, setPlans] = useState([]);
  const [loading, setLoading] = useState(false);
  const [refreshing, setRefreshing] = useState(false);
  const [planContent, setPlanContent] = useState(null);

  // 初始化：加载今日计划
  useEffect(() => {
    loadTodayPlan();
  }, []);

  // 加载今日计划
  const loadTodayPlan = async () => {
    setLoading(true);
    try {
      const data = await planAPI.getToday();
      if (data && data.content) {
        setPlanContent(data.content);
        // 尝试解析计划内容
        try {
          const parsed = JSON.parse(data.content);
          if (parsed.tasks && Array.isArray(parsed.tasks)) {
            setPlans(parsed.tasks.map((task, index) => ({
              id: index + 1,
              goalName: task.goal_name || task.title || '学习任务',
              taskContent: task.title || task.description || '',
              learningSteps: task.steps || [task.description || ''],
              estimatedMinutes: task.duration_minutes || 60,
              status: 'pending',
              priority: task.priority === 'high' ? 1 : task.priority === 'medium' ? 2 : 3
            })));
          } else {
            // 如果不是结构化数据，显示原始内容
            setPlans([{
              id: 1,
              goalName: '今日计划',
              taskContent: '查看详细计划',
              learningSteps: [data.content],
              estimatedMinutes: 120,
              status: 'pending',
              priority: 2
            }]);
          }
        } catch (e) {
          // JSON解析失败，显示原始内容
          setPlans([{
            id: 1,
            goalName: '今日计划',
            taskContent: '查看详细计划',
            learningSteps: [data.content],
            estimatedMinutes: 120,
            status: 'pending',
            priority: 2
          }]);
        }
      }
      MessagePlugin.success('计划已加载');
    } catch (error) {
      console.error('加载计划失败:', error);
      MessagePlugin.error('加载计划失败: ' + error.message);
    } finally {
      setLoading(false);
    }
  };

  // 刷新计划（强制重新生成）
  const handleRefreshPlan = async () => {
    setRefreshing(true);
    try {
      const data = await planAPI.refreshToday();
      if (data && data.content) {
        setPlanContent(data.content);
        // 尝试解析计划内容
        try {
          const parsed = JSON.parse(data.content);
          if (parsed.tasks && Array.isArray(parsed.tasks)) {
            setPlans(parsed.tasks.map((task, index) => ({
              id: index + 1,
              goalName: task.goal_name || task.title || '学习任务',
              taskContent: task.title || task.description || '',
              learningSteps: task.steps || [task.description || ''],
              estimatedMinutes: task.duration_minutes || 60,
              status: 'pending',
              priority: task.priority === 'high' ? 1 : task.priority === 'medium' ? 2 : 3
            })));
          } else {
            // 如果不是结构化数据，显示原始内容
            setPlans([{
              id: 1,
              goalName: '今日计划',
              taskContent: '查看详细计划',
              learningSteps: [data.content],
              estimatedMinutes: 120,
              status: 'pending',
              priority: 2
            }]);
          }
        } catch (e) {
          // JSON解析失败，显示原始内容
          setPlans([{
            id: 1,
            goalName: '今日计划',
            taskContent: '查看详细计划',
            learningSteps: [data.content],
            estimatedMinutes: 120,
            status: 'pending',
            priority: 2
          }]);
        }
      }
      MessagePlugin.success('✨ 计划已刷新！');
    } catch (error) {
      console.error('刷新计划失败:', error);
      MessagePlugin.error('刷新计划失败: ' + error.message);
    } finally {
      setRefreshing(false);
    }
  };

  const totalMinutes = plans.reduce((sum, plan) => sum + plan.estimatedMinutes, 0);
  const completedCount = plans.filter(p => p.status === 'completed').length;

  const handleToggleComplete = (id) => {
    setPlans(plans.map(plan => 
      plan.id === id 
        ? { ...plan, status: plan.status === 'completed' ? 'pending' : 'completed' }
        : plan
    ));
    MessagePlugin.success('状态已更新');
  };

  const getPriorityColor = (priority) => {
    const colors = {
      1: 'error',
      2: 'warning',
      3: 'default',
      4: 'success',
      5: 'primary'
    };
    return colors[priority] || 'default';
  };

  if (loading) {
    return (
      <div style={{ textAlign: 'center', padding: '60px 0' }}>
        <Loading />
        <p style={{ marginTop: 16, color: '#999' }}>正在加载计划...</p>
      </div>
    );
  }

  return (
    <div>
      <div className="page-header" style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <div>
          <h3>📅 今日学习计划</h3>
          <p>{dayjs().format('YYYY年MM月DD日 dddd')}</p>
        </div>
        <Button
          theme="primary"
          icon={<RefreshIcon />}
          loading={refreshing}
          onClick={handleRefreshPlan}
          style={{ marginTop: 8 }}
        >
          {refreshing ? '刷新中...' : '🔄 刷新计划'}
        </Button>
      </div>

      <Row gutter={16} style={{ marginBottom: 24 }}>
        <Col span={4}>
          <div className="stat-card" style={{ background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)' }}>
            <h4>{plans.length}</h4>
            <p>今日任务</p>
          </div>
        </Col>
        <Col span={4}>
          <div className="stat-card" style={{ background: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)' }}>
            <h4>{totalMinutes}</h4>
            <p>预计时长(分钟)</p>
          </div>
        </Col>
        <Col span={4}>
          <div className="stat-card" style={{ background: 'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)' }}>
            <h4>{completedCount}/{plans.length}</h4>
            <p>完成进度</p>
          </div>
        </Col>
      </Row>

      <Space direction="vertical" style={{ width: '100%' }} size="large">
        {plans.map((plan) => (
          <Card
            key={plan.id}
            bordered
            hoverShadow
            style={{
              opacity: plan.status === 'completed' ? 0.7 : 1,
              borderLeft: plan.status === 'completed' ? '4px solid #52c41a' : '4px solid #1890ff'
            }}
          >
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' }}>
              <div style={{ flex: 1 }}>
                <Space>
                  <Checkbox
                    checked={plan.status === 'completed'}
                    onChange={() => handleToggleComplete(plan.id)}
                  />
                  <h4 style={{ 
                    margin: 0, 
                    textDecoration: plan.status === 'completed' ? 'line-through' : 'none',
                    color: plan.status === 'completed' ? '#999' : '#000'
                  }}>
                    {plan.taskContent}
                  </h4>
                </Space>
                
                <div style={{ marginTop: 12, marginLeft: 32 }}>
                  <Space>
                    <Tag theme="primary" variant="light">{plan.goalName}</Tag>
                    <Tag theme={getPriorityColor(plan.priority)} variant="light">
                      优先级 {plan.priority}
                    </Tag>
                    <Tag icon={<TimeIcon />} variant="light">
                      {plan.estimatedMinutes} 分钟
                    </Tag>
                  </Space>
                </div>

                <Divider style={{ margin: '16px 0' }} />

                <div style={{ marginLeft: 32 }}>
                  <p style={{ fontWeight: 600, marginBottom: 8, color: '#666' }}>📝 学习步骤：</p>
                  <ol style={{ paddingLeft: 20, margin: 0 }}>
                    {plan.learningSteps.map((step, index) => (
                      <li key={index} style={{ 
                        marginBottom: 8,
                        color: plan.status === 'completed' ? '#999' : '#333'
                      }}>
                        {step}
                      </li>
                    ))}
                  </ol>
                </div>
              </div>

              {plan.status === 'completed' && (
                <CheckCircleIcon size="32px" style={{ color: '#52c41a', marginLeft: 16 }} />
              )}
            </div>
          </Card>
        ))}
      </Space>

      {plans.length === 0 && (
        <Card style={{ textAlign: 'center', padding: '60px 0' }}>
          <p style={{ fontSize: 16, color: '#999' }}>今日暂无学习计划</p>
          <Button theme="primary" style={{ marginTop: 16 }}>生成今日计划</Button>
        </Card>
      )}
    </div>
  );
};

export default TodayPlan;
