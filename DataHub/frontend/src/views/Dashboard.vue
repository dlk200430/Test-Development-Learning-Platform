<template>
  <div class="dashboard-container">
    <!-- 动态背景 -->
    <div class="bg-grid"></div>
    
    <div class="dashboard-content">
      <!-- 标题区域 -->
      <div class="header">
        <h1>DataHub</h1>
        <p>测试开发平台 · 实时监控系统</p>
      </div>

      <!-- 指标卡片 -->
      <div class="metrics-grid">
        <div class="card metric-card">
          <div class="metric-icon blue">📋</div>
          <div class="metric-info">
            <div class="metric-label">测试用例数</div>
            <div class="metric-value" data-target="3">0</div>
            <div class="metric-change">↑ 2 个较上周</div>
          </div>
        </div>

        <div class="card metric-card">
          <div class="metric-icon purple">▶️</div>
          <div class="metric-info">
            <div class="metric-label">执行次数</div>
            <div class="metric-value" data-target="5">0</div>
            <div class="metric-change">↑ 3 次较上周</div>
          </div>
        </div>

        <div class="card metric-card">
          <div class="metric-icon green">✅</div>
          <div class="metric-info">
            <div class="metric-label">通过率</div>
            <div class="metric-value" data-target="80" data-suffix="%">0</div>
            <div class="metric-change">↑ 5% 较上周</div>
          </div>
        </div>
      </div>

      <!-- 图表卡片 -->
      <div class="card chart-card">
        <div class="card-header">
          <h3>📈 测试趋势</h3>
          <div class="chart-tabs">
            <button class="chart-tab active">周</button>
            <button class="chart-tab">月</button>
            <button class="chart-tab">年</button>
          </div>
        </div>
        <div class="chart-container">
          <div class="chart-bars">
            <div class="chart-bar" data-value="1" style="--height: 40%;"></div>
            <div class="chart-bar" data-value="2" style="--height: 60%;"></div>
            <div class="chart-bar" data-value="1" style="--height: 40%;"></div>
            <div class="chart-bar" data-value="3" style="--height: 80%;"></div>
            <div class="chart-bar" data-value="2" style="--height: 60%;"></div>
            <div class="chart-bar" data-value="1" style="--height: 50%;"></div>
            <div class="chart-bar" data-value="2" style="--height: 70%;"></div>
          </div>
        </div>
        <div class="chart-labels">
          <span class="chart-label">周一</span>
          <span class="chart-label">周二</span>
          <span class="chart-label">周三</span>
          <span class="chart-label">周四</span>
          <span class="chart-label">周五</span>
          <span class="chart-label">周六</span>
          <span class="chart-label">周日</span>
        </div>
      </div>

      <!-- 活动列表 -->
      <div class="card activity-card">
        <h3>🕐 最近执行记录</h3>
        <div class="activity-list">
          <div v-for="report in recentReports" :key="report.id" class="activity-item" @click="viewReport(report.id)">
            <div :class="['activity-icon', report.status === 'passed' ? 'success' : 'alert']">
              {{ report.status === 'passed' ? '✅' : '❌' }}
            </div>
            <div class="activity-content">
              <div class="activity-title">{{ report.testCaseName }}</div>
              <div class="activity-time">{{ report.createTime }} · 耗时 {{ report.duration }}ms</div>
            </div>
            <span :class="['activity-badge', report.status === 'passed' ? 'badge-success' : 'badge-error']">
              {{ report.status === 'passed' ? '通过' : '失败' }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { testcases, reports } from '../utils/axios'

console.log('📊 Dashboard 组件已加载')

const currentUser = computed(() => {
  const user = localStorage.getItem('user')
  return user ? JSON.parse(user).username : '用户'
})

const stats = ref({
  caseCount: 0,
  execCount: 0,
  passRate: 0,
  failCount: 0
})

const recentReports = ref([])

const viewReport = (id) => {
  console.log('🔍 查看报告详情, ID:', id)
  window.location.href = `/reports/${id}`
}

// 数字滚动动画
const animateValue = (element, start, end, duration) => {
  const startTime = performance.now()
  
  const update = (currentTime) => {
    const elapsed = currentTime - startTime
    const progress = Math.min(elapsed / duration, 1)
    const easeOutExpo = 1 - Math.pow(2, -10 * progress)
    const current = start + (end - start) * easeOutExpo
    
    let displayValue
    if (end >= 1000) {
      displayValue = Math.floor(current).toLocaleString()
    } else if (end < 100) {
      displayValue = Math.floor(current)
    } else {
      displayValue = Math.floor(current)
    }
    
    element.textContent = displayValue
    
    if (progress < 1) {
      requestAnimationFrame(update)
    }
  }
  
  requestAnimationFrame(update)
}

onMounted(async () => {
  console.log('📊 Dashboard onMounted 钩子触发')
  try {
    console.log('📊 Dashboard 正在加载数据...')
    
    const [casesRes, reportsRes] = await Promise.all([
      testcases.getAll(),
      reports.getAll()
    ])
    
    console.log('✅ 测试用例数据:', casesRes)
    console.log('✅ 执行报告数据:', reportsRes)
    
    const caseList = casesRes.data
    const reportList = reportsRes.data
    
    console.log('📋 测试用例数量:', caseList.length)
    console.log('📝 执行报告数量:', reportList.length)
    
    stats.value.caseCount = caseList.length
    stats.value.execCount = reportList.length
    
    const passedCount = reportList.filter(r => r.status === 'passed').length
    stats.value.passRate = reportList.length > 0 
      ? Math.round((passedCount / reportList.length) * 100) 
      : 0
    stats.value.failCount = reportList.length - passedCount

    recentReports.value = reportList.slice(0, 5).map(r => ({
      id: r.id,
      testCaseName: r.testCase?.name || '未知',
      status: r.status,
      duration: r.duration,
      createTime: new Date(r.createdAt).toLocaleString('zh-CN')
    }))
    
    // 启动数字动画
    setTimeout(() => {
      const metricValues = document.querySelectorAll('.metric-value')
      metricValues.forEach((el, index) => {
        const target = parseFloat(el.dataset.target) || stats.value.caseCount
        const suffix = el.dataset.suffix || ''
        animateValue(el, 0, target, 1500)
      })
    }, 100)
  } catch (error) {
    console.error('❌ 加载数据失败', error)
  }
})
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.dashboard-container {
  min-height: 100vh;
  background-color: #0a0a0f;
  position: relative;
}

.bg-grid {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: 
    linear-gradient(rgba(0, 212, 255, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(0, 212, 255, 0.03) 1px, transparent 1px);
  background-size: 50px 50px;
  pointer-events: none;
  z-index: 0;
}

.dashboard-content {
  max-width: 1400px;
  margin: 0 auto;
  padding: 40px 20px;
  position: relative;
  z-index: 1;
}

/* 标题区域 */
.header {
  text-align: center;
  margin-bottom: 50px;
  position: relative;
  padding: 30px;
}

.header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 200px;
  height: 3px;
  background: linear-gradient(90deg, transparent, #00d4ff, #a855f7, transparent);
  animation: borderFlow 3s linear infinite;
}

.header::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 300px;
  height: 1px;
  background: linear-gradient(90deg, transparent, #a855f7, #00d4ff, transparent);
}

@keyframes borderFlow {
  0% { background-position: -200px 0; }
  100% { background-position: 200px 0; }
}

.header h1 {
  font-size: 42px;
  font-weight: 700;
  background: linear-gradient(135deg, #00d4ff, #a855f7);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 10px;
  text-shadow: 0 0 30px rgba(0, 212, 255, 0.3);
}

.header p {
  color: #a0a0b0;
  font-size: 16px;
}

/* 指标卡片网格 */
.metrics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
  margin-bottom: 40px;
}

/* 卡片基础样式 */
.card {
  background: rgba(20, 20, 30, 0.6);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 20px;
  padding: 28px;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #00d4ff, #a855f7);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.card:hover {
  transform: translateY(-8px);
  box-shadow: 
    0 20px 40px rgba(0, 0, 0, 0.4),
    0 0 30px rgba(0, 212, 255, 0.1);
}

.card:hover::before {
  opacity: 1;
}

/* 指标卡片 */
.metric-card {
  display: flex;
  align-items: center;
  gap: 20px;
}

.metric-icon {
  width: 64px;
  height: 64px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 28px;
  position: relative;
}

.metric-icon.blue {
  background: rgba(0, 212, 255, 0.15);
  box-shadow: 0 0 20px rgba(0, 212, 255, 0.2);
}

.metric-icon.purple {
  background: rgba(168, 85, 247, 0.15);
  box-shadow: 0 0 20px rgba(168, 85, 247, 0.2);
}

.metric-icon.green {
  background: rgba(34, 197, 94, 0.15);
  box-shadow: 0 0 20px rgba(34, 197, 94, 0.2);
}

.metric-info {
  flex: 1;
}

.metric-label {
  color: #a0a0b0;
  font-size: 14px;
  margin-bottom: 8px;
}

.metric-value {
  font-size: 36px;
  font-weight: 700;
  background: linear-gradient(135deg, #ffffff, #00d4ff);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.metric-change {
  font-size: 13px;
  color: #22c55e;
  margin-top: 4px;
}

/* 图表卡片 */
.chart-card {
  margin-bottom: 40px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.card-header h3 {
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
}

.chart-tabs {
  display: flex;
  gap: 12px;
}

.chart-tab {
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 13px;
  color: #a0a0b0;
  background: transparent;
  border: 1px solid rgba(255, 255, 255, 0.1);
  cursor: pointer;
  transition: all 0.3s ease;
}

.chart-tab.active {
  background: rgba(0, 212, 255, 0.15);
  color: #00d4ff;
  border-color: #00d4ff;
}

/* 柱状图 */
.chart-container {
  height: 280px;
  display: flex;
  align-items: flex-end;
  justify-content: space-around;
  padding: 20px 0;
  position: relative;
}

.chart-bars {
  display: flex;
  align-items: flex-end;
  justify-content: space-around;
  width: 100%;
  height: 220px;
  position: relative;
}

.chart-bar {
  width: 45px;
  background: linear-gradient(180deg, #00d4ff, #a855f7);
  border-radius: 8px 8px 0 0;
  position: relative;
  transition: all 0.3s ease;
  cursor: pointer;
  animation: barGrow 1s ease-out forwards;
  height: var(--height, 50%);
}

@keyframes barGrow {
  from { transform: scaleY(0); transform-origin: bottom; }
  to { transform: scaleY(1); transform-origin: bottom; }
}

.chart-bar:hover {
  filter: brightness(1.2);
  box-shadow: 0 0 20px rgba(0, 212, 255, 0.5);
}

.chart-bar::before {
  content: attr(data-value);
  position: absolute;
  top: -30px;
  left: 50%;
  transform: translateX(-50%);
  font-size: 14px;
  font-weight: 600;
  color: #00d4ff;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.chart-bar:hover::before {
  opacity: 1;
}

.chart-labels {
  display: flex;
  justify-content: space-around;
  margin-top: 12px;
  padding: 0 10px;
}

.chart-label {
  font-size: 12px;
  color: #a0a0b0;
}

/* 活动列表 */
.activity-card h3 {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 20px;
  color: #ffffff;
}

.activity-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.activity-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 16px;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 12px;
  border: 1px solid transparent;
  transition: all 0.3s ease;
  cursor: pointer;
}

.activity-item:hover {
  background: rgba(255, 255, 255, 0.06);
  border-color: rgba(255, 255, 255, 0.1);
}

.activity-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.activity-icon.success { 
  background: rgba(34, 197, 94, 0.15);
}

.activity-icon.alert { 
  background: rgba(239, 68, 68, 0.15);
}

.activity-content {
  flex: 1;
}

.activity-title {
  font-size: 14px;
  font-weight: 500;
  color: #ffffff;
  margin-bottom: 4px;
}

.activity-time {
  font-size: 12px;
  color: #a0a0b0;
}

.activity-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
}

.badge-success {
  background: rgba(34, 197, 94, 0.15);
  color: #22c55e;
}

.badge-error {
  background: rgba(239, 68, 68, 0.15);
  color: #ef4444;
}

/* 响应式 */
@media (max-width: 768px) {
  .dashboard-content {
    padding: 20px 16px;
  }

  .header h1 {
    font-size: 28px;
  }

  .metrics-grid {
    grid-template-columns: 1fr;
  }

  .chart-bars {
    height: 180px;
  }

  .chart-bar {
    width: 30px;
  }

  .chart-labels {
    display: none;
  }
}
</style>