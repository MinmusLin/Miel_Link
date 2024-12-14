<!--suppress HtmlUnknownTag-->
<template>
  <div>
    <el-form ref='loginForm' :model='loginForm' :rules='loginRules'>
      <div v-show='isLoginPage'>
        <el-form-item prop='username'>
          <el-input ref='username'
                    v-model='loginForm.username'
                    placeholder='请输入用户名'
                    name='username'
                    @keyup.enter.native='handleLogin'/>
        </el-form-item>
        <el-form-item prop='password'>
          <el-input ref='password'
                    v-model='loginForm.password'
                    type='password'
                    placeholder='请输入密码'
                    name='password'
                    @keyup.enter.native='handleLogin'/>
        </el-form-item>
        <el-button type='text' @click='handleRegister'>
          注册
        </el-button>
        <el-button :loading='loading' type='primary' @click.native.prevent='handleLogin'>
          登录
        </el-button>
      </div>
    </el-form>
    <el-form ref='registerForm' :model='registerForm' :rules='registerRules'>
      <div v-show='!isLoginPage'>
        <el-form-item prop='username'>
          <el-input v-model='registerForm.username' placeholder='请输入用户名' name='username'/>
        </el-form-item>
        <el-form-item prop='password'>
          <el-input ref='password' v-model='registerForm.password' placeholder='请输入密码' name='password'/>
        </el-form-item>
        <el-form-item prop='password2'>
          <el-input v-model='registerForm.password2' placeholder='请再次输入密码' name='password'/>
        </el-form-item>
        <el-form-item>
          <el-select v-model='registerForm.userType' placeholder='请选择角色'>
            <el-option v-for='item in options' :key='item.value' :label='item.label' :value='item.value'/>
          </el-select>
        </el-form-item>
        <el-button :loading='loading' type='text' @click='handleRegister'>
          登录
        </el-button>
        <el-button :loading='loading' type='primary' @click.native.prevent='submitRegister'>
          注册
        </el-button>
      </div>
    </el-form>
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
          required: true
        }],
        password: [{
          required: true
        }]
      },
      loading: false,
      redirect: undefined,
      isLoginPage: true,
      registerForm: {
        username: '',
        password: '',
        password2: '',
        userType: ''
      },
      registerRules: {
        username: [{
          required: true
        }],
        password: [{
          required: true
        }],
        password2: [{
          required: true
        }],
        userType: [{
          required: true
        }]
      },
      options: [{
        value: '种植户',
        label: '种植户'
      }, {
        value: '工厂',
        label: '工厂'
      }, {
        value: '运输司机',
        label: '运输司机'
      }, {
        value: '商店',
        label: '商店'
      }, {
        value: '消费者',
        label: '消费者'
      }]
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
      this.loading = true
      this.$store.dispatch('user/login', this.loginForm).then(() => {
        this.$router.push({path: this.redirect || '/'})
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    handleRegister() {
      this.isLoginPage = !this.isLoginPage
    },
    submitRegister() {
      if (this.registerForm.password !== this.registerForm.password2) {
        this.$message.error('两次密码不一致')
        return
      }
      const loading = this.$loading({
        lock: true,
        text: '注册中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      this.$store.dispatch('user/register', this.registerForm).then(() => {
        this.$router.push({path: this.redirect || '/'})
        this.loading = false
        loading.close()
        this.handleRegister()
      }).catch(() => {
        loading.close()
      })
    }
  }
}
</script>

<style lang='scss'>

</style>
