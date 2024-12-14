const chokidar = require('chokidar')
const bodyParser = require('body-parser')
const chalk = require('chalk')
const path = require('path')
const Mock = require('mockjs')
const mockDir = path.join(process.cwd(), 'mock')

function registerRoutes(app) {
    let mockLastIndex
    const { mocks } = require('./index.js')
    const mocksForServer = mocks.map(route => {
        return responseFake(route.url, route.type, route.response)
    })
    for (const mock of mocksForServer) {
        app[mock.type](mock.url, mock.response)
        mockLastIndex = app._router.stack.length
    }
    const mockRoutesLength = Object.keys(mocksForServer).length
    return {
        mockRoutesLength: mockRoutesLength,
        mockStartIndex: mockLastIndex - mockRoutesLength
    }
}

function unregisterRoutes() {
    Object.keys(require.cache).forEach(i => {
        if (i.includes(mockDir)) {
            delete require.cache[require.resolve(i)]
        }
    })
}

const responseFake = (url, type, respond) => {
    return {
        url: new RegExp(`http://43.156.142.179:9090${url}`),
        type: type || 'get',
        response(req, res) {
            console.log('Request invoke: ' + req.path)
            res.json(Mock.mock(respond instanceof Function ? respond(req, res) : respond))
        }
    }
}

module.exports = app => {
    app.use(bodyParser.json())
    app.use(bodyParser.urlencoded({
        extended: true
    }))
    const mockRoutes = registerRoutes(app)
    var mockRoutesLength = mockRoutes.mockRoutesLength
    var mockStartIndex = mockRoutes.mockStartIndex
    chokidar.watch(mockDir, {
        ignored: /mock-server/,
        ignoreInitial: true
    }).on('all', (event, path) => {
        if (event === 'change' || event === 'add') {
            try {
                app._router.stack.splice(mockStartIndex, mockRoutesLength)
                unregisterRoutes()
                const mockRoutes = registerRoutes(app)
                mockRoutesLength = mockRoutes.mockRoutesLength
                mockStartIndex = mockRoutes.mockStartIndex
                console.log(chalk.magentaBright(`Mock Server hot reload success! Changed ${path}`))
            } catch (error) {
                console.log(chalk.redBright(error))
            }
        }
    })
}