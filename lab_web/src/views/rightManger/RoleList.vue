<template>
  <div>
    <el-table :data="roleData" style="width: 100%" row-key="id">
      <el-table-column label="角色类型" style="width: 50%">
        <template #default="scope">
          <el-popover
            placement="right"
            title="权限详情"
            :width="200"
            trigger="hover"
            @before-enter="handlePopoverEnter(scope.row)"
          >
            <template #reference>
              <el-button class="m-2">{{
                roleTypeMap[scope.row.roleType]
              }}</el-button>
            </template>
            <template #default>
              <el-tree
                style="max-width: 600px"
                :data="rightData"
                :props="defaultProps"
                :default-expand-all="true"
                :render-content="renderContent"
              />
            </template>
          </el-popover>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button round @click="handleUpdate(scope.row)"> 更新 </el-button>
          <el-popconfirm
            confirm-button-text="Yes"
            cancel-button-text="No"
            title="你确定要删除吗？"
            @confirm="handleDelete(scope.row.ID)"
          >
            <template #reference>
              <el-button round type="danger"> 删除 </el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogFormVisible" title="权限更新" width="500">
      <el-form :model="updateFormRef">
        <el-form-item label="权限类型" :label-width="formLabelWidth">
          <el-select
            v-model="updateFormRef.roleType"
            placeholder="请设置权限类型"
          >
            <el-option
              v-for="role in roleData"
              :key="role.ID"
              :label="roleTypeMap[role.roleType]"
              :value="role.roleType">
              <span>{{ roleTypeMap[role.roleType] }}</span>    
            </el-option>
        </el-select>
        </el-form-item>
        <el-form-item label="权限" :label-width="formLabelWidth">
            <el-tree
                style="max-width: 600px"
                show-checkbox
                node-key="path"
                ref="treeRef"
                :data="rightData"
                :props="defaultProps"
              />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取消</el-button>
          <el-button type="primary" @click="handleConfirm()">
            确定更新
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { roleTypeMap } from "../../util/roles";
import { ref, reactive, onMounted, nextTick} from "vue";
import axios from "axios";

const formLabelWidth = "120px";
const roleData = ref([]);
const rightData = ref([]);
const dialogFormVisible = ref(false);
const updateFormRef = reactive({
    roleType: "",
    Rights: [],
});

onMounted(() => {
  getRoleList();
  getRightList();
});

const getRoleList = async () => {
  var res = await axios.get("admin/api/roles");
  roleData.value = res.data.data;
};

const getRightList = async () => {
  var res = await axios.get("admin/api/rights");
  rightData.value = res.data.data;
};

const treeRef = ref([])
const currentItem = ref({})
const handleUpdate = (item) => {
  console.log(item);
  dialogFormVisible.value = true;
  updateFormRef.roleType = item.roleType;
  currentItem.value = item;
  nextTick(() => {
    var checkedKeys = item.rights.map((right) => right.path);
    treeRef.value.setCheckedKeys(checkedKeys);
  });

};

const handleConfirm = async () => {
  const checkedKeys = treeRef.value.getCheckedKeys();
  var res = await axios.get("admin/api/rights");
  var rights = res.data.data;
  const checkedRights = rights.filter(right => checkedKeys.includes(right.path));
  console.log(checkedRights);
  await axios.put(`admin/api/roles/${currentItem.value.ID}`, {
    role_type: updateFormRef.roleType,
    rights: checkedRights,
  });
  dialogFormVisible.value = false;
};

const handleDelete = async (id) => {
  await axios.delete(`/admin/api/roles/${id}`);
  await getRoleList();
};


const currentRights = ref([]);
const handlePopoverEnter = ({ rights }) => {
  currentRights.value = rights;
};

const renderContent = (h, { node, data, store }) => {
  const isActive = currentRights.value.some((right) => right.ID === data.id);
  return h(
    "span",
    {
      class: `custom-tree-node ${isActive ? "active" : ""}`,
    },
    h("span", null, node.label)
  );
};

const defaultProps = {
  children: "children",
  label: "title",
};
</script>

<style lang="scss" scoped>
:deep(.active) {
  color: #d71fa6;
}
</style>
