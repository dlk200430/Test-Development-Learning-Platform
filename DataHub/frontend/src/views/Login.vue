<template>
  <div class="login-container">
    <canvas id="particles" class="particles-canvas"></canvas>
    
    <div class="gradient-overlay"></div>
    
    <div class="login-wrapper">
      <div class="login-card">
        <div class="login-header">
          <div class="logo-box">
            <svg class="logo-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 11-8 0v4h8z"></path>
            </svg>
          </div>
          <h1 class="login-title">测试开发学习平台</h1>
          <p class="login-subtitle">欢迎回来，请登录您的账户</p>
        </div>

        <form @submit.prevent="handleLogin" class="login-form">
          <div class="form-group">
            <label class="form-label">用户名</label>
            <div class="input-wrapper">
              <svg class="input-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
              <input 
                v-model="form.username" 
                type="text" 
                placeholder="请输入用户名" 
                class="form-input"
                :class="{ 'error': errors.username }"
              />
            </div>
            <span v-if="errors.username" class="error-text">{{ errors.username }}</span>
          </div>

          <div class="form-group">
            <label class="form-label">密码</label>
            <div class="input-wrapper">
              <svg class="input-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
              </svg>
              <input 
                v-model="form.password" 
                :type="showPassword ? 'text' : 'password'" 
                placeholder="请输入密码" 
                class="form-input"
                :class="{ 'error': errors.password }"
              />
              <button 
                type="button" 
                @click="showPassword = !showPassword"
                class="password-toggle"
              >
                <svg v-if="!showPassword" class="toggle-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                </svg>
                <svg v-else class="toggle-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"></path>
                </svg>
              </button>
            </div>
            <span v-if="errors.password" class="error-text">{{ errors.password }}</span>
          </div>

          <div class="form-row">
            <label class="checkbox-label">
              <input 
                v-model="form.rememberMe" 
                type="checkbox" 
                class="checkbox-input"
              />
              <span class="checkbox-custom"></span>
              <span class="checkbox-text">记住我</span>
            </label>
            <a href="#" class="forgot-link">忘记密码？</a>
          </div>

          <button 
            type="submit" 
            :disabled="loading"
            class="login-button"
          >
            <span v-if="loading" class="loading-content">
              <svg class="loading-icon" fill="none" viewBox="0 0 24 24">
                <circle class="loading-circle" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="loading-path" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              登录中...
            </span>
            <span v-else>登录</span>
          </button>
        </form>

        <div class="login-footer">
          <p class="footer-text">
            还没有账户？
            <button @click="goRegister" class="register-link">立即注册</button>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { auth } from '../utils/axios'

const router = useRouter()
const loading = ref(false)
const showPassword = ref(false)

const form = reactive({
  username: '',
  password: '',
  rememberMe: false
})

const errors = reactive({
  username: '',
  password: ''
})

const validateForm = () => {
  errors.username = ''
  errors.password = ''
  
  if (!form.username.trim()) {
    errors.username = '请输入用户名'
    return false
  }
  
  if (!form.password) {
    errors.password = '请输入密码'
    return false
  }
  
  if (form.password.length < 6) {
    errors.password = '密码至少需要6位'
    return false
  }
  
  return true
}

const handleLogin = async () => {
  if (!validateForm()) return
  
  console.log('========================================')
  console.log('🔐 开始登录请求')
  console.log('========================================')
  console.log('📋 请求参数:')
  console.log('   ├─ 用户名:', form.username)
  console.log('   ├─ 密码:', '•'.repeat(form.password.length))
  console.log('   ├─ 记住我:', form.rememberMe)
  console.log('   ├─ 时间戳:', new Date().toLocaleString())
  console.log('   └─ API端点: /api/auth/login')
  
  loading.value = true
  try {
    console.log('\n⏳ 正在发送请求...')
    
    const response = await auth.login({
      username: form.username,
      password: form.password
    })
    
    console.log('\n✅ 登录成功!')
    console.log('========================================')
    console.log('📦 响应数据:')
    console.log('   ├─ 状态: 成功')
    console.log('   ├─ Token:', response.token)
    console.log('   ├─ 用户信息:')
    console.log('   │   ├─ ID:', response.user.id)
    console.log('   │   ├─ 用户名:', response.user.username)
    console.log('   │   └─ 邮箱:', response.user.email)
    console.log('   └─ 响应时间:', new Date().toLocaleString())
    
    console.log('\n💾 正在存储登录信息...')
    localStorage.setItem('token', response.token)
    localStorage.setItem('user', JSON.stringify(response.user))
    
    if (form.rememberMe) {
      localStorage.setItem('rememberMe', 'true')
      localStorage.setItem('savedUsername', form.username)
      console.log('   ├─ 记住我: ✓ (已保存用户名)')
    } else {
      localStorage.removeItem('rememberMe')
      localStorage.removeItem('savedUsername')
      console.log('   ├─ 记住我: ✗')
    }
    
    console.log('   ├─ Token: ✓ 已保存到localStorage')
    console.log('   └─ 用户信息: ✓ 已保存到localStorage')
    
    console.log('\n🔄 准备跳转到仪表盘...')
    console.log('========================================\n')
    
    router.push('/dashboard')
  } catch (error) {
    console.log('\n❌ 登录失败!')
    console.log('========================================')
    console.log('📦 错误信息:')
    console.log('   ├─ 状态码:', error.response?.status || '网络错误')
    console.log('   ├─ 错误消息:', error.response?.data?.error || error.message || '未知错误')
    console.log('   └─ 发生时间:', new Date().toLocaleString())
    console.log('========================================\n')
    
    errors.password = error.response?.data?.error || '登录失败，请检查用户名或密码'
  } finally {
    loading.value = false
  }
}

const goRegister = () => {
  router.push('/register')
}

let particles = []
let animationId = null

const initParticles = () => {
  const canvas = document.getElementById('particles')
  if (!canvas) return
  
  const ctx = canvas.getContext('2d')
  const resizeCanvas = () => {
    canvas.width = window.innerWidth
    canvas.height = window.innerHeight
  }
  resizeCanvas()
  window.addEventListener('resize', resizeCanvas)
  
  class Particle {
    constructor() {
      this.reset()
    }
    
    reset() {
      this.x = Math.random() * canvas.width
      this.y = Math.random() * canvas.height
      this.size = Math.random() * 3 + 1
      this.speedX = (Math.random() - 0.5) * 0.5
      this.speedY = (Math.random() - 0.5) * 0.5
      this.opacity = Math.random() * 0.5 + 0.2
    }
    
    update() {
      this.x += this.speedX
      this.y += this.speedY
      
      if (this.x < 0 || this.x > canvas.width) this.speedX *= -1
      if (this.y < 0 || this.y > canvas.height) this.speedY *= -1
    }
    
    draw() {
      ctx.beginPath()
      ctx.arc(this.x, this.y, this.size, 0, Math.PI * 2)
      ctx.fillStyle = `rgba(139, 92, 246, ${this.opacity})`
      ctx.fill()
    }
  }
  
  for (let i = 0; i < 50; i++) {
    particles.push(new Particle())
  }
  
  const animate = () => {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    
    particles.forEach(p => {
      p.update()
      p.draw()
    })
    
    particles.forEach((p1, i) => {
      particles.slice(i + 1).forEach(p2 => {
        const dx = p1.x - p2.x
        const dy = p1.y - p2.y
        const distance = Math.sqrt(dx * dx + dy * dy)
        
        if (distance < 150) {
          ctx.beginPath()
          ctx.moveTo(p1.x, p1.y)
          ctx.lineTo(p2.x, p2.y)
          ctx.strokeStyle = `rgba(139, 92, 246, ${0.1 * (1 - distance / 150)})`
          ctx.lineWidth = 0.5
          ctx.stroke()
        }
      })
    })
    
    animationId = requestAnimationFrame(animate)
  }
  
  animate()
}

onMounted(() => {
  initParticles()
  
  const savedRememberMe = localStorage.getItem('rememberMe') === 'true'
  const savedUsername = localStorage.getItem('savedUsername')
  
  if (savedRememberMe && savedUsername) {
    form.username = savedUsername
    form.rememberMe = true
  }
})

onUnmounted(() => {
  if (animationId) {
    cancelAnimationFrame(animationId)
  }
})
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  background-color: #0f172a;
  position: relative;
  overflow: hidden;
}

.particles-canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}

.gradient-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, rgba(99, 102, 241, 0.3) 0%, rgba(139, 92, 246, 0.2) 50%, rgba(30, 41, 59, 0.5) 100%);
  z-index: 1;
}

.login-wrapper {
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 20px;
  box-sizing: border-box;
}

.login-card {
  width: 100%;
  max-width: 400px;
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 24px;
  padding: 32px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3), inset 0 1px 0 rgba(255, 255, 255, 0.05);
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.logo-box {
  width: 64px;
  height: 64px;
  margin: 0 auto 16px;
  border-radius: 16px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 4px 20px rgba(99, 102, 241, 0.3);
}

.logo-icon {
  width: 32px;
  height: 32px;
  color: #ffffff;
}

.login-title {
  font-size: 24px;
  font-weight: bold;
  color: #ffffff;
  margin: 0 0 8px;
}

.login-subtitle {
  font-size: 14px;
  color: #94a3b8;
  margin: 0;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  font-size: 14px;
  font-weight: 500;
  color: #cbd5e1;
  margin-bottom: 8px;
}

.input-wrapper {
  position: relative;
}

.input-icon {
  position: absolute;
  left: 16px;
  top: 50%;
  transform: translateY(-50%);
  width: 20px;
  height: 20px;
  color: #94a3b8;
}

.form-input {
  width: 100%;
  padding: 12px 16px 12px 48px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 12px;
  color: #ffffff;
  font-size: 14px;
  outline: none;
  box-sizing: border-box;
  transition: all 0.3s ease;
}

.form-input::placeholder {
  color: rgba(148, 163, 184, 0.6);
}

.form-input:focus {
  border-color: #6366f1;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2);
}

.form-input.error {
  border-color: #ef4444;
}

.form-input.error:focus {
  border-color: #ef4444;
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.2);
}

.error-text {
  font-size: 12px;
  color: #ef4444;
  margin-top: 4px;
}

.password-toggle {
  position: absolute;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
}

.toggle-icon {
  width: 20px;
  height: 20px;
  color: #94a3b8;
  transition: color 0.3s ease;
}

.password-toggle:hover .toggle-icon {
  color: #ffffff;
}

.form-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.checkbox-label {
  display: flex;
  align-items: center;
  cursor: pointer;
  user-select: none;
}

.checkbox-input {
  display: none;
}

.checkbox-custom {
  width: 20px;
  height: 20px;
  border: 1px solid #94a3b8;
  border-radius: 6px;
  background: rgba(255, 255, 255, 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 10px;
  transition: all 0.3s ease;
}

.checkbox-input:checked + .checkbox-custom {
  background: #6366f1;
  border-color: #6366f1;
}

.checkbox-custom svg {
  width: 12px;
  height: 12px;
  color: #ffffff;
}

.checkbox-text {
  font-size: 14px;
  color: #cbd5e1;
  transition: color 0.3s ease;
}

.checkbox-label:hover .checkbox-text {
  color: #ffffff;
}

.checkbox-label:hover .checkbox-custom {
  border-color: #6366f1;
}

.forgot-link {
  font-size: 14px;
  color: #6366f1;
  text-decoration: none;
  transition: color 0.3s ease;
}

.forgot-link:hover {
  color: #818cf8;
}

.login-button {
  width: 100%;
  padding: 12px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  border: none;
  border-radius: 12px;
  color: #ffffff;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 20px rgba(99, 102, 241, 0.3);
}

.login-button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 25px rgba(99, 102, 241, 0.4);
}

.login-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.loading-content {
  display: flex;
  align-items: center;
  justify-content: center;
}

.loading-icon {
  width: 20px;
  height: 20px;
  margin-right: 8px;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.login-footer {
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.footer-text {
  text-align: center;
  font-size: 14px;
  color: #94a3b8;
  margin: 0;
}

.register-link {
  background: none;
  border: none;
  color: #6366f1;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: color 0.3s ease;
}

.register-link:hover {
  color: #818cf8;
}

@media (max-width: 480px) {
  .login-card {
    padding: 24px;
    margin: 0 16px;
  }
  
  .login-title {
    font-size: 20px;
  }
}
</style>