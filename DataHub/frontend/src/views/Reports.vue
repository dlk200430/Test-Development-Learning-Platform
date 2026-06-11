<template>
  <div class="reports-content">
      <div class="reports-header">
        <div>
          <h1 class="reports-title">执行报告</h1>
          <p class="reports-subtitle">查看测试执行记录和结果</p>
        </div>
      </div>

      <div class="stats-row">
        <div class="stat-item">
          <div class="stat-num">{{ totalCount }}</div>
          <div class="stat-label">总执行数</div>
        </div>
        <div class="stat-item success">
          <div class="stat-num">{{ passedCount }}</div>
          <div class="stat-label">通过</div>
        </div>
        <div class="stat-item error">
          <div class="stat-num">{{ failedCount }}</div>
          <div class="stat-label">失败</div>
        </div>
      </div>

      <div class="reports-list">
        <div v-for="report in reports" :key="report.id" class="report-card" @click="goDetail(report.id)">
          <div class="report-header">
            <div class="report-name">{{ report.testCaseName }}</div>
            <span :class="['status-tag', report.status]">
              {{ report.status === 'passed' ? '✓ 通过' : '✗ 失败' }}
            </span>
          </div>
          <div class="report-info">
            <div class="info-row">
              <span class="info-label">响应码</span>
              <span :class="['code-badge', report.responseCode]">{{ report.responseCode }}</span>
            </div>
            <div class="info-row">
              <span class="info-label">耗时</span>
              <span class="info-value">{{ report.duration }} ms</span>
            </div>
            <div class="info-row">
              <span class="info-label">执行时间</span>
              <span class="info-value">{{ report.createTime }}</span>
            </div>
          </div>
          <div v-if="report.errorMessage" class="report-error">
            <span class="error-icon">⚠️</span>
            <span class="error-text">{{ report.errorMessage }}</span>
          </div>
        </div>
      </div>

      <div v-if="reports.length === 0" class="empty-state">
        <svg class="empty-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        <p class="empty-text">暂无执行报告</p>
        <p class="empty-hint">执行测试用例后会在这里显示报告</p>
      </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { reports as reportsAPI } from '../utils/axios'

const router = useRouter()
const reports = ref([])

const totalCount = computed(() => reports.value.length)
const passedCount = computed(() => reports.value.filter(r => r.status === 'passed').length)
const failedCount = computed(() => reports.value.filter(r => r.status === 'failed').length)

onMounted(async () => {
  await loadReports()
})

const loadReports = async () => {
  try {
    const response = await reportsAPI.getAll()
    reports.value = response.data.map(r => ({
      id: r.id,
      testCaseName: r.testCase?.name || '未知',
      status: r.status,
      responseCode: r.responseCode,
      duration: r.duration,
      errorMessage: r.errorMessage,
      createTime: new Date(r.createdAt).toLocaleString('zh-CN')
    }))
  } catch (error) {
    console.error('加载报告失败', error)
  }
}

const goDetail = (id) => {
  router.push(`/reports/${id}`)
}
</script>

<style scoped>
.reports-content {
  max-width: 1400px;
}

.reports-header {
  margin-bottom: 24px;
}

.reports-title {
  font-size: 28px;
  font-weight: bold;
  color: #ffffff;
  margin: 0 0 8px;
}

.reports-subtitle {
  font-size: 14px;
  color: #94a3b8;
  margin: 0;
}

.stats-row {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
}

.stat-item {
  flex: 1;
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 20px;
  text-align: center;
}

.stat-item.success {
  border-color: rgba(34, 197, 94, 0.3);
}

.stat-item.error {
  border-color: rgba(239, 68, 68, 0.3);
}

.stat-num {
  font-size: 32px;
  font-weight: bold;
  color: #ffffff;
  margin: 0 0 8px;
}

.stat-item.success .stat-num {
  color: #22c55e;
}

.stat-item.error .stat-num {
  color: #ef4444;
}

.stat-label {
  font-size: 13px;
  color: #94a3b8;
}

.reports-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(360px, 1fr));
  gap: 16px;
}

.report-card {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 20px;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.report-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.3);
  border-color: rgba(99, 102, 241, 0.3);
}

.report-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 16px;
}

.report-name {
  font-size: 16px;
  font-weight: 600;
  color: #ffffff;
  flex: 1;
  margin-right: 12px;
}

.status-tag {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
}

.status-tag.passed {
  background: rgba(34, 197, 94, 0.2);
  color: #22c55e;
}

.status-tag.failed {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.report-info {
  margin-bottom: 12px;
}

.info-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.info-row:last-child {
  margin-bottom: 0;
}

.info-label {
  font-size: 12px;
  color: #64748b;
}

.info-value {
  font-size: 13px;
  color: #cbd5e1;
}

.code-badge {
  padding: 2px 8px;
  border-radius: 6px;
  font-size: 12px;
  font-weight: 600;
}

.code-badge[class*="20"] {
  background: rgba(34, 197, 94, 0.2);
  color: #22c55e;
}

.code-badge[class*="40"] {
  background: rgba(245, 158, 11, 0.2);
  color: #f59e0b;
}

.code-badge[class*="50"] {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.report-error {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px;
  background: rgba(239, 68, 68, 0.1);
  border-radius: 8px;
}

.error-icon {
  font-size: 16px;
}

.error-text {
  font-size: 12px;
  color: #fca5a5;
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
}

.empty-icon {
  width: 64px;
  height: 64px;
  color: #64748b;
  margin-bottom: 16px;
}

.empty-text {
  font-size: 16px;
  color: #94a3b8;
  margin: 0 0 8px;
}

.empty-hint {
  font-size: 14px;
  color: #64748b;
  margin: 0;
}

@media (max-width: 768px) {
  .reports-list {
    grid-template-columns: 1fr;
  }
  
  .stats-row {
    flex-direction: column;
  }
}
</style>