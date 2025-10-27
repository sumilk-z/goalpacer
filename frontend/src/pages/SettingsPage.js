import React from 'react';
import { Tabs } from 'tdesign-react';
import TimeConfig from './TimeConfig';
import NotificationConfig from './NotificationConfig';

const { TabPanel } = Tabs;

const SettingsPage = () => {
  return (
    <div>
      <div className="page-header">
        <h3>⚙️ 系统设置</h3>
        <p>管理应用的基础配置和个性化选项</p>
      </div>
      <Tabs defaultValue="time">
        <TabPanel value="time" label="时间配置">
          <div style={{ paddingTop: '24px' }}>
            <TimeConfig />
          </div>
        </TabPanel>
        <TabPanel value="notification" label="提醒配置">
          <div style={{ paddingTop: '24px' }}>
            <NotificationConfig />
          </div>
        </TabPanel>
      </Tabs>
    </div>
  );
};

export default SettingsPage;
