/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/9 10:04
 * @version     v1.0
 * @filename    index.js
 * @description
 ***************************************************************************/

const cp = require('child_process')
const http = require('http')

http.createServer((req, res) => {
    let proc = cp.exec('./deploy.sh', ()=>{})
    proc.stdout.pipe(process.stdout)
    proc.stderr.pipe(process.stderr)
    res.end('deploy maxmicro')
}).listen(4000, () => {
    console.log('deploy be running on 4000')
})