import React from 'react';
import { Card, Empty } from 'tdesign-react';
import { InfoCircleIcon } from 'tdesign-icons-react';

const NotificationConfig = () => {
  return (
    <Card>
      <Empty
        description="提醒配置功能正在开发中，敬请期待！"
        icon={<InfoCircleIcon size="64px" />}
      />
    </Card>
  );
};

export default NotificationConfig;
