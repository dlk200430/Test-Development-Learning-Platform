import axios from 'axios'

const mockUsers = [
  { id: 1, username: 'testuser', password: '123456', email: 'test@example.com' }
]

const mockToken = 'mock-jwt-token-12345'

const mockTestCases = [
  {
    id: 1,
    name: '测试用户登录接口',
    description: '测试登录API的正常和异常情况',
    protocol: 'HTTP',
    method: 'POST',
    url: 'http://localhost:8080/api/auth/login',
    headers: JSON.stringify({ 'Content-Type': 'application/json' }),
    params: JSON.stringify({}),
    body: JSON.stringify({ username: 'test', password: '123456' }),
    expectedStatus: 200,
    createdAt: '2026-06-01T10:00:00Z',
    updatedAt: '2026-06-01T10:00:00Z'
  },
  {
    id: 2,
    name: '测试获取用户列表',
    description: '测试获取用户列表接口',
    protocol: 'HTTP',
    method: 'GET',
    url: 'http://localhost:8080/api/users',
    headers: JSON.stringify({}),
    params: JSON.stringify({ page: 1, size: 10 }),
    body: JSON.stringify({}),
    expectedStatus: 200,
    createdAt: '2026-06-02T14:00:00Z',
    updatedAt: '2026-06-02T14:00:00Z'
  },
  {
    id: 3,
    name: '测试创建用户',
    description: '测试POST创建新用户',
    protocol: 'HTTP',
    method: 'POST',
    url: 'http://localhost:8080/api/auth/register',
    headers: JSON.stringify({ 'Content-Type': 'application/json' }),
    params: JSON.stringify({}),
    body: JSON.stringify({ username: 'newuser', password: '123456', email: 'new@example.com' }),
    expectedStatus: 201,
    createdAt: '2026-06-03T09:00:00Z',
    updatedAt: '2026-06-03T09:00:00Z'
  }
]

const mockReports = [
  {
    id: 1,
    testCaseId: 1,
    testCase: mockTestCases[0],
    status: 'passed',
    responseCode: 200,
    responseBody: JSON.stringify({ token: 'abc123', user: { id: 1, username: 'testuser' } }),
    duration: 125,
    errorMessage: '',
    createdAt: '2026-06-05T10:30:00Z'
  },
  {
    id: 2,
    testCaseId: 1,
    testCase: mockTestCases[0],
    status: 'failed',
    responseCode: 401,
    responseBody: JSON.stringify({ error: 'Unauthorized' }),
    duration: 89,
    errorMessage: '用户名或密码错误',
    createdAt: '2026-06-05T11:00:00Z'
  },
  {
    id: 3,
    testCaseId: 2,
    testCase: mockTestCases[1],
    status: 'passed',
    responseCode: 200,
    responseBody: JSON.stringify({ data: [], total: 0 }),
    duration: 45,
    errorMessage: '',
    createdAt: '2026-06-06T14:20:00Z'
  },
  {
    id: 4,
    testCaseId: 3,
    testCase: mockTestCases[2],
    status: 'passed',
    responseCode: 201,
    responseBody: JSON.stringify({ message: '注册成功' }),
    duration: 156,
    errorMessage: '',
    createdAt: '2026-06-07T09:15:00Z'
  },
  {
    id: 5,
    testCaseId: 1,
    testCase: mockTestCases[0],
    status: 'passed',
    responseCode: 200,
    responseBody: JSON.stringify({ token: 'def456', user: { id: 1, username: 'testuser' } }),
    duration: 98,
    errorMessage: '',
    createdAt: '2026-06-10T08:00:00Z'
  }
]

const instance = axios.create({
  baseURL: 'http://localhost:8080/api',
  timeout: 10000
})

instance.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  response => {
    return response.data
  },
  error => {
    if (error.response && error.response.status === 401) {
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.href = '/'
    }
    return Promise.reject(error)
  }
)

const auth = {
  login: async (data) => {
    console.log('\n🌐 [Auth Module] 收到登录请求')
    console.log('   ├─ 请求数据:', JSON.stringify(data, null, 2))
    console.log('   ├─ 时间:', new Date().toISOString())
    
    const user = mockUsers.find(u => u.username === data.username && u.password === data.password)
    if (user) {
      console.log('   └─ ✅ 用户验证成功')
      const response = {
        token: mockToken,
        user: {
          id: user.id,
          username: user.username,
          email: user.email
        }
      }
      console.log('   └─ 📤 返回响应:', JSON.stringify(response, null, 2))
      return response
    }
    
    console.log('   └─ ❌ 用户验证失败')
    throw { response: { data: { error: '用户名或密码错误' } } }
  },
  
  register: async (data) => {
    const exists = mockUsers.find(u => u.username === data.username)
    if (exists) {
      throw { response: { data: { error: '用户名已存在' } } }
    }
    const newUser = {
      id: mockUsers.length + 1,
      username: data.username,
      password: data.password,
      email: data.email
    }
    mockUsers.push(newUser)
    return {
      token: mockToken,
      user: newUser
    }
  }
}

const testcases = {
  getAll: async () => {
    return { data: mockTestCases }
  },
  
  getById: async (id) => {
    const tc = mockTestCases.find(t => t.id === parseInt(id))
    if (tc) return tc
    throw { response: { data: { error: '用例不存在' } } }
  },
  
  create: async (data) => {
    const newCase = {
      id: mockTestCases.length + 1,
      ...data,
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    }
    mockTestCases.push(newCase)
    return newCase
  },
  
  update: async (id, data) => {
    const index = mockTestCases.findIndex(t => t.id === parseInt(id))
    if (index === -1) {
      throw { response: { data: { error: '用例不存在' } } }
    }
    mockTestCases[index] = {
      ...mockTestCases[index],
      ...data,
      updatedAt: new Date().toISOString()
    }
    return mockTestCases[index]
  },
  
  delete: async (id) => {
    const index = mockTestCases.findIndex(t => t.id === parseInt(id))
    if (index === -1) {
      throw { response: { data: { error: '用例不存在' } } }
    }
    mockTestCases.splice(index, 1)
    return { message: '删除成功' }
  }
}

const reports = {
  getAll: async () => {
    return { data: mockReports }
  },
  
  getById: async (id) => {
    const report = mockReports.find(r => r.id === parseInt(id))
    if (report) return report
    throw { response: { data: { error: '报告不存在' } } }
  },
  
  execute: async (testCaseId) => {
    const tc = mockTestCases.find(t => t.id === parseInt(testCaseId))
    if (!tc) {
      throw { response: { data: { error: '用例不存在' } } }
    }
    
    const isSuccess = Math.random() > 0.3
    const duration = Math.floor(Math.random() * 200) + 50
    
    const newReport = {
      id: mockReports.length + 1,
      testCaseId: tc.id,
      testCase: tc,
      status: isSuccess ? 'passed' : 'failed',
      responseCode: isSuccess ? 200 : 400,
      responseBody: isSuccess 
        ? JSON.stringify({ success: true, data: {} })
        : JSON.stringify({ error: '测试失败' }),
      duration: duration,
      errorMessage: isSuccess ? '' : '模拟测试失败',
      createdAt: new Date().toISOString()
    }
    
    mockReports.unshift(newReport)
    return newReport
  }
}

export { auth, testcases, reports }
export default instance