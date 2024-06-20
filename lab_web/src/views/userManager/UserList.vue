<template>
  <div>
    <el-button type="primary" style="margin-bottom: 10px" @click="handleAdd"
      >添加用户</el-button
    >
    <el-table :data="userList" style="width: 100%">
      <el-table-column label="用户名" prop="username" style="width: 10%" />
      <el-table-column label="用户角色" style="width: 10%">
        <template #default="{ row }">
          <span>{{ roleTypeMap[row.role.roleType] }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button
            round
            :disabled="scope.row.is_default"
            @click="handleUpdate(scope.row)"
          >
            更新
          </el-button>
          <el-popconfirm
            confirm-button-text="Yes"
            cancel-button-text="No"
            title="你确定要删除吗？"
            @confirm="handleDelete(scope.row.ID)"
          >
            <template #reference>
              <el-button round :disabled="scope.row.is_default" type="danger">
                删除
              </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogFormVisible" title="用户更新" width="500">
      <el-form ref="updateFormRef" :model="updateForm" :rules="rules">
        <el-form-item
          label="用户名"
          prop="username"
          :label-width="formLabelWidth"
        >
          <el-input v-model="updateForm.username" autocomplete="off" />
        </el-form-item>
        <el-form-item
          label="密码"
          prop="password"
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="updateForm.password"
            type="password"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="角色" prop="roleID" :label-width="formLabelWidth">
          <el-select v-model="updateForm.roleID" placeholder="请设置权限">
            <el-option
              v-for="role in roleList"
              :key="role.ID"
              :label="roleTypeMap[role.roleType]"
              :value="role.ID"
            />
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

    <el-dialog v-model="dialogAddVisible" title="添加用户" width="500">
      <el-form ref="updateFormRef" :model="addForm" :rules="rules">
        <el-form-item
          label="用户名"
          prop="username"
          :label-width="formLabelWidth"
        >
          <el-input v-model="addForm.username" autocomplete="off" />
        </el-form-item>
        <el-form-item
          label="密码"
          prop="password"
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="addForm.password"
            type="password"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="角色" prop="roleID" :label-width="formLabelWidth">
          <el-select v-model="addForm.roleID" placeholder="请设置权限">
            <el-option
              v-for="role in roleList"
              :key="role.ID"
              :label="roleTypeMap[role.roleType]"
              :value="role.ID"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogAddVisible = false">取消</el-button>
          <el-button type="primary" @click="handleAddConfirm(updateFormRef)">
            确定添加
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from "vue";
import { roleTypeMap } from "../../util/roles";
import axios from "axios";
import { ElMessage } from "element-plus";

const formLabelWidth = "120px";
const userList = ref();
const roleList = ref();
const updateFormRef = ref();
const currentItem = ref({});
const updateForm = reactive({
  username: "",
  password: "",
  roleID: "",
});
const addForm = reactive({
  username: "",
  password: "",
  roleID: "",
});
const rules = {
  username: [
    { required: true, message: "请输入用户名", trigger: "blur" },
    {
      min: 3,
      max: 30,
      message: "长度应该在 3 - 30 这个范围内",
      trigger: "blur",
    },
  ],
  password: [
    { required: true, message: "请输入密码", trigger: "blur" },
    {
      min: 1,
      max: 11,
      message: "长度应该在 8 - 11 这个范围内",
      trigger: "blur",
    },
  ],
  roleID: [{ required: true, message: "请选择角色", trigger: "blur" }],
};
const dialogFormVisible = ref(false);
const dialogAddVisible = ref(false);

onMounted(() => {
  getUserList();
  getRoleList();
});

const getUserList = async () => {
  var res = await axios.get("/admin/api/users");
  userList.value = res.data.data;
};

const getRoleList = async () => {
  var res = await axios.get("/admin/api/roles");
  roleList.value = res.data.data;
};

const handleUpdate = (item) => {
  // console.log(item);
  currentItem.value = item;
  updateForm.username = item.username;
  updateForm.password = item.password;
  updateForm.roleID = item.role.ID;
  dialogFormVisible.value = true;
};

const handleConfirm = (formEl) => {
  if (!formEl) return;
  formEl.validate(async (valid, fields) => {
    if (valid) {
      console.log(updateForm.roleID);
      await axios.put(`/admin/api/users/${currentItem.value.ID}`, updateForm);
      await getUserList();
      dialogFormVisible.value = false;
      ElMessage.success("更新成功");
    } else {
      console.log("error submit!!");
    }
  });
};

const handleAdd = () => {
  dialogAddVisible.value = true;
};

const handleAddConfirm = (formEl) => {
  if (!formEl) return;
  formEl.validate(async (valid, fields) => {
    if (valid) {
      await axios.post(`/admin/api/users`, addForm);
      await getUserList();
      dialogAddVisible.value = false;
      ElMessage.success("添加成功");
    } else {
      console.log("error submit!!");
    }
  });
};

const handleDelete = async (id) => {
  // console.log(id);
  await axios.delete(`/admin/api/users/${id}`);
  await getUserList();
};
</script>
