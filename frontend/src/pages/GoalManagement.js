import React, { useState, useEffect } from 'react';
import { Table, Button, Dialog, Form, Input, Select, Tag, Space, MessagePlugin, Popconfirm, Loading } from 'tdesign-react';
import { AddIcon, EditIcon, DeleteIcon } from 'tdesign-icons-react';
import dayjs from 'dayjs';
import { goalAPI } from '../services/api';

const { FormItem } = Form;

const GoalManagement = () => {
  const [goals, setGoals] = useState([]);
  const [loading, setLoading] = useState(false);
  const [dialogVisible, setDialogVisible] = useState(false);
  const [editingGoal, setEditingGoal] = useState(null);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    status: 'active'
  });

  const statusOptions = [
    { label: '进行中', value: 'active' },
    { label: '已完成', value: 'completed' },
    { label: '已归档', value: 'archived' }
  ];

  // 初始化：加载目标列表
  useEffect(() => {
    loadGoals();
  }, []);

  const loadGoals = async () => {
    setLoading(true);
    try {
      const data = await goalAPI.getAll();
      setGoals(data || []);
    } catch (error) {
      MessagePlugin.error('加载目标失败: ' + error.message);
    } finally {
      setLoading(false);
    }
  };

  const columns = [
    {
      colKey: 'name',
      title: '目标名称',
      width: 200,
      cell: ({ row }) => <strong>{row.name}</strong>
    },
    {
      colKey: 'description',
      title: '描述',
      ellipsis: true
    },
    {
      colKey: 'status',
      title: '状态',
      width: 100,
      cell: ({ row }) => {
        const statusMap = {
          active: { label: '进行中', theme: 'success' },
          completed: { label: '已完成', theme: 'primary' },
          archived: { label: '已归档', theme: 'default' }
        };
        const status = statusMap[row.status];
        return <Tag theme={status.theme}>{status.label}</Tag>;
      }
    },
    {
      colKey: 'created_at',
      title: '创建时间',
      width: 120,
      cell: ({ row }) => row.created_at ? dayjs(row.created_at).format('YYYY-MM-DD') : '-'
    },
    {
      colKey: 'action',
      title: '操作',
      width: 150,
      cell: ({ row }) => (
        <Space>
          <Button
            theme="primary"
            variant="text"
            icon={<EditIcon />}
            onClick={() => handleEdit(row)}
          >
            编辑
          </Button>
          <Popconfirm
            content="确定要删除这个目标吗？"
            onConfirm={() => handleDelete(row.id)}
          >
            <Button
              theme="danger"
              variant="text"
              icon={<DeleteIcon />}
            >
              删除
            </Button>
          </Popconfirm>
        </Space>
      )
    }
  ];

  const handleAdd = () => {
    setEditingGoal(null);
    setFormData({
      name: '',
      description: '',
      status: 'active'
    });
    setDialogVisible(true);
  };

  const handleEdit = (goal) => {
    setEditingGoal(goal);
    setFormData({
      name: goal.name,
      description: goal.description,
      status: goal.status
    });
    setDialogVisible(true);
  };

  const handleDelete = async (id) => {
    try {
      await goalAPI.delete(id);
      MessagePlugin.success('删除成功');
      loadGoals();
    } catch (error) {
      MessagePlugin.error('删除失败: ' + error.message);
    }
  };

  const handleSubmit = async () => {
    if (!formData.name) {
      MessagePlugin.warning('请输入目标名称');
      return;
    }

    try {
      if (editingGoal) {
        await goalAPI.update(editingGoal.id, formData);
        MessagePlugin.success('更新成功');
      } else {
        await goalAPI.create(formData);
        MessagePlugin.success('添加成功');
      }
      setDialogVisible(false);
      loadGoals();
    } catch (error) {
      MessagePlugin.error('操作失败: ' + error.message);
    }
  };

  return (
    <div>
      <div className="page-header">
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <div>
            <h3>🎯 目标管理</h3>
            <p>管理您的学习目标</p>
          </div>
          <Button
            theme="primary"
            icon={<AddIcon />}
            onClick={handleAdd}
          >
            添加目标
          </Button>
        </div>
      </div>

      <Loading loading={loading}>
        <Table
          data={goals}
          columns={columns}
          rowKey="id"
          bordered
          hover
          stripe
        />
      </Loading>

      <Dialog
        header={editingGoal ? '编辑目标' : '添加目标'}
        visible={dialogVisible}
        onClose={() => setDialogVisible(false)}
        onConfirm={handleSubmit}
        width={600}
      >
        <Form labelWidth={100}>
          <FormItem label="目标名称" required>
            <Input
              value={formData.name}
              onChange={(value) => setFormData({ ...formData, name: value })}
              placeholder="例如：算法刷题"
            />
          </FormItem>
          <FormItem label="状态">
            <Select
              value={formData.status}
              onChange={(value) => setFormData({ ...formData, status: value })}
              options={statusOptions}
            />
          </FormItem>
          <FormItem label="描述">
            <Input
              value={formData.description}
              onChange={(value) => setFormData({ ...formData, description: value })}
              placeholder="描述您的学习目标"
              autosize={{ minRows: 3 }}
            />
          </FormItem>
        </Form>
      </Dialog>
    </div>
  );
};

export default GoalManagement;
