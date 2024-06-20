<template>
  <div>
    <el-table :data="tableData" style="width: 100%" row-key="id">
      <el-table-column prop="title" label="权限名称" width="200" />
      <el-table-column prop="path" label="路径" width="280" />
      <!-- <el-table-column prop="permissions" label="权限" width="200" /> -->
      <el-table-column label="图标" width="200">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <el-icon size="25px">
              <component :is="mapIcons[scope.row.icon]" />
            </el-icon>
            <span style="margin-left: 10px">{{ scope.row.date }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button round @click="handleUpdate(scope.row)"> 更新 </el-button>
          <el-popconfirm
            confirm-button-text="Yes"
            cancel-button-text="No"
            title="你确定要删除吗？"
            @confirm="handleDelete(scope.row.id)"
          >
            <template #reference>
              <el-button round type="danger"> 删除 </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>
    <!-- dialog -->
    <el-dialog v-model="dialogFormVisible" title="权限更新" width="500">
      <el-form ref="updateFormRef" :model="updateForm" :rules="rules">
        <el-form-item
          label="路径名称"
          prop="title"
          :label-width="formLabelWidth"
        >
          <el-input v-model="updateForm.title" autocomplete="off" />
        </el-form-item>
        <el-form-item
          label="当前路径"
          prop="path"
          :label-width="formLabelWidth"
        >
          <el-input v-model="updateForm.path" autocomplete="off" />
        </el-form-item>
        <!-- <el-form-item
          label="角色选择"
          :label-width="formLabelWidth"
        >
          <el-select v-model="updateForm.permissions" placeholder="请设置权限">
            <el-option label="实验室管理员" value="1000" />
            <el-option label="教师" value="1001" />
            <el-option label="学生" value="1002" />
            <el-option label="访客" value="1003" />
          </el-select>
        </el-form-item> -->
        <el-form-item label="图标" prop="icon" :label-width="formLabelWidth">
          <el-select v-model="updateForm.icon" placeholder="请设置图标">
            <el-option
              v-for="(icon, name) in mapIcons"
              :key="name"
              :label="name"
              :value="name"
            >
              <span>{{ name }}</span>
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取消</el-button>
          <el-button type="primary" @click="handleConfirm(updateFormRef)">
            确定更新
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { mapIcons } from "../../util/icons";
import { onMounted, ref, reactive } from "vue";
import axios from "axios";

const formLabelWidth = "120px";
const tableData = ref([]);
const dialogFormVisible = ref(false);
const currentItem = ref({});
const updateFormRef = ref();
const updateForm = reactive({
  title: "",
  path: "",
  icon: "",
});
const rules = reactive({
  title: [
    { required: true, message: "请输入权限名称", trigger: "blur" },
    { min: 1, message: "路径名称长度必须大于等于1", trigger: "blur" },
  ],
  path: [
    { required: true, message: "请输入路径", trigger: "blur" },
    { min: 1, message: "路径长度必须大于等于1", trigger: "blur" },
  ],
  icon: [{ required: true, message: "请选择图标", trigger: "blur" }],
});

onMounted(() => {
  getList();
});

const getList = async () => {
  var res = await axios.get("admin/api/rights");
  tableData.value = res.data.data;
};

const handleUpdate = (item) => {
  // console.log(item);
  currentItem.value = item;
  updateForm.title = item.title;
  updateForm.path = item.path;
  updateForm.icon = item.icon;
  updateForm.permissions = item.permissions;
  dialogFormVisible.value = true;
};

const handleConfirm = (formEl) => {
  if (!formEl) return;
  console.log(updateForm)
  formEl.validate(async (valid, fields) => {
    if (valid) {
      await axios.put(`admin/api/rights/${currentItem.value.id}`, updateForm);
      await getList();
      dialogFormVisible.value = false;
    } else {
      console.log("error submit!", fields);
      return false;
    }
  });
};

const handleDelete = async (id) => {
  console.log(id);
  await axios.delete(`/admin/api/rights/${id}`);
  await getList();
};
</script>
