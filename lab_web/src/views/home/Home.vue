<template>
    <div class="home-container" style="margin: 5px;">
    <el-row :gutter="20">
      <el-col :span="8">
        <el-card>
          <div class="grid-content">
            总用户数: <strong>{{ totalUsers }}</strong>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div class="grid-content">
            总实验室数: <strong>{{ totalLabs }}</strong>
          </div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card>
          <div class="grid-content">
            总预约数: <strong>{{ totalBooks }} </strong>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
    <div>
        <el-table :data="labList" style="width: 100%" @row-dblclick="handleRowClick">
      <el-table-column label="实验室名称" prop="lab_name" style="width: 10%" />
      <el-table-column label="所属学院" style="width: 10%">
        <template #default="{ row }">
          <span>{{ college_method(row.college_type) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="实验室类型" style="width: 5%">
        <template #default="{ row }">
          <span>{{ lab_method(row.lab_type) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="实验室安全等级" style="width: 5%">
        <template #default="{ row }">
          <span>{{ risk_method(row.risk_type) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="实验室容量" prop="capacity" style="width: 5%" />
    </el-table>

    </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import { LabType, RiskType, CollegeType } from "../../util/labs";
import { useRouter } from "vue-router";

const labList = ref([]);
const router = useRouter()
const totalLabs = ref(0);
const totalUsers = ref(0);
const totalBooks = ref(0);


onMounted(() => {
  getLabList();
  getLabCount();
  getUserCount();
  getBookCount();
});

const getUserCount = async () => {
  const res = await axios.get("/admin/api/users/count");
  totalUsers.value = res.data.data
};

const getLabCount = async () => {
  const res = await axios.get("/admin/api/labs/count");
  totalLabs.value = res.data.data
};

const getBookCount = async () => {
  const res = await axios.get("/admin/api/books/count");
  totalBooks.value = res.data.data
};

const getLabList = async () => {
  const res = await axios.get("/admin/api/labs");
  labList.value = res.data.data;
};

const college_method = (value) => {
  return CollegeType.find((item) => item.value === value).label;
};

const risk_method = (value) => {
  return RiskType.find((item) => item.value === value).label;
};

const lab_method = (value) => {
  return LabType.find((item) => item.value === value).label;
};

const handleRowClick = (row) => {
  router.push(`/labs/${row.ID}`);
};
</script>