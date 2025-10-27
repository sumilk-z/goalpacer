import React, { useState } from 'react';
import { Card, Button, Form, Input, Space, MessagePlugin, Divider, Tag } from 'tdesign-react';
import { SaveIcon } from 'tdesign-icons-react';

const { FormItem } = Form;

const TimeConfig = () => {
  const [timeRules, setTimeRules] = useState({
    weekday: 120,  // 工作日：2小时
    weekend: 300,  // 周末：5小时
    custom: []     // 自定义日期
  });

  const [customDate, setCustomDate] = useState('');
  const [customMinutes, setCustomMinutes] = useState(null);

  const handleSaveBasic = () => {
    MessagePlugin.success('基础时间规则保存成功');
  };

  const handleAddCustom = () => {
    if (!customDate || !customMinutes) {
      MessagePlugin.warning('请填写完整的自定义日期和时长');
      return;
    }

    const newCustom = {
      date: customDate,
      minutes: customMinutes
    };

    setTimeRules({
      ...timeRules,
      custom: [...timeRules.custom, newCustom]
    });

    setCustomDate('');
    setCustomMinutes(null);
    MessagePlugin.success('自定义时间规则添加成功');
  };

  const handleDeleteCustom = (index) => {
    const newCustom = timeRules.custom.filter((_, i) => i !== index);
    setTimeRules({
      ...timeRules,
      custom: newCustom
    });
    MessagePlugin.success('删除成功');
  };

  const formatMinutes = (minutes) => {
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return mins > 0 ? `${hours}小时${mins}分钟` : `${hours}小时`;
  };

  return (
    <div>
      <div className="page-header">
        <h3>⏰ 时间配置</h3>
        <p>设置您的可用学习时间，帮助系统生成更合理的计划</p>
      </div>

      <Card title="基础时间规则" bordered style={{ marginBottom: 24 }}>
        <Form labelWidth={120}>
          <FormItem label="工作日时长">
            <Space>
              <Input
                type="number"
                value={timeRules.weekday}
                onChange={(value) => setTimeRules({ ...timeRules, weekday: Number(value) })}
                style={{ width: 200 }}
                suffix="分钟"
              />
              <Tag variant="light">{formatMinutes(timeRules.weekday)}</Tag>
            </Space>
            <p style={{ margin: '8px 0 0 0', color: '#999', fontSize: 12 }}>
              周一至周五每天的可用学习时长
            </p>
          </FormItem>
          <FormItem label="周末时长">
            <Space>
              <Input
                type="number"
                value={timeRules.weekend}
                onChange={(value) => setTimeRules({ ...timeRules, weekend: Number(value) })}
                style={{ width: 200 }}
                suffix="分钟"
              />
              <Tag variant="light">{formatMinutes(timeRules.weekend)}</Tag>
            </Space>
            <p style={{ margin: '8px 0 0 0', color: '#999', fontSize: 12 }}>
              周六和周日每天的可用学习时长
            </p>
          </FormItem>
          <FormItem>
            <Button theme="primary" icon={<SaveIcon />} onClick={handleSaveBasic}>
              保存基础规则
            </Button>
          </FormItem>
        </Form>
      </Card>

      <Card title="自定义日期规则" bordered>
        <p style={{ color: '#666', marginBottom: 16 }}>
          为特定日期设置不同的学习时长（例如节假日、考试前等）
        </p>
        
        <Form labelWidth={120}>
          <FormItem label="日期">
            <Input
              type="date"
              value={customDate}
              onChange={setCustomDate}
              style={{ width: 200 }}
            />
          </FormItem>
          <FormItem label="可用时长">
            <Space>
              <Input
                type="number"
                value={customMinutes}
                onChange={(value) => setCustomMinutes(Number(value))}
                style={{ width: 200 }}
                suffix="分钟"
                placeholder="输入分钟数"
              />
              {customMinutes && <Tag variant="light">{formatMinutes(customMinutes)}</Tag>}
            </Space>
          </FormItem>
          <FormItem>
            <Button theme="primary" onClick={handleAddCustom}>
              添加自定义规则
            </Button>
          </FormItem>
        </Form>

        {timeRules.custom.length > 0 && (
          <>
            <Divider />
            <h4 style={{ marginBottom: 16 }}>已设置的自定义规则</h4>
            <Space direction="vertical" style={{ width: '100%' }} size="small">
              {timeRules.custom.map((rule, index) => (
                <Card key={index} size="small" bordered>
                  <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                    <Space>
                      <Tag theme="primary">{rule.date}</Tag>
                      <span>{formatMinutes(rule.minutes)}</span>
                    </Space>
                    <Button
                      theme="danger"
                      variant="text"
                      size="small"
                      onClick={() => handleDeleteCustom(index)}
                    >
                      删除
                    </Button>
                  </div>
                </Card>
              ))}
            </Space>
          </>
        )}

        {timeRules.custom.length === 0 && (
          <div style={{ textAlign: 'center', padding: '40px 0', color: '#999' }}>
            暂无自定义时间规则
          </div>
        )}
      </Card>

      <Card style={{ marginTop: 24, background: '#f6f9ff', border: '1px solid #d9e6ff' }}>
        <h4 style={{ margin: '0 0 12px 0', color: '#1890ff' }}>💡 使用提示</h4>
        <ul style={{ margin: 0, paddingLeft: 20, color: '#666' }}>
          <li>基础规则会应用到所有工作日和周末</li>
          <li>自定义规则的优先级高于基础规则</li>
          <li>系统会根据您设置的时间自动分配学习任务</li>
          <li>建议根据实际情况合理设置时长，避免过度疲劳</li>
        </ul>
      </Card>
    </div>
  );
};

export default TimeConfig;
