<template>
  <div class='navbar'>
    <hamburger :is-active='sidebar.opened' class='hamburger-container' @toggleClick='toggleSideBar'/>
    <breadcrumb class='breadcrumb-container'/>
    <div class='right-menu'>
      <span style='display:block' @click='logout' class='logout-icon'>
        <span>注销</span>
        <!--suppress HtmlUnknownTag-->
        <el-icon class='el-icon-switch-button' style='font-size: 19px; margin-left: 4px; transform: translateY(2px)'/>
      </span>
    </div>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import Hamburger from '@/components/Hamburger'
import sidebar from '@/layout/components/Sidebar'

export default {
  components: {
    Breadcrumb,
    Hamburger
  },
  computed: {
    sidebar() {
      return sidebar
    },
    ...mapGetters([
      'sidebar',
      'avatar'
    ])
  },
  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    async logout() {
      await this.$store.dispatch('user/logout')
      this.$router.push(`/login?redirect=${this.$route.fullPath}`)
    }
  }
}
</script>

<style scoped lang='scss'>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  box-shadow: 0 1px 4px rgba(0, 21, 41, .08);

  .hamburger-container {
    line-height: 46px;
    height: 100%;
    float: left;
    cursor: pointer;
    transition: background .3s;
    -webkit-tap-highlight-color: transparent;

    &:hover {
      background: rgba(0, 0, 0, .025)
    }
  }

  .breadcrumb-container {
    float: left;
  }

  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;

    .logout-icon {
      font-size: 14px;
      color: #666666;
      margin-right: 16px;
      font-weight: bold;
    }

    &:hover .logout-icon {
      color: #E2A130;
      cursor: pointer;
    }

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5A5E66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background .3s;

        &:hover {
          background: rgba(0, 0, 0, .025)
        }
      }
    }

    .avatar-container {
      margin-right: 30px;

      .avatar-wrapper {
        margin-top: 5px;
        position: relative;

        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
    }
  }
}
</style>

<style scoped lang='scss'>
* {
  user-select: none;
}
</style>
