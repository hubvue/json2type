// import { spawn } from 'child_process'
const { exec } = require('child_process')

const args = process.argv.slice(2)

console.log(args)

const execCommand = `./json2type ${args.join(' ')}`
console.log('execCommand', execCommand)
exec(execCommand, (err, stdout, stderr) => {
  if (err) {
    console.log(err)
    return
  }
  if (stdout) {
    console.log(stdout);
  }
  if (stderr) {
    console.log(stderr)
  }
})