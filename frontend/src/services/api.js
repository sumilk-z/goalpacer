// API 服务层 - 与后端通信
const API_BASE_URL = 'http://localhost:8080/api';

// 通用请求方法
const request = async (method, endpoint, data = null) => {
  try {
    const options = {
      method,
      headers: {
        'Content-Type': 'application/json',
      },
    };

    if (data) {
      options.body = JSON.stringify(data);
    }

    const response = await fetch(`${API_BASE_URL}${endpoint}`, options);
    const result = await response.json();

    if (result.code !== 0) {
      throw new Error(result.message || '请求失败');
    }

    return result.data;
  } catch (error) {
    console.error(`API 请求失败: ${method} ${endpoint}`, error);
    throw error;
  }
};

// ========== 目标管理 API ==========
export const goalAPI = {
  // 获取所有目标
  getAll: () => request('GET', '/goals'),

  // 创建目标
  create: (data) => request('POST', '/goals', data),

  // 更新目标
  update: (id, data) => request('PUT', `/goals/${id}`, data),

  // 删除目标
  delete: (id) => request('DELETE', `/goals/${id}`),
};

// ========== 学习记录 API ==========
export const logAPI = {
  // 获取学习记录
  getAll: (params = {}) => {
    const query = new URLSearchParams(params).toString();
    return request('GET', `/logs${query ? '?' + query : ''}`);
  },

  // 创建学习记录
  create: (data) => request('POST', '/logs', data),
};

// ========== 时间规则 API ==========
export const timeRuleAPI = {
  // 获取时间规则
  getAll: () => request('GET', '/time-rules'),

  // 设置时间规则
  set: (data) => request('POST', '/time-rules', data),
};

// ========== 学习计划 API ==========
export const planAPI = {
  // 获取今日计划
  getToday: () => request('GET', '/plan/today'),

  // 刷新今日计划（强制重新生成）
  refreshToday: () => request('POST', '/plan/today/refresh'),

  // 获取指定日期计划
  getByDate: (date) => request('GET', `/plan?date=${date}`),

  // 创建计划
  create: (data) => request('POST', '/plan', data),

  // 更新计划
  update: (id, data) => request('PUT', `/plan/${id}`, data),

  // 删除计划
  delete: (id) => request('DELETE', `/plan/${id}`),
};
