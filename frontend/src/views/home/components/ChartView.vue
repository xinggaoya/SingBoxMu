<template>
  <v-chart class="chart" :option="option" />
</template>

<script lang="ts" setup>
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { LineChart } from "echarts/charts";
import {
  GridComponent,
  LegendComponent,
  TitleComponent,
  ToolboxComponent,
  TooltipComponent
} from "echarts/components";
import VChart, { THEME_KEY } from "vue-echarts";
import {provide, ref, watch} from "vue";
import {useEventsStore} from "@/stores/events/EventsStore";

use([
  CanvasRenderer,
  LineChart,
  TitleComponent,
  TooltipComponent,
  GridComponent,
  ToolboxComponent,
  LegendComponent
]);

provide(THEME_KEY, "light");

const option = ref<any>({
  title: {
    text: "实时流量 (最近1分钟)"
  },
  tooltip: {
    trigger: "axis",
    axisPointer: {
      type: "cross",
      label: {
        backgroundColor: "#6a7985"
      }
    }
  },
  legend: {
    data: ["上行", "下行"]
  },
  toolbox: {
    feature: {
      saveAsImage: {}
    }
  },
  grid: {
    left: "3%",
    right: "4%",
    bottom: "3%",
    containLabel: true
  },
  xAxis: {
    type: "time",
    splitLine: {
      show: false
    }
  },
  yAxis: {
    type: "value"
  },
  series: [
    {
      name: "上行",
      type: "line",
      smooth: true,
      areaStyle: {},
      data: []
    },
    {
      name: "下行",
      type: "line",
      smooth: true,
      areaStyle: {},
      data: []
    }
  ]
});

const MAX_DATA_POINTS = 6; // 限制数据点数量
const events = useEventsStore()

watch(events.traffic, (value)=>{
  handleTrafficData(value)
},{
  deep: true
})


function handleTrafficData(data:any) {
  const { up, down } = data;
  updateData(up, down);
}

function updateData(up: number, down: number) {
  const now = Date.now();

  if (option.value.series[0].data.length >= MAX_DATA_POINTS) {
    // 移除旧的数据点
    option.value.series[0].data.shift();
  }

  // 更新数据
  option.value.series[0].data.push([now, up]);
  option.value.series[1].data.push([now, down]);
}

</script>

<style scoped>
.chart {
  height: 100%;
}
</style>
