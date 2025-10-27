import React, { useState } from 'react';
import { Layout, Menu } from 'tdesign-react';
import { DashboardIcon, FlagIcon, EditIcon, SettingIcon } from 'tdesign-icons-react';
import './App.css';

import TodayPlan from './pages/TodayPlan';
import GoalManagement from './pages/GoalManagement';
import LearningRecords from './pages/LearningRecords';
import SettingsPage from './pages/SettingsPage';

const { Header, Content, Footer, Aside } = Layout;
const { MenuItem, SubMenu } = Menu;

function App() {
  const [activeMenu, setActiveMenu] = useState('1');

  const renderContent = () => {
    if (activeMenu.startsWith('4-')) {
      return <SettingsPage />;
    }
    switch (activeMenu) {
      case '1':
        return <TodayPlan />;
      case '2':
        return <GoalManagement />;
      case '3':
        return <LearningRecords />;
      default:
        return <TodayPlan />;
    }
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Aside style={{ background: '#001529' }}>
        <div className="logo">
          GoalPacer
        </div>
        <Menu
          theme="dark"
          value={activeMenu}
          onChange={(value) => setActiveMenu(value)}
          style={{ background: '#001529' }}
        >
          <MenuItem value="1" icon={<DashboardIcon />}>
            今日计划
          </MenuItem>
          <MenuItem value="2" icon={<FlagIcon />}>
            目标管理
          </MenuItem>
          <MenuItem value="3" icon={<EditIcon />}>
            学习记录
          </MenuItem>
          <SubMenu value="4" title={<span><SettingIcon />设置</span>}>
            <MenuItem value="4-1">
              时间配置
            </MenuItem>
            <MenuItem value="4-2">
              提醒配置
            </MenuItem>
          </SubMenu>
        </Menu>
      </Aside>
      <Layout>
        <Header style={{ background: '#fff', padding: '0 24px', boxShadow: '0 2px 8px rgba(0,0,0,0.1)' }}>
          <h2 style={{ margin: 0 }}>智能学习规划助手</h2>
        </Header>
        <Content style={{ margin: '24px', background: '#f0f2f5' }}>
          <div style={{ padding: 24, background: '#fff', minHeight: 'calc(100vh - 180px)', borderRadius: '8px' }}>
            {renderContent()}
          </div>
        </Content>
        <Footer style={{ textAlign: 'center', background: '#f0f2f5' }}>
          GoalPacer ©2025 - 让学习更高效
        </Footer>
      </Layout>
    </Layout>
  );
}

export default App;
