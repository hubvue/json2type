const { resolve } = require('path')
const { exec } = require('child_process')
const fs = require('fs/promises')

const BIN_PATH = '/Users/didi/work/bin/'
const ARCH_MAPPING = {
  "ia32": "386",
  "x64": "amd64",
  "arm": "arm"
}


const install = async () => {
  try {
    const commandPath = resolve(__dirname, `../dist/json2type_${process.platform}_${ARCH_MAPPING[process.arch]}/json2type`)
    await execShellCommand(`cp ${commandPath} ${BIN_PATH}/json2type`)
    console.log('Install cli successfully!')
  } catch(err) {
    console.warn('Install fail: ', err)
  }
}

const uninstall = async () => {
  try {
    await fs.unlink(`${BIN_PATH}/json2type`)
    console.log("Uninstall cli successfully!")
  } catch(err) {
    console.warn('Uninstall fail: ', err)
  }
}

const execShellCommand = (cmd) => {
  return new Promise((resolve, reject) => {
    exec(cmd, (error, stdout, stderr) => {
      if (error) {
        reject(error)
      }
      resolve(stdout? stdout : stderr)
    })
  })
}

const actions = {
  install,
  uninstall
}

const run = async () => {
  try {
    const args = process.argv.slice(2)
    const cmd = args[0]
    if (!actions[cmd]) {
      console.log("Invalid command. `install` and `uninstall` are the only supported commands");
      process.exit(1);
    }
    await actions[cmd]()
    process.exit(0)
  } catch(err) {
    console.log(err)
    process.exit(1)
  }
}

run()