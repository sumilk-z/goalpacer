import React, { useState } from 'react';
import { Card, Button, Checkbox, Tag, Space, Divider, Row, Col, MessagePlugin } from 'tdesign-react';
import { CheckCircleIcon, TimeIcon } from 'tdesign-icons-react';
import dayjs from 'dayjs';

const TodayPlan = () => {
  const [plans, setPlans] = useState([
    {
      id: 1,
      goalName: '算法刷题',
      taskContent: '学习二叉树的遍历算法',
      learningSteps: [
        '阅读《算法导论》第12章二叉树部分',
        '理解前序、中序、后序遍历的区别',
        '在LeetCode上完成3道相关题目',
        '总结遍历算法的时间复杂度'
      ],
      estimatedMinutes: 90,
      status: 'pending',
      priority: 1
    },
    {
      id: 2,
      goalName: 'Golang学习',
      taskContent: '深入理解 channel 的阻塞机制',
      learningSteps: [
        '阅读Go官方文档关于channel的部分',
        '编写示例代码测试有缓冲和无缓冲channel',
        '理解select语句的使用场景',
        '完成一个生产者-消费者模式的实践'
      ],
      estimatedMinutes: 120,
      status: 'pending',
      priority: 2
    },
    {
      id: 3,
      goalName: 'Agent开发',
      taskContent: '学习LangChain框架基础',
      learningSteps: [
        '安装LangChain并配置开发环境',
        '学习Chain和Agent的概念区别',
        '运行官方提供的3个示例代码',
        '尝试构建一个简单的问答Agent'
      ],
      estimatedMinutes: 60,
      status: 'pending',
      priority: 3
    }
  ]);

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

  return (
    <div>
      <div className="page-header">
        <h3>📅 今日学习计划</h3>
        <p>{dayjs().format('YYYY年MM月DD日 dddd')}</p>
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
