<template>
  <div class="sidebar">
    <div class="sidebar-header">
      <div class="logo-section">
        <div class="logo-icon">🧪</div>
        <div class="logo-text">
          <div class="logo-title">DataHub</div>
          <div class="logo-subtitle">测试开发平台</div>
        </div>
      </div>
    </div>

    <div class="sidebar-menu">
      <button 
        class="menu-item" 
        :class="{ active: currentRoute === 'dashboard' }"
        @click="navigate('/dashboard')"
      >
        <span class="menu-icon">📊</span>
        <span class="menu-text">仪表盘</span>
      </button>

      <button 
        class="menu-item" 
        :class="{ active: currentRoute === 'testcases' || currentRoute === 'createtestcase' || currentRoute === 'edittestcase' }"
        @click="navigate('/testcases')"
      >
        <span class="menu-icon">📋</span>
        <span class="menu-text">测试用例</span>
      </button>

      <button 
        class="menu-item" 
        :class="{ active: currentRoute === 'reports' || currentRoute === 'reportdetail' }"
        @click="navigate('/reports')"
      >
        <span class="menu-icon">📝</span>
        <span class="menu-text">执行报告</span>
      </button>
    </div>

    <div class="sidebar-footer">
      <div class="user-info">
        <div class="user-avatar">{{ user?.username?.charAt(0).toUpperCase() || 'U' }}</div>
        <div class="user-details">
          <div class="username">{{ user?.username || '用户' }}</div>
          <div class="user-role">测试工程师</div>
        </div>
      </div>
      <button class="logout-btn" @click="handleLogout">
        <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path>
        </svg>
        退出
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const user = ref(null)
const currentRoute = computed(() => route.name?.toLowerCase())

onMounted(() => {
  const userStr = localStorage.getItem('user')
  if (userStr) {
    user.value = JSON.parse(userStr)
  }
})

const navigate = (path) => {
  console.log('🧭 Sidebar 导航到:', path)
  router.push(path)
}

const handleLogout = () => {
  console.log('👋 退出登录')
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/login')
}
</script>

<style scoped>
.sidebar {
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  width: 260px;
  background: linear-gradient(180deg, #1e293b 0%, #0f172a 100%);
  border-right: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  flex-direction: column;
  z-index: 1000;
}

.sidebar-header {
  padding: 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-icon {
  font-size: 36px;
}

.logo-text {
  display: flex;
  flex-direction: column;
}

.logo-title {
  font-size: 18px;
  font-weight: bold;
  color: #ffffff;
}

.logo-subtitle {
  font-size: 11px;
  color: #64748b;
  margin-top: 2px;
}

.sidebar-menu {
  flex: 1;
  padding: 20px 12px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 16px;
  background: transparent;
  border: none;
  border-radius: 12px;
  color: #94a3b8;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  text-align: left;
  width: 100%;
}

.menu-item:hover {
  background: rgba(99, 102, 241, 0.15);
  color: #ffffff;
  transform: translateX(4px);
}

.menu-item.active {
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  color: #ffffff;
  box-shadow: 0 4px 15px rgba(99, 102, 241, 0.3);
  transform: translateX(4px);
}

.menu-icon {
  font-size: 20px;
  width: 24px;
  text-align: center;
}

.menu-text {
  font-weight: 600;
}

.sidebar-footer {
  padding: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  font-weight: bold;
  color: #ffffff;
}

.user-details {
  flex: 1;
}

.username {
  font-size: 14px;
  font-weight: 600;
  color: #ffffff;
}

.user-role {
  font-size: 12px;
  color: #64748b;
  margin-top: 2px;
}

.logout-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 12px;
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.2);
  border-radius: 10px;
  color: #ef4444;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.logout-btn svg {
  width: 18px;
  height: 18px;
}

.logout-btn:hover {
  background: rgba(239, 68, 68, 0.2);
  border-color: rgba(239, 68, 68, 0.3);
}

@media (max-width: 768px) {
  .sidebar {
    width: 100%;
    height: auto;
    top: auto;
    bottom: 0;
    left: 0;
    right: 0;
    border-right: none;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
    flex-direction: row;
  }

  .sidebar-header,
  .sidebar-footer {
    display: none;
  }

  .sidebar-menu {
    flex-direction: row;
    padding: 12px;
    width: 100%;
    justify-content: space-around;
    gap: 8px;
  }

  .menu-item {
    flex-direction: column;
    gap: 4px;
    padding: 8px 12px;
    font-size: 11px;
    transform: none !important;
  }

  .menu-item:hover,
  .menu-item.active {
    transform: none !important;
  }

  .menu-icon {
    font-size: 22px;
  }
}
</style>