<template>
  <div class="detail-content">
      <div class="detail-header">
        <div>
          <button class="back-btn" @click="goBack">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
            </svg>
            返回
          </button>
          <h1 class="detail-title">报告详情</h1>
        </div>
      </div>

      <div v-if="report" class="report-content">
        <div class="summary-card">
          <div class="summary-header">
            <h2 class="case-name">{{ report.testCase?.name }}</h2>
            <span :class="['status-badge', report.status]">
              {{ report.status === 'passed' ? '✓ 通过' : '✗ 失败' }}
            </span>
          </div>
          <div class="summary-stats">
            <div class="stat">
              <span class="stat-label">执行时间</span>
              <span class="stat-value">{{ formatTime(report.createdAt) }}</span>
            </div>
            <div class="stat">
              <span class="stat-label">响应码</span>
              <span :class="['stat-value', 'code-' + report.responseCode]">{{ report.responseCode }}</span>
            </div>
            <div class="stat">
              <span class="stat-label">耗时</span>
              <span class="stat-value">{{ report.duration }} ms</span>
            </div>
          </div>
        </div>

        <div class="info-section">
          <h3 class="section-title">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
            </svg>
            请求信息
          </h3>
          <div class="info-grid">
            <div class="info-item">
              <span class="info-label">请求方法</span>
              <span :class="['method-badge', report.testCase?.method]">{{ report.testCase?.method }}</span>
            </div>
            <div class="info-item">
              <span class="info-label">请求URL</span>
              <span class="info-value url">{{ report.testCase?.url }}</span>
            </div>
          </div>
          <div v-if="report.testCase?.headers" class="json-section">
            <span class="json-label">请求头</span>
            <pre class="json-content">{{ formatJson(report.testCase.headers) }}</pre>
          </div>
          <div v-if="report.testCase?.params" class="json-section">
            <span class="json-label">URL参数</span>
            <pre class="json-content">{{ formatJson(report.testCase.params) }}</pre>
          </div>
          <div v-if="report.testCase?.body" class="json-section">
            <span class="json-label">请求体</span>
            <pre class="json-content">{{ formatJson(report.testCase.body) }}</pre>
          </div>
        </div>

        <div class="info-section">
          <h3 class="section-title">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
            响应信息
          </h3>
          <div v-if="report.responseBody" class="json-section">
            <span class="json-label">响应体</span>
            <pre class="json-content">{{ formatJson(report.responseBody) }}</pre>
          </div>
          <div v-else class="empty-msg">暂无响应数据</div>
        </div>

        <div v-if="report.errorMessage" class="error-section">
          <h3 class="section-title error">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            错误信息
          </h3>
          <pre class="error-content">{{ report.errorMessage }}</pre>
        </div>
      </div>

      <div v-else class="loading-state">
        <div class="loading-spinner"></div>
        <p>加载中...</p>
      </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { reports as reportsAPI } from '../utils/axios'

const router = useRouter()
const route = useRoute()
const report = ref(null)

onMounted(async () => {
  await loadReport()
})

const loadReport = async () => {
  try {
    report.value = await reportsAPI.getById(route.params.id)
  } catch (error) {
    alert('加载报告失败')
    router.push('/reports')
  }
}

const formatTime = (timeStr) => {
  if (!timeStr) return '-'
  return new Date(timeStr).toLocaleString('zh-CN')
}

const formatJson = (jsonStr) => {
  try {
    return JSON.stringify(JSON.parse(jsonStr), null, 2)
  } catch {
    return jsonStr
  }
}

const goBack = () => {
  router.push('/reports')
}
</script>

<style scoped>
.detail-content {
  max-width: 1400px;
}

.detail-header {
  margin-bottom: 24px;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background: none;
  border: none;
  color: #94a3b8;
  font-size: 14px;
  cursor: pointer;
  margin-bottom: 12px;
  transition: color 0.2s;
}

.back-btn svg {
  width: 18px;
  height: 18px;
}

.back-btn:hover {
  color: #ffffff;
}

.detail-title {
  font-size: 28px;
  font-weight: bold;
  color: #ffffff;
  margin: 0;
}

.report-content {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.summary-card {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 24px;
}

.summary-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 20px;
}

.case-name {
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  margin: 0;
}

.status-badge {
  padding: 6px 14px;
  border-radius: 20px;
  font-size: 13px;
  font-weight: 600;
}

.status-badge.passed {
  background: rgba(34, 197, 94, 0.2);
  color: #22c55e;
}

.status-badge.failed {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.summary-stats {
  display: flex;
  gap: 32px;
}

.stat {
  display: flex;
  flex-direction: column;
}

.stat-label {
  font-size: 12px;
  color: #64748b;
  margin-bottom: 4px;
}

.stat-value {
  font-size: 16px;
  font-weight: 600;
  color: #ffffff;
}

.stat-value[class*="code-20"] {
  color: #22c55e;
}

.stat-value[class*="code-40"] {
  color: #f59e0b;
}

.stat-value[class*="code-50"] {
  color: #ef4444;
}

.info-section {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 24px;
}

.section-title {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 16px;
  font-weight: 600;
  color: #ffffff;
  margin: 0 0 20px;
}

.section-title svg {
  width: 20px;
  height: 20px;
  color: #6366f1;
}

.section-title.error svg {
  color: #ef4444;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  margin-bottom: 20px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.info-label {
  font-size: 12px;
  color: #64748b;
}

.info-value {
  font-size: 14px;
  color: #cbd5e1;
}

.info-value.url {
  word-break: break-all;
}

.method-badge {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 600;
  width: fit-content;
}

.method-badge.GET {
  background: rgba(34, 197, 94, 0.2);
  color: #22c55e;
}

.method-badge.POST {
  background: rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}

.method-badge.PUT {
  background: rgba(245, 158, 11, 0.2);
  color: #f59e0b;
}

.method-badge.DELETE {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.json-section {
  margin-bottom: 16px;
}

.json-section:last-child {
  margin-bottom: 0;
}

.json-label {
  font-size: 12px;
  color: #64748b;
  margin-bottom: 8px;
  display: block;
}

.json-content {
  background: rgba(0, 0, 0, 0.2);
  border-radius: 10px;
  padding: 16px;
  font-family: 'Consolas', 'Monaco', monospace;
  font-size: 13px;
  color: #cbd5e1;
  white-space: pre-wrap;
  word-break: break-all;
  overflow-x: auto;
}

.error-section {
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.2);
  border-radius: 16px;
  padding: 24px;
}

.error-content {
  font-family: 'Consolas', 'Monaco', monospace;
  font-size: 14px;
  color: #fca5a5;
  white-space: pre-wrap;
}

.empty-msg {
  color: #64748b;
  font-size: 14px;
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 100px 20px;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid rgba(255, 255, 255, 0.1);
  border-top-color: #6366f1;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.loading-state p {
  color: #94a3b8;
}

@media (max-width: 768px) {
  .summary-stats {
    flex-direction: column;
    gap: 12px;
  }
  
  .info-grid {
    grid-template-columns: 1fr;
  }
}
</style>