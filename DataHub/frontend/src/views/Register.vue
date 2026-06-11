<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center">
    <div class="bg-white rounded-2xl shadow-2xl p-8 w-full max-w-md">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-800 mb-2">测试开发学习平台</h1>
        <p class="text-gray-500">创建新账户</p>
      </div>

      <n-form ref="formRef" :model="form" label-placement="top">
        <n-form-item label="用户名" path="username">
          <n-input v-model:value="form.username" placeholder="请输入用户名" />
        </n-form-item>

        <n-form-item label="邮箱" path="email">
          <n-input v-model:value="form.email" type="email" placeholder="请输入邮箱" />
        </n-form-item>

        <n-form-item label="密码" path="password">
          <n-input v-model:value="form.password" type="password" placeholder="请输入密码（至少6位）" />
        </n-form-item>

        <n-form-item>
          <n-button type="primary" block @click="handleRegister" :loading="loading">
            注册
          </n-button>
        </n-form-item>
      </n-form>

      <div class="mt-4 text-center">
        <span class="text-gray-500">已有账户？</span>
        <n-button type="text" @click="goLogin">立即登录</n-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { auth } from '../utils/axios'

const router = useRouter()
const formRef = ref(null)
const loading = ref(false)

const form = ref({
  username: '',
  email: '',
  password: ''
})

const handleRegister = async () => {
  if (!form.value.username || !form.value.email || !form.value.password) {
    return
  }

  loading.value = true
  try {
    await auth.register(form.value)
    alert('注册成功，请登录')
    router.push('/login')
  } catch (error) {
    alert(error.response?.data?.error || '注册失败')
  } finally {
    loading.value = false
  }
}

const goLogin = () => {
  router.push('/login')
}
</script>