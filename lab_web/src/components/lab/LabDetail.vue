<template>
  <div>
    <el-card class="box-card">
      <el-header> 实验室详情 </el-header>

      <el-descriptions
        v-if="labDetail"
        title="实验室基本信息"
        :column="1"
        border
      >
        <el-descriptions-item label="实验室名称">{{
          labDetail.lab.lab_name
        }}</el-descriptions-item>
        <el-descriptions-item label="实验室类型">{{
          lab_method(labDetail.lab.lab_type)
        }}</el-descriptions-item>
        <el-descriptions-item label="风险类型">{{
          risk_method(labDetail.lab.risk_type)
        }}</el-descriptions-item>
        <el-descriptions-item label="实验室所属学院">{{
          college_method(labDetail.lab.college_type)
        }}</el-descriptions-item>
        <el-descriptions-item label="容纳人数">{{
          labDetail.lab.capacity
        }}</el-descriptions-item>
      </el-descriptions>

      <el-divider></el-divider>
      <h3>实验环境传感器数据</h3>
      <el-row :gutter="16">
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic v-if="sensorData" :value="sensorData.Temperature + '℃'">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  当前温度
                  <el-tooltip
                    effect="dark"
                    content="当前实验室温度"
                    placement="top"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning />
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>最后更新时间</span>
                <span
                  :class="{
                    red: lastUpdate,
                  }"
                >
                  <span class="time">{{ lastUpdate }}</span>
                </span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic v-if="sensorData" :value="sensorData.Humidity + '%'">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  当前湿度
                  <el-tooltip
                    effect="dark"
                    content="当前实验室湿度"
                    placement="top"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning />
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>最后更新时间</span>
                <span
                  :class="{
                    red: lastUpdate,
                  }"
                >
                  <span class="time">{{ lastUpdate }}</span>
                </span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic v-if="sensorData" :value="sensorData.Fire == 1 ? '正常' : '异常'">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  当前火焰状态
                  <el-tooltip
                    effect="dark"
                    content="当前实验室火焰传感器状态"
                    placement="top"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning />
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div v-if="sensorData" class="footer-item">
                <span>火焰继电器状态</span>
                <span :class="{ green: sensorData.Fire === 1, red: sensorData.Fire === 0 }">
                {{ sensorData.Fire === 1 ? "正常" : "异常" }}
              </span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic v-if="sensorData" :value="sensorData.Smoke == 1 ? '正常' : '异常'">
              <template #title>
                <div style="display: inline-flex; align-items: center">
                  当前烟雾状态
                  <el-tooltip
                    effect="dark"
                    content="当前实验室烟雾传感器状态"
                    placement="top"
                  >
                    <el-icon style="margin-left: 4px" :size="12">
                      <Warning />
                    </el-icon>
                  </el-tooltip>
                </div>
              </template>
            </el-statistic>
            <div v-if="sensorData" class="statistic-footer">
              <div class="footer-item">
                <span>烟雾继电器状态</span>
                <span :class="{ green: sensorData.Smoke === 1, red: sensorData.Smoke === 0 }">
                {{ sensorData.Smoke === 1 ? "正常" : "异常" }}
                </span>
              </div>
            </div>
          </div>
        </el-col> 
      </el-row>
    </el-card>
  </div>
</template>
<script setup>
import { onMounted, ref, onUnmounted } from "vue";
import { useRoute } from "vue-router";
import axios from "axios";
import {
  ArrowRight,
  CaretBottom,
  CaretTop,
  Warning,
} from "@element-plus/icons-vue";
import { LabType, RiskType, CollegeType } from "../../util/labs";

const labDetail = ref(null);
const sensorData = ref(null);
const fireStatus = ref('');
let intervalId = null;
const lastUpdate = ref();

const college_method = (value) => {
  return CollegeType.find((item) => item.value === value).label;
};

const risk_method = (value) => {
  return RiskType.find((item) => item.value === value).label;
};

const lab_method = (value) => {
  return LabType.find((item) => item.value === value).label;
};

onMounted(() => {
  const route = useRoute();
  const id = route.params.labId;
  getLabDetail(id);
  fetchSensorData(id);
  intervalId = setInterval(fetchSensorData, 6000);
});

onUnmounted(() => {
  clearInterval(intervalId);
});

const getLabDetail = async (id) => {
  try {
    const res = await axios.get(`/admin/api/labs/${id}`);
    labDetail.value = res.data.data;
    console.log(labDetail.value);
  } catch (error) {
    console.error(error);
  }
};

const fetchSensorData = async () => {
  try {
    const res = await axios.get(`/admin/api/sensor/${labDetail.value.lab.ID}`);
    sensorData.value = res.data.data;
    lastUpdate.value = new Date(sensorData.value.CreatedAt).toLocaleString();
    console.log(sensorData.value);
  } catch (error) {
    console.error(error);
  }
};
</script>

<style scoped>
.el-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background-color: #409eff;
  color: #fff;
  text-align: right;
}

:global(h2#card-usage ~ .example .example-showcase) {
  background-color: var(--el-fill-color) !important;
}

.el-statistic {
  --el-statistic-content-font-size: 28px;
}

.statistic-card {
  height: 100%;
  padding: 20px;
  border-radius: 4px;
  background-color: var(--el-bg-color-overlay);
}

.statistic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  font-size: 12px;
  color: var(--el-text-color-regular);
  margin-top: 16px;
}

.statistic-footer .footer-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.statistic-footer .footer-item span:last-child {
  display: inline-flex;
  align-items: center;
  margin-left: 4px;
}

.green {
  color: var(--el-color-success);
}
.red {
  color: var(--el-color-error);
}
</style>
