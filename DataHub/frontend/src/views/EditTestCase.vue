<template>
  <div class="testcase-content">
    <div class="testcase-header">
      <h1 class="testcase-title">编辑测试用例</h1>
      <button class="back-btn" @click="goBack">
        <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
        </svg>
        返回
      </button>
    </div>

    <div class="testcase-form-card">
      <div class="form-group">
        <label class="form-label">用例名称 *</label>
        <input 
          v-model="form.name" 
          type="text" 
          class="form-input" 
          placeholder="请输入用例名称"
        />
      </div>

      <div class="form-group">
        <label class="form-label">描述</label>
        <textarea 
          v-model="form.description" 
          class="form-textarea" 
          placeholder="请输入用例描述"
        ></textarea>
      </div>

      <div class="form-row">
        <div class="form-group">
          <label class="form-label">协议</label>
          <select v-model="form.protocol" class="form-select">
            <option value="http">HTTP</option>
            <option value="https">HTTPS</option>
          </select>
        </div>

        <div class="form-group">
          <label class="form-label">请求方法</label>
          <select v-model="form.method" class="form-select">
            <option value="GET">GET</option>
            <option value="POST">POST</option>
            <option value="PUT">PUT</option>
            <option value="DELETE">DELETE</option>
            <option value="PATCH">PATCH</option>
          </select>
        </div>
      </div>

      <div class="form-group">
        <label class="form-label">URL *</label>
        <input 
          v-model="form.url" 
          type="text" 
          class="form-input" 
          placeholder="请输入请求URL"
        />
      </div>

      <div class="form-group">
        <label class="form-label">请求头 (JSON格式)</label>
        <textarea 
          v-model="form.headers" 
          class="form-textarea" 
          placeholder='{"Content-Type": "application/json"}'
        ></textarea>
      </div>

      <div class="form-group">
        <label class="form-label">URL参数 (JSON格式)</label>
        <textarea 
          v-model="form.params" 
          class="form-textarea" 
          placeholder='{"key": "value"}'
        ></textarea>
      </div>

      <div class="form-group">
        <label class="form-label">请求体 (JSON格式)</label>
        <textarea 
          v-model="form.body" 
          class="form-textarea" 
          placeholder='{"key": "value"}'
        ></textarea>
      </div>

      <div class="form-group">
        <label class="form-label">期望结果</label>
        <textarea 
          v-model="form.expected" 
          class="form-textarea" 
          placeholder="请输入期望结果"
        ></textarea>
      </div>

      <div class="form-actions">
        <button class="submit-btn" @click="handleSubmit" :disabled="loading">
          <span v-if="loading" class="loading-spinner"></span>
          {{ loading ? '更新中...' : '更新' }}
        </button>
        <button class="cancel-btn" @click="goBack">取消</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { testcases } from '../utils/axios'

const router = useRouter()
const route = useRoute()
const loading = ref(false)

const form = ref({
  name: '',
  description: '',
  protocol: 'http',
  method: 'GET',
  url: '',
  headers: '',
  params: '',
  body: '',
  expected: ''
})

onMounted(async () => {
  try {
    const data = await testcases.getById(route.params.id)
    form.value = {
      name: data.name,
      description: data.description || '',
      protocol: data.protocol || 'http',
      method: data.method || 'GET',
      url: data.url,
      headers: data.headers || '',
      params: data.params || '',
      body: data.body || '',
      expected: data.expected || ''
    }
  } catch (error) {
    alert('加载用例失败')
    router.push('/testcases')
  }
})

const handleSubmit = async () => {
  if (!form.value.name || !form.value.url) {
    alert('请填写必填字段（用例名称和URL）')
    return
  }

  loading.value = true
  try {
    await testcases.update(route.params.id, form.value)
    router.push('/testcases')
  } catch (error) {
    alert('更新失败')
  } finally {
    loading.value = false
  }
}

const goBack = () => {
  router.push('/testcases')
}
</script>

<style scoped>
.testcase-content {
  max-width: 800px;
}

.testcase-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.testcase-title {
  font-size: 28px;
  font-weight: bold;
  color: #ffffff;
  margin: 0;
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 16px;
  background: rgba(71, 85, 105, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  color: #94a3b8;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.back-btn:hover {
  background: rgba(71, 85, 105, 0.5);
  color: #ffffff;
}

.back-btn svg {
  width: 16px;
  height: 16px;
}

.testcase-form-card {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 32px;
}

.form-group {
  margin-bottom: 20px;
}

.form-label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: #e2e8f0;
  margin-bottom: 8px;
}

.form-input,
.form-select {
  width: 100%;
  padding: 12px 14px;
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  color: #ffffff;
  font-size: 14px;
  outline: none;
  transition: all 0.2s;
}

.form-input:focus,
.form-select:focus {
  border-color: #6366f1;
  box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
}

.form-input::placeholder {
  color: #64748b;
}

.form-textarea {
  width: 100%;
  padding: 12px 14px;
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  color: #ffffff;
  font-size: 14px;
  outline: none;
  transition: all 0.2s;
  min-height: 100px;
  resize: vertical;
  font-family: monospace;
}

.form-textarea:focus {
  border-color: #6366f1;
  box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.2);
}

.form-textarea::placeholder {
  color: #64748b;
}

.form-row {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
}

.form-actions {
  display: flex;
  gap: 12px;
  margin-top: 28px;
}

.submit-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 14px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  border: none;
  border-radius: 12px;
  color: #ffffff;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.submit-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(99, 102, 241, 0.4);
}

.submit-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loading-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: #ffffff;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.cancel-btn {
  padding: 14px 24px;
  background: rgba(71, 85, 105, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  color: #94a3b8;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.cancel-btn:hover {
  background: rgba(71, 85, 105, 0.5);
  color: #ffffff;
}

@media (max-width: 768px) {
  .form-row {
    grid-template-columns: 1fr;
  }

  .form-actions {
    flex-direction: column;
  }
}
</style>