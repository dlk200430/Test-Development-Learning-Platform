<template>
  <div class="testcases-content">
      <div class="testcases-header">
        <div>
          <h1 class="testcases-title">测试用例管理</h1>
          <p class="testcases-subtitle">管理和执行HTTP测试用例</p>
        </div>
        <button class="create-btn" @click="goCreate">
          <span class="btn-icon">➕</span>
          新建用例
        </button>
      </div>

      <div class="filter-bar">
        <div class="search-box">
          <svg class="search-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
          <input 
            v-model="searchQuery" 
            type="text" 
            placeholder="搜索用例名称..." 
            class="search-input"
          />
        </div>
      </div>

      <div class="cases-list">
        <div v-for="item in filteredCases" :key="item.id" class="case-card">
          <div class="case-header">
            <div class="case-name">{{ item.name }}</div>
            <span :class="['case-badge', item.status]">
              {{ item.status === 'active' ? '活跃' : '草稿' }}
            </span>
          </div>
          <p class="case-desc">{{ item.description }}</p>
          <div class="case-info">
            <div class="info-item">
              <span class="info-label">方法</span>
              <span :class="['method-badge', item.method]">{{ item.method }}</span>
            </div>
            <div class="info-item">
              <span class="info-label">URL</span>
              <span class="info-value">{{ item.url }}</span>
            </div>
          </div>
          <div class="case-footer">
            <span class="case-time">{{ formatTime(item.createdAt) }}</span>
            <div class="case-actions">
              <button class="action-btn edit-btn" @click="goEdit(item.id)">
                <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                </svg>
              </button>
              <button class="action-btn exec-btn" @click="handleExecute(item.id)">
                <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </button>
              <button class="action-btn delete-btn" @click="handleDelete(item.id)">
                <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>

      <div v-if="filteredCases.length === 0" class="empty-state">
        <svg class="empty-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
        </svg>
        <p class="empty-text">暂无测试用例</p>
        <button class="empty-btn" @click="goCreate">创建第一个用例</button>
      </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { testcases, reports } from '../utils/axios'

const router = useRouter()
const cases = ref([])
const searchQuery = ref('')

const filteredCases = computed(() => {
  return cases.value.filter(item => {
    return item.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  })
})

const formatTime = (dateStr) => {
  return new Date(dateStr).toLocaleString('zh-CN')
}

onMounted(async () => {
  await loadCases()
})

const loadCases = async () => {
  try {
    const response = await testcases.getAll()
    cases.value = response.data.map(c => ({
      ...c,
      status: 'active'
    }))
  } catch (error) {
    console.error('加载用例失败', error)
  }
}

const goCreate = () => {
  router.push('/testcases/create')
}

const goEdit = (id) => {
  router.push(`/testcases/${id}/edit`)
}

const handleExecute = async (id) => {
  try {
    await reports.execute(id)
    alert('执行成功！报告已生成')
  } catch (error) {
    alert('执行失败')
  }
}

const handleDelete = async (id) => {
  if (!confirm('确定删除此用例？')) return
  try {
    await testcases.delete(id)
    await loadCases()
  } catch (error) {
    alert('删除失败')
  }
}
</script>

<style scoped>
.testcases-content {
  max-width: 1400px;
}

.testcases-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.testcases-title {
  font-size: 28px;
  font-weight: bold;
  color: #ffffff;
  margin: 0 0 8px;
}

.testcases-subtitle {
  font-size: 14px;
  color: #94a3b8;
  margin: 0;
}

.create-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  border: none;
  border-radius: 12px;
  padding: 10px 20px;
  color: #ffffff;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.create-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(99, 102, 241, 0.4);
}

.btn-icon {
  font-size: 18px;
}

.filter-bar {
  margin-bottom: 24px;
}

.search-box {
  position: relative;
  max-width: 400px;
}

.search-icon {
  position: absolute;
  left: 14px;
  top: 50%;
  transform: translateY(-50%);
  width: 18px;
  height: 18px;
  color: #94a3b8;
}

.search-input {
  width: 100%;
  padding: 12px 16px 12px 48px;
  background: rgba(30, 41, 59, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  color: #ffffff;
  font-size: 14px;
  outline: none;
  transition: border-color 0.2s;
}

.search-input::placeholder {
  color: #64748b;
}

.search-input:focus {
  border-color: #6366f1;
}

.cases-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 20px;
}

.case-card {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 20px;
  transition: transform 0.2s, box-shadow 0.2s;
}

.case-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.3);
}

.case-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 12px;
}

.case-name {
  font-size: 16px;
  font-weight: 600;
  color: #ffffff;
}

.case-badge {
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.case-badge.active {
  background: rgba(34, 197, 94, 0.2);
  color: #22c55e;
}

.case-badge.draft {
  background: rgba(148, 163, 184, 0.2);
  color: #94a3b8;
}

.case-desc {
  font-size: 13px;
  color: #94a3b8;
  margin: 0 0 16px;
  line-height: 1.5;
}

.case-info {
  margin-bottom: 16px;
}

.info-item {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 8px;
}

.info-item:last-child {
  margin-bottom: 0;
}

.info-label {
  font-size: 12px;
  color: #64748b;
  min-width: 50px;
}

.info-value {
  font-size: 13px;
  color: #cbd5e1;
  word-break: break-all;
}

.method-badge {
  padding: 3px 10px;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 600;
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

.method-badge.PATCH {
  background: rgba(139, 92, 246, 0.2);
  color: #8b5cf6;
}

.case-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 12px;
  border-top: 1px solid rgba(255, 255, 255, 0.05);
}

.case-time {
  font-size: 12px;
  color: #64748b;
}

.case-actions {
  display: flex;
  gap: 8px;
}

.action-btn {
  width: 36px;
  height: 36px;
  border: none;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s;
}

.action-btn svg {
  width: 16px;
  height: 16px;
}

.edit-btn {
  background: rgba(99, 102, 241, 0.2);
  color: #6366f1;
}

.edit-btn:hover {
  background: rgba(99, 102, 241, 0.3);
}

.exec-btn {
  background: rgba(34, 197, 94, 0.2);
  color: #22c55e;
}

.exec-btn:hover {
  background: rgba(34, 197, 94, 0.3);
}

.delete-btn {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.delete-btn:hover {
  background: rgba(239, 68, 68, 0.3);
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
  margin: 0 0 20px;
}

.empty-btn {
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  border: none;
  border-radius: 12px;
  padding: 12px 24px;
  color: #ffffff;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
}

@media (max-width: 768px) {
  .cases-list {
    grid-template-columns: 1fr;
  }
  
  .testcases-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }
}
</style>