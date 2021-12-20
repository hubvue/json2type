import { exec } from 'child_process'
import { readFile, stat } from 'fs/promises'
import { resolve } from 'path'

type Language = 'typescript' | 'go'
export interface CommandOptions {
  input: string
  language: Language
  output?: string
  name?: string
}

const ARCH_MAPPING: Record<string, string> = {
  "ia32": "386",
  "x64": "amd64",
  "arm": "arm"
}

const json2type = async (options: CommandOptions) => {
  try {
    // path detection
    if (options.input[0] !== '/' || (options.output && options.output[0] !== '/')) {
      throw new Error('please use absolute path by ouput or input')
    }
    // command params
    const params = getCommandParams(options)

    // get command path
    const extname = process.platform === 'win32' ? '.exe' : ''
    const commandPath = resolve(__dirname, `../dist/json2type_${process.platform}_${ARCH_MAPPING[process.arch]}/json2type${extname}`)

    const result = await execShellCommand(`${commandPath} ${params}`)
    return result
  } catch(err) {
    console.log(err)
  }
}

const getCommandParams = (options: CommandOptions) => {
  let params = ''
  Object.keys(options).forEach(key => {
    params += `-${key}=${options[key as keyof CommandOptions]} `
  })
  return params
}

const execShellCommand = (cmd: string) => {
  return new Promise((resolve, reject) => {
    exec(cmd, (error, stdout, stderr) => {
      if (error) {
        reject(error)
      }
      resolve(stdout? stdout : stderr)
    })
  })
}

export default json2type