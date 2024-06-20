<template>
    <div class="not-found">
      <div class="not-found-content">
        <img :src="notFoundImage" alt="404 Not Found" class="not-found-image">
        <h1 class="not-found-title">404 Not Found</h1>
        <p class="not-found-description">抱歉,你访问的页面不存在或已被移动。</p>
        <el-button type="primary" @click="goHome">返回首页</el-button>
      </div>
    </div>
  </template>
  
  <script>
  import { defineComponent } from 'vue'
  import { useRouter } from 'vue-router'
  import notFoundImage from '/src/assets/404.svg'
  
  export default defineComponent({
    setup() {
      const router = useRouter()
  
      const goHome = () => {
        router.push('/')
      }
  
      return {
        notFoundImage,
        goHome
      }
    }
  })
  </script>
  
  <style scoped>
  .not-found {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f5f5f5;
  }
  
  .not-found-content {
    text-align: center;
  }
  
  .not-found-image {
    width: 300px;
    height: auto;
  }
  
  .not-found-title {
    font-size: 2rem;
    margin-top: 1rem;
  }
  
  .not-found-description {
    font-size: 1.2rem;
    margin-top: 0.5rem;
    color: #666;
  }
  </style>