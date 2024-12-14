import store from '@/store'

const {body} = document
const WIDTH = 992

export default {
    watch: {
        $route() {
            if (this.device === 'mobile' && this.sidebar.opened) {
                // noinspection JSIgnoredPromiseFromCall
                store.dispatch('app/closeSideBar', {withoutAnimation: false})
            }
        }
    },
    beforeMount() {
        window.addEventListener('resize', this.$_resizeHandler)
    },
    beforeDestroy() {
        window.removeEventListener('resize', this.$_resizeHandler)
    },
    mounted() {
        const isMobile = this.$_isMobile()
        if (isMobile) {
            // noinspection JSIgnoredPromiseFromCall
            store.dispatch('app/toggleDevice', 'mobile')
            // noinspection JSIgnoredPromiseFromCall
            store.dispatch('app/closeSideBar', {withoutAnimation: true})
        }
    },
    methods: {
        $_isMobile() {
            const rect = body.getBoundingClientRect()
            return rect.width - 1 < WIDTH
        },
        $_resizeHandler() {
            if (!document.hidden) {
                const isMobile = this.$_isMobile()
                // noinspection JSIgnoredPromiseFromCall
                store.dispatch('app/toggleDevice', isMobile ? 'mobile' : 'desktop')
                if (isMobile) {
                    // noinspection JSIgnoredPromiseFromCall
                    store.dispatch('app/closeSideBar', {withoutAnimation: true})
                }
            }
        }
    }
}
