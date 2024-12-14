<!--suppress HtmlUnknownTag-->
<template>
  <div>
    <p>MielLink：基于 Hyperledger Fabric 的蜂蜜产品供应链溯源系统</p>
    <p @click="goToPage('/uplink')">信息录入</p>
    <p @click="goToPage('/trace')">溯源查询</p>
    <p @click="openExternalLink('http://43.156.142.179:8080')">Hyperledger 浏览器</p>
    <p @click="goToPage('/login')">登录 / 注册</p>
  </div>
</template>

<script>
export default {
  methods: {
    goToPage(path) {
      this.$router.push(path)
    },
    openExternalLink(url) {
      window.open(url, '_blank')
    }
  }
}
</script>

<style scoped lang='scss'>
</style>
