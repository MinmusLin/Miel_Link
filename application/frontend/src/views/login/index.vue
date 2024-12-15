<!--suppress HtmlUnknownTag-->
<template>
  <div class='container'>
    <div class='content'>
      <el-form ref='loginForm' :model='loginForm' :rules='loginRules'>
        <div v-show='isLoginPage'>
          <el-form-item prop='username'>
            <el-input v-model='loginForm.username'
                      placeholder='请输入用户名'
                      name='username'/>
          </el-form-item>
          <el-form-item prop='password'>
            <el-input v-model='loginForm.password'
                      type='password'
                      placeholder='请输入密码'
                      name='password'/>
          </el-form-item>
          <el-button :loading='loading' type='warning' @click='handleLogin' style='width: 100%'>
            登录
          </el-button>
          <el-button type='text' @click='changePage'>
            没有账号？前往注册
          </el-button>
        </div>
      </el-form>
      <el-form ref='registerForm' :model='registerForm' :rules='registerRules'>
        <div v-show='!isLoginPage'>
          <el-form-item prop='username'>
            <el-input v-model='registerForm.username'
                      placeholder='请输入用户名'
                      name='username'/>
          </el-form-item>
          <el-form-item prop='password'>
            <el-input v-model='registerForm.password'
                      placeholder='请输入密码'
                      type='password'
                      name='password'/>
          </el-form-item>
          <el-form-item prop='password2'>
            <el-input v-model='registerForm.password2'
                      placeholder='请确认密码'
                      type='password'
                      name='password'/>
          </el-form-item>
          <el-form-item prop='userType'>
            <el-radio-group v-model='registerForm.userType'>
              <el-radio-button label='养蜂场'>养蜂场</el-radio-button>
              <el-radio-button label='加工厂'>加工厂</el-radio-button>
              <el-radio-button label='批发商'>批发商</el-radio-button>
              <el-radio-button label='零售商'>零售商</el-radio-button>
              <el-radio-button label='消费者'>消费者</el-radio-button>
            </el-radio-group>
          </el-form-item>
          <el-button :loading='loading' type='warning' @click='handleRegister' style='width: 100%'>
            注册
          </el-button>
          <el-button :loading='loading' type='text' @click='changePage'>
            已有账号？前往登录
          </el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{
          required: true,
          message: '用户名不能为空',
          trigger: 'blur'
        }],
        password: [{
          required: true,
          message: '密码不能为空',
          trigger: 'blur'
        }]
      },
      loading: false,
      redirect: undefined,
      isLoginPage: true,
      registerForm: {
        username: '',
        password: '',
        password2: '',
        userType: '消费者'
      },
      registerRules: {
        username: [{
          required: true,
          message: '用户名不能为空',
          trigger: 'blur'
        }],
        password: [{
          required: true,
          message: '密码不能为空',
          trigger: 'blur'
        }],
        password2: [{
          required: true,
          message: '请确认密码',
          trigger: 'blur'
        }],
        userType: [{
          required: true,
          message: '请选择用户类型',
          trigger: 'change'
        }]
      }
    }
  },
  watch: {
    $route: {
      handler: function (route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    handleLogin() {
      if (!this.loginForm.username || !this.loginForm.password) {
        this.$message.warning('请输入用户名和密码')
        return
      }
      this.loading = true
      this.$store.dispatch('user/login', this.loginForm).then(() => {
        this.$router.push({path: this.redirect || '/home'})
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    handleRegister() {
      if (!this.registerForm.username || !this.registerForm.password || !this.registerForm.password2) {
        this.$message.warning('请输入用户名和密码')
        return
      }
      if (this.registerForm.password !== this.registerForm.password2) {
        this.$message.warning('密码不一致')
        return
      }
      const loading = this.$loading({
        lock: true,
        text: '注册中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      this.$store.dispatch('user/register', this.registerForm).then(() => {
        this.$router.push({path: this.redirect || '/home'})
        this.loading = false
        loading.close()
        this.isLoginPage = !this.isLoginPage
      }).catch(() => {
        loading.close()
      })
    },
    changePage() {
      this.isLoginPage = !this.isLoginPage
      if (this.isLoginPage) {
        this.loginForm.username = ''
        this.loginForm.password = ''
      } else {
        this.registerForm.username = ''
        this.registerForm.password = ''
        this.registerForm.password2 = ''
        this.registerForm.userType = '消费者'
      }
    }
  }
}
</script>

<style scoped lang='scss'>
.container {
  width: 100%;
  height: 100vh;
  position: relative;
  background-color: white;
  overflow: hidden;
}

.container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(100deg, #FFF2E4 50%, transparent 50%);
  z-index: 0;
}

.content {
  position: relative;
  z-index: 1;
  width: 416px;
}
</style>
