<template>
  <div class="loading-wrapper">
    <n-spin :description="props.description"/>
  </div>
</template>

<script lang="ts" setup>
import {withDefaults} from 'vue'

defineOptions({
  name: 'SysLoading'
})
const props = withDefaults(
    defineProps<{
      description: string
    }>(),
    {
      description: undefined
    }
)
</script>

<style lang="scss" scoped>
.loading-wrapper {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  width: 100%;
  overflow: hidden;
  background: rgba($color: #ffffff, $alpha: 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
