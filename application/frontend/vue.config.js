'use strict'

const path = require('path')
const defaultSettings = require('./src/settings.js')

function resolve(dir) {
    return path.join(__dirname, dir)
}

const name = defaultSettings.title || 'UserInterface'
const port = process.env.port || process.env.npm_config_port || 3000

// noinspection JSUnusedGlobalSymbols
module.exports = {
    publicPath: '/',
    outputDir: 'dist',
    assetsDir: 'static',
    productionSourceMap: false,
    devServer: {
        port: port,
        open: true,
        overlay: {
            warnings: false,
            errors: true
        }
    },
    configureWebpack: {
        name: name,
        resolve: {
            alias: {
                '@': resolve('src')
            }
        }
    },
    chainWebpack(config) {
        config.plugin('preload').tap(() => [{
            rel: 'preload',
            fileBlacklist: [/\.map$/, /hot-update\.js$/, /runtime\..*\.js$/],
            include: 'initial'
        }])
        config.plugins.delete('prefetch')
        config.when(process.env.NODE_ENV !== 'development',
            config => {
                config.optimization.splitChunks({
                    chunks: 'all',
                    cacheGroups: {
                        libs: {
                            name: 'chunk-libs',
                            test: /[\\/]node_modules[\\/]/,
                            priority: 10,
                            chunks: 'initial'
                        },
                        elementUI: {
                            name: 'chunk-elementUI',
                            priority: 20,
                            test: /[\\/]node_modules[\\/]_?element-ui(.*)/
                        },
                        commons: {
                            name: 'chunk-commons',
                            test: resolve('src/components'),
                            minChunks: 3,
                            priority: 5,
                            reuseExistingChunk: true
                        }
                    }
                })
                config.optimization.runtimeChunk('single')
            }
        )
    }
}
