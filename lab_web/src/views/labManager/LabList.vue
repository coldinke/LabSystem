<template>
  <div>
    <el-button type="primary" style="margin-bottom: 10px" @click="handleAdd"
      >添加实验室</el-button
    >
    <el-table :data="labList" style="width: 100%">
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
      <el-table-column label="操作" style="width: 30%">
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

    <el-dialog v-model="dialogFormVisible" title="更新实验室" width="500">
      <el-form ref="updateFormRef" :model="updateForm" :rules="rules">
        <el-form-item
          label="实验室名"
          prop="lab_name"
          :label-width="formLabelWidth"
        >
          <el-input v-model="updateForm.lab_name" autocomplete="off" />
        </el-form-item>
        <el-form-item
          label="所属学院"
          prop="college_type"
          :label-width="formLabelWidth"
        >
          <el-select v-model="updateForm.college_type" placeholder="请设置学院">
            <el-option
              v-for="item in CollegeType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="实验室类型"
          prop="lab_type"
          :label-width="formLabelWidth"
        >
          <el-select
            v-model="updateForm.lab_type"
            placeholder="请设置实验室类型"
          >
            <el-option
              v-for="item in LabType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="实验室安全等级"
          prop="risk_type"
          :label-width="formLabelWidth"
        >
          <el-select
            v-model="updateForm.risk_type"
            placeholder="请设置实验室安全等级"
          >
            <el-option
              v-for="item in RiskType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="实验室容量"
          prop="capacity"
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="updateForm.capacity"
            type="number"
            autocomplete="off"
          />
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

    
    <el-dialog v-model="dialogAddVisible" title="添加实验室" width="500">
      <el-form ref="updateFormRef" :model="addForm" :rules="rules">
        <el-form-item
          label="实验室名"
          prop="lab_name"
          :label-width="formLabelWidth"
        >
          <el-input v-model="addForm.lab_name" autocomplete="off" />
        </el-form-item>
        <el-form-item
          label="所属学院"
          prop="college_type"
          :label-width="formLabelWidth"
        >
          <el-select v-model="addForm.college_type" placeholder="请设置学院">
            <el-option
              v-for="item in CollegeType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="实验室类型"
          prop="lab_type"
          :label-width="formLabelWidth"
        >
          <el-select v-model="addForm.lab_type" placeholder="请设置实验室类型">
            <el-option
              v-for="item in LabType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="实验室安全等级"
          prop="risk_type"
          :label-width="formLabelWidth"
        >
          <el-select
            v-model="addForm.risk_type"
            placeholder="请设置实验室安全等级"
          >
            <el-option
              v-for="item in RiskType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="实验室容量"
          prop="capacity"
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="addForm.capacity"
            type="number"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogaddVisible = false">取消</el-button>
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
import axios from "axios";
import { LabType, RiskType, CollegeType } from "../../util/labs";
import { ElMessage } from "element-plus";


const labList = ref([]);
const updateFormRef = ref();
const updateForm = reactive({
  lab_name: "",
  lab_type: "",
  risk_type: "",
  college_type: "",
  capacity: 0,
});
const addForm = reactive({
  lab_name: "",
  lab_type: "",
  risk_type: "",
  college_type: "",
  capacity: 0,
});
const dialogFormVisible = ref(false);
const dialogAddVisible = ref(false);
const currentItem = ref({});
const formLabelWidth = "120px";

onMounted(() => {
  getLabList();
});

const rules = {
  lab_name: [
    { required: true, message: "请输入实验室名称", trigger: "blur" },
    { type: "string", min: 3, message: "长度至少为3个字符", trigger: "blur" },
  ],
  lab_type: [
    { required: true, message: "请选择实验室类型", trigger: "change" },
    {
      type: "int",
      min: 1,
      max: 5,
      message: "实验室类型为1-5",
      trigger: "change",
    },
  ],
  risk_type: [
    { required: true, message: "请选择实验室安全等级", trigger: "change" },
    {
      type: "int",
      min: 1,
      max: 4,
      message: "实验室安全等级为1-4",
      trigger: "change",
    },
  ],
  college_type: [
    { required: true, message: "请选择所属学院", trigger: "change" },
    {
      type: "int",
      min: 1,
      max: 5,
      message: "所属学院为1-2",
      trigger: "change",
    },
  ],
  capacity: [
    { required: true, message: "请输入实验室容量", trigger: "blur" },
    {
      type: "int",
      min: 1,
      max: 1000,
    },
  ],
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

const handleUpdate = (item) => {
  currentItem.value = item;
  updateForm.lab_name = item.lab_name;
  updateForm.lab_type = item.lab_type;
  updateForm.college_type = item.college_type;
  updateForm.risk_type = item.risk_type;
  updateForm.capacity = item.capacity;
  dialogFormVisible.value = true;
};

const handleConfirm = (formEl) => {
  if (!formEl) return;
  formEl.validate(async (valid, fields) => {
    if (valid) {
      updateForm.capacity = parseInt(updateForm.capacity, 10);
      await axios.put(`/admin/api/labs/${currentItem.value.ID}`, updateForm);
      await getLabList();
      dialogFormVisible.value = false;
      ElMessage.success("更新成功");
    } else {
      console.log("error submit!!");
    }
  });
};

const handleDelete = async (id) => {
  console.log("delete", id);
  await axios.delete(`/admin/api/labs/${id}`);
  await getLabList();
};

const handleAdd = () => {
  dialogAddVisible.value = true;
};

const handleAddConfirm = (formEl) => {
  if (!formEl) return;
  formEl.validate(async (valid, fields) => {
    if (valid) {
      addForm.capacity = parseInt(addForm.capacity, 10);
      await axios.post(`/admin/api/labs`, addForm);
      await getLabList();
      dialogAddVisible.value = false;
      ElMessage.success("添加成功");
    } else {
      console.log("error submit!!");
    }
  });
};
</script>
