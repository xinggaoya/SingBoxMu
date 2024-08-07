import {defineConfig} from "vite";
import {fileURLToPath, URL} from 'node:url'
import vue from "@vitejs/plugin-vue";
import {internalIpV4} from "internal-ip";
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {NaiveUiResolver} from 'unplugin-vue-components/resolvers'
import VueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig(async () => ({
    plugins: [
        vue(),
        VueDevTools(),
        AutoImport({
            resolvers: [NaiveUiResolver()],
        }),
        Components({
            resolvers: [NaiveUiResolver()],
        }),
    ],

    // Vite options tailored for Tauri development and only applied in `tauri dev` or `tauri build`
    //
    // 1. prevent vite from obscuring rust errors
    clearScreen: false,
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url)),
            '@api': fileURLToPath(new URL('./bindings', import.meta.url)),
        }
    },
    // 2. tauri expects a fixed port, fail if that port is not available
    server: {
        port: 1420,
        strictPort: true,
        host: "0.0.0.0",
        hmr: {
            protocol: "ws",
            host: await internalIpV4(),
            port: 1421,
        },
        watch: {
            // 3. tell vite to ignore watching `src-tauri`
            ignored: ["**/src-tauri/**"],
        },
    },
}));
