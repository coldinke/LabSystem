<template>
  <div>
    <el-table :data="fileterBookList" style="width: 100%">
      <el-table-column label="实验室名称" style="width: 10%">
        <template #default="{row}">
            <span>{{ getLabName(row.LabID)  }}</span>
        </template>
        </el-table-column>
      <el-table-column label="预约用户" style="width: 20%">
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
      <el-table-column label="操作" style="width: 30%">
        <template #default="scope">
          <el-popconfirm
            confirm-button-text="Yes"
            cancel-button-text="No"
            title="你确定要撤销吗？"
            @confirm="handleDelete(scope.row.ID)"
          >
            <template #reference>
              <el-button round type="danger"> 撤销 </el-button>
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
  var res;
  if (user.role.roleType == 1000) {
    res = await axios.get("/admin/api/books");
  } else {
    res = await axios.get(`/admin/api/books/${user.username}`);
  }
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

const handleDelete = (id) => {
  axios.delete(`/admin/api/books/${id}`).then(() => {
    getBookList();
  });
};

onMounted(() => {
  getBookList();
  getLabList();
});


const vTeacher = {
  mounted(el) {
    if (user.role.roleType !== 1001) {
      el.remove();
    }
  }
}

const vAdmin = {
  mounted(el) {
    if (user.role.roleType !== 1000) {
      el.remove();
    }
  }
}
</script>
