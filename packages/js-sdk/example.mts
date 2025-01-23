import { Sandbox } from './dist'

import dotenv from 'dotenv'

dotenv.config()

for (let i = 0; i < 40; i++) {
  console.log('creating sandbox', i)
  const sbx = await Sandbox.create('fq7u0hwuq4wp1fsr1lnr', { timeoutMs: 5000 })

  await sbx.pause()

  try {
    const isRunning = await sbx.isRunning()
    if (isRunning) {
      console.log('sandbox is running after pause, breaking')

      break
    }
  } catch (e) {
    console.error('error checking if sandbox is running after pause, continuing', e)
    continue
  }

  await Sandbox.resume(sbx.sandboxId, { timeoutMs: 5000 })

  try {
    const isRunning = await sbx.isRunning()
    if (!isRunning) {
      console.log('sandbox is not running after resume, breaking')

      break
    }
  } catch (e) {
    console.error('error checking if sandbox is running after resume', e)
  }
}
