<!--suppress HtmlUnknownTag-->
<template>
  <div class='container'>
    <div class='left-illustration'>
      <img src='/images/illustration.png' class='center-img' alt='illustration'>
    </div>
    <div class='right-container'>
      <img src='/images/logo.png' alt='logo'>
      <div class='login'>
        <el-form ref='loginForm' :model='loginForm' :rules='loginRules'>
          <div v-show='isLoginPage'>
            <div class='form-item-title'>用户名</div>
            <el-form-item prop='username'>
              <el-input v-model='loginForm.username'
                        placeholder='请输入用户名'
                        name='username'/>
            </el-form-item>
            <div class='form-item-title'>密码</div>
            <el-form-item prop='password'>
              <el-input v-model='loginForm.password'
                        type='password'
                        placeholder='请输入密码'
                        name='password'/>
            </el-form-item>
            <div class='button-container'>
              <el-button :loading='loading' type='warning' @click='handleLogin' class='login-button'>
                登录
              </el-button>
              <el-button type='text' @click='changePage' class='register-button'>
                没有账号？前往注册
              </el-button>
            </div>
          </div>
        </el-form>
        <el-form ref='registerForm' :model='registerForm' :rules='registerRules'>
          <div v-show='!isLoginPage'>
            <div class='form-item-title'>用户名</div>
            <el-form-item prop='username'>
              <el-input v-model='registerForm.username'
                        placeholder='请输入用户名'
                        name='username'/>
            </el-form-item>
            <div class='form-item-title'>密码</div>
            <el-form-item prop='password'>
              <el-input v-model='registerForm.password'
                        placeholder='请输入密码'
                        type='password'
                        name='password'/>
            </el-form-item>
            <div class='form-item-title'>确认密码</div>
            <el-form-item prop='password2'>
              <el-input v-model='registerForm.password2'
                        placeholder='请确认密码'
                        type='password'
                        name='password'/>
            </el-form-item>
            <div class='form-item-title'>角色</div>
            <el-form-item prop='userType'>
              <el-radio-group v-model='registerForm.userType' style="width: 100%;">
                <el-radio-button label='养蜂场'>养蜂场</el-radio-button>
                <el-radio-button label='加工厂'>加工厂</el-radio-button>
                <el-radio-button label='批发商'>批发商</el-radio-button>
                <el-radio-button label='零售商'>零售商</el-radio-button>
                <el-radio-button label='消费者'>消费者</el-radio-button>
              </el-radio-group>
            </el-form-item>
            <div class='button-container'>
              <el-button :loading='loading' type='warning' @click='handleRegister' class='login-button'>
                注册
              </el-button>
              <el-button :loading='loading' type='text' @click='changePage' class='register-button'>
                已有账号？前往登录
              </el-button>
            </div>
          </div>
        </el-form>
      </div>
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
  height: 100vh;
  width: 100%;
  position: relative;
  display: flex;
  flex-direction: row;
  background-color: white;
  overflow: hidden;

  .left-illustration {
    flex: 1;
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: hidden;

    .center-img {
      position: absolute;
      left: 8%;
      top: 15%;
      width: 70%;
      height: auto;
      z-index: 1;
    }
  }

  .right-container {
    position: relative;
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;

    img {
      position: absolute;
      top: 15%;
      left: 50%;
      transform: translateX(-50%);
      margin-bottom: 50px;
      width: 50%;
      height: auto;
    }

    .login {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 50%;
      padding: 20px;

      .form-item-title {
        font-size: 24px;
        color: #333;
        font-weight: bold;
        margin-bottom: 12px;
      }

      ::v-deep .el-radio-button {
        margin-right: 6%;

        &.is-active {
          .el-radio-button__inner {
            background-color: #E2A130 !important;
            border-color: #E2A130 !important;
            color: white !important;
            box-shadow: -1px 0 0 0 #E2A130;
          }
        }

        .el-radio-button__inner {
          width: 142%;
          color: #E2A130 !important;
        }
      }

      .button-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 50px;

        .register-button {
          display: block;
          margin-top: 10px;
          font-size: 14px;
          color: #B16400;
        }

        .login-button {
          display: block;
          width: 100%;
          height: 40px;
          font-size: 16px;
        }
      }
    }
  }
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
</style>
