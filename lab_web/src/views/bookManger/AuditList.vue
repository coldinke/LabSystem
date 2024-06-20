<template>
  <div>
    <el-table :data="fileterBookList" style="width: 100%">
      <el-table-column label="实验室名称" style="width: 10%">
        <template #default="{row}">
            <span>{{ getLabName(row.LabID)  }}</span>
        </template>
        </el-table-column>
      <el-table-column label="预约用户" style="width: 20%"  v-admin>
        <template #header>
          <el-input v-model="search" size="mini" placeholder="搜索" @change="getBookList" />
        </template>
      <template #default="scope">
        <span>{{ scope.row.Username }}</span>
      </template>
      </el-table-column>
      <el-table-column label="预约时间" style="width: 20%">
        <template #default="scope">
          <el-tag style="height: 40px;line-height: 20px;"> {{ moment(scope.row.BookTime).format("YYYY-MM-DD") }}
          <br/>
            {{  class_method(scope.row.ClassID) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="预约状态" prop="Status" style="width: auto">
        <template #default="{ row }">
          <el-tag :type=getStatusType(row.State)>{{ getStatus(row.State) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" style="width: 30%" v-teacher>
        <template #default="scope">
          <el-popconfirm
            confirm-button-text="Yes"
            cancel-button-text="No"
            title="你确定要允许吗？"
            @confirm="handleUpdate(scope.row.ID, 0)"
          >
            <template #reference>
              <el-button round type="primary"> 允许 </el-button>
            </template>
          </el-popconfirm>

          <el-popconfirm
            confirm-button-text="Yes"
            cancel-button-text="No"
            title="你确定要驳回吗？"
            @confirm="handleUpdate(scope.row.ID, 1)"
          >
            <template #reference>
              <el-button round type="danger"> 驳回 </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import moment from "moment";
import { useUserStore } from "../../store/useUserStore";
import { ClassType } from "../../util/labs";

const {user} = useUserStore();
const bookList = ref([]);
const labList = ref([]);
const search = ref("");

const fileterBookList = computed(() => {
  return bookList.value.filter(item=>item.Username.includes(search.value))
})

const getLabList = async () => {
  console.log(user.role.roleType)
  var res = await axios.get("/admin/api/labs");
  console.log(res.data.data)
  labList.value = res.data.data;
};

const getBookList = async ()   => {
  var res = await axios.get("/admin/api/books/process");
  bookList.value = res.data.data;
};

const getLabName = (id) => {
  const name = labList.value.find((item) => item.ID === id)?.lab_name;
  return name || '未找到实验室名称';
};

const statusTextMap = {
  2: '审核中',
  1: '拒绝',
  0: '接受'
}

const getStatus = (status) => statusTextMap[status] || '未知状态'

const getStatusType = (status) => {
  switch (status) {
    case 2:
      return 'info';
    case 0:
      return 'success';
    case 1:
      return 'danger';
    default:
      return 'info'; // 默认返回 'info'
  }
};

const class_method = (value) => {
  return ClassType.find((item) => item.value === value).label;
};

const handleUpdate = async (id, status) => {
  const res = await axios.put(`/admin/api/books/${id}`, { Status: status });
  if (res.data.code === 200) {
    getBookList();
  }
};

onMounted(() => {
  getBookList();
  getLabList();
});
</script>