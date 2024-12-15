<!--suppress HtmlUnknownTag-->
<template>
  <div class='page-container'>
    <div class='background'>
      <img src='/images/homeBackground.png' alt='Home background image'>
      <img src='/images/logo.png' alt='Website logo'>
      <img src='/images/title.png' alt='Page title'>
      <img src='/images/honey.png' alt='Honey icon'>
      <div class="login">
        <button @click="goToPage('/login')" class='login-button'>登陆</button>
        <button @click="goToPage('/login')" class='signup-button'>注册</button>
      </div>
    </div>

    <div class='selection'>
      <div @click="goToPage('/uplink')" class='selectionCard'>
        <img src='/images/selectionContainer.png' alt='Data entry selection background'>
        <img src='/images/selection1.png' alt='Data entry icon'>
        <button class='selection-button'>信息录入</button>
      </div>
      <div @click="goToPage('/trace')" class="selectionCard">
        <img src='/images/selectionContainer.png' alt='Trace query selection background'>
        <img src='/images/selection2.png' alt='Trace query icon'>
        <button class='selection-button'>溯源查询</button>
      </div>
      <div @click="openExternalLink('http://43.156.142.179:8080')" class='selectionCard'>
        <img src='/images/selectionContainer.png' alt='Hyperledger browser selection background'>
        <img src='/images/selection3.png' alt='Hyperledger browser icon'>
        <button class='selection-button'>Hyperledger 浏览器</button>
      </div>
    </div>
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
.page-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  background: #FFF6EB;

  .background {
    position: relative;
    width: 100%;
    height: 100vh;

    img:first-of-type {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: auto;
      object-fit: cover;
      z-index: 1;
    }

    img:nth-of-type(2) {
      position: absolute;
      top: 1%;
      left: 1%;
      width: 25%;
      height: auto;
      z-index: 2;
    }

    img:nth-of-type(3) {
      position: absolute;
      top: 38%;
      left: 16%;
      width: 35%;
      height: auto;
      z-index: 3;
    }

    img:nth-of-type(4) {
      position: absolute;
      top: 17%;
      right: 10%;
      width: 25%;
      height: auto;
      z-index: 4;
    }

    .login {
      position: absolute;
      top: 3%;
      right: 2%;
      display: flex;
      align-items: center;
      gap: 5px;
      padding: 5px;
      z-index: 5;
    }

    .login-button,
    .signup-button {
      padding: 10px 20px;
      font-size: 14px;
      border: none;
      cursor: pointer;
    }

    .login-button {
      background-color: #B16400;
      color: #FFF;
      border-radius: 20px 0 0 20px;

      &:hover {
        background-color: #A35C00;
      }

      &:active {
        background-color: #6F3C00;
      }
    }

    .signup-button {
      background-color: #FAFAFA;
      color: #8F5100;
      border-radius: 0 20px 20px 0;

      &:hover {
        background-color: #FFFFFF;
      }

      &:active {
        background-color: #EDEDED;
      }
    }
  }

  .selection {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    position: absolute;
    bottom: 10%;
    padding: 20px 200px;
    background: #FFF6EB;
    cursor: pointer;

    .selectionCard {
      position: relative;
      width: 100%;
      max-width: 300px;
      margin: 0 auto;
      text-align: center;
      border-radius: 10px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      transition: box-shadow 0.3s ease, transform 0.3s ease;

      &:hover {
        box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
        transform: scale(1.05);
      }

      &:active {
        box-shadow: inset 0 0 8px rgba(0, 0, 0, 0.3);
        transform: scale(0.98);
      }

      img:first-of-type {
        position: relative;
        width: 100%;
        height: auto;
      }

      img:last-of-type {
        position: absolute;
        top: 20%;
        left: 25%;
        width: 50%;
        height: auto;
        z-index: 2;
      }

      .selection-button {
        cursor: pointer;
        position: absolute;
        top: 80%;
        left: 50%;
        transform: translateX(-50%);
        width: 78%;
        padding: 10px 20px;
        font-size: 20px;
        color: #FFF;
        background-color: #8A4B00;
        border: none;
        border-radius: 20px;
      }
    }
  }
}
</style>
