<!--suppress JSUnresolvedReference-->
<!--suppress HtmlUnknownTag-->
<template>
  <div v-if='!item.hidden'>
    <template v-if='hasOneShowingChild(item.children,item) && (!onlyOneChild.children || onlyOneChild.noShowingChildren) && !item.alwaysShow'>
      <app-link v-if='onlyOneChild.meta' :to='resolvePath(onlyOneChild.path)'>
        <el-menu-item :index='resolvePath(onlyOneChild.path)' :class="{ 'submenu-title-noDropdown': !isNest }">
          <item :icon='onlyOneChild.meta.icon || (item.meta && item.meta.icon)' :title='onlyOneChild.meta.title'/>
        </el-menu-item>
      </app-link>
    </template>
    <el-submenu v-else ref='subMenu' :index='resolvePath(item.path)' popper-append-to-body>
      <!--suppress HtmlDeprecatedAttribute -->
      <template slot='title'>
        <item v-if='item.meta' :icon='item.meta && item.meta.icon' :title='item.meta.title'/>
      </template>
      <sidebar-item v-for='child in item.children'
                    :key='child.path'
                    :is-nest='true'
                    :item='child'
                    :base-path='resolvePath(child.path)'
                    class='nest-menu'/>
    </el-submenu>
  </div>
</template>

<script>
import path from 'path'
import {isExternal} from '@/utils/validate'
import Item from './Item'
import AppLink from './Link'

// noinspection JSUnusedLocalSymbols
export default {
  name: 'SidebarItem',
  components: {Item, AppLink},
  props: {
    item: {
      type: Object,
      required: true
    },
    isNest: {
      type: Boolean,
      default: false
    },
    basePath: {
      type: String,
      default: ''
    }
  },
  data() {
    return {}
  },
  methods: {
    hasOneShowingChild(children = [], parent) {
      const showingChildren = children.filter(item => {
        return !item.hidden
      })
      if (showingChildren.length === 1) {
        return true
      }
      return showingChildren.length === 0
    },
    resolvePath(routePath) {
      if (isExternal(routePath)) {
        return routePath
      }
      if (isExternal(this.basePath)) {
        return this.basePath
      }
      return path.resolve(this.basePath, routePath)
    }
  }
}
</script>
