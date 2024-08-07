import { App } from 'vue'
import { loading } from './loading'

export const createDirectives = (app: App) => {
  loading(app)
}
