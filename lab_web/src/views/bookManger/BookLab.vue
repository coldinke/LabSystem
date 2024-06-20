<template>
  <div>
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
      <el-table-column
        label="实验室可容纳人数"
        prop="capacity"
        style="width: 5%"
      />
      <el-table-column label="操作" style="width: 30%">
        <template #default="scope">
          <el-button round type="primary" @click="handleAdd(scope.row)">
            预约
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogAddVisible" title="预约实验室" width="500">
      <el-form ref="updateFormRef" :model="addForm" :rules="rules">
        <el-form-item
          label="预约时间"
          prop="book_time"
          :label-width="formLabelWidth"
        >
          <el-date-picker
            v-model="addForm.book_time"
            type="date"
            placeholder="选择日期"
            :disabled-date="disabledDate"
            style="width: 100%"
            @change="handleSelectChange"
          />
        </el-form-item>
        <el-form-item
          label="预约课时"
          prop="book_class"
          :label-width="formLabelWidth"
        >
          <el-select v-model="addForm.book_class" placeholder="请选择预约课时">
            <el-option
              v-for="item in ClassType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
              :disabled="disableSelect?.includes(item.value)"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="预约原因"
          prop="book_reason"
          
          :label-width="formLabelWidth"
        >
          <el-input v-model="addForm.book_reason" placeholder="请填写预约原因" type="textarea" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="handleCancel">取消</el-button>
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
import { ElMessage } from "element-plus";
import axios from "axios";
import { LabType, CollegeType, ClassType } from "../../util/labs";
import { useUserStore } from "../../store/useUserStore";

const {user} = useUserStore();
const labList = ref([]);
const updateFormRef = ref();
const addForm = reactive({
  book_time: "",
  book_class: "",
  book_reason: "",
});

const dialogAddVisible = ref(false);
const currentItem = ref({});
const formLabelWidth = "120px";

onMounted(() => {
  getLabList();
});

const rules = {
  book_time: [{ required: true, message: "请选择预约时间", trigger: "blur" }],
  book_class: [{ required: true, message: "请选择预约课时", trigger: "blur" }],
  book_reason: [{ required: true, message: "请填写预约原因", trigger: "blur" }],
};

const getLabList = async () => {
  const res = await axios.get("/admin/api/labs");
  labList.value = res.data.data;
};

const college_method = (value) => {
  return CollegeType.find((item) => item.value === value).label;
};

const lab_method = (value) => {
  return LabType.find((item) => item.value === value).label;
};

const disabledDate = (time) => {
  return time.getTime() < Date.now() || time.getTime() > (Date.now() + 7 * 24 * 60 * 60 * 1000);
};

const handleAdd = (item) => {
  dialogAddVisible.value = true;
  currentItem.value = item;
};

const clear = () => {
  addForm.book_time = "";
  addForm.book_class = "";
  addForm.book_reason = "";
};

const handleCancel = () => {
  clear();
  dialogAddVisible.value = false;
};

const handleAddConfirm = (formEl) => {
  if (!formEl) return;
  formEl.validate(async (valid, fields) => {
    if (valid) {
      await axios.post(`/admin/api/books`, {
        ...addForm,
        lab_id: currentItem.value.ID,
        book_username: user.username,
      });
      console.log(addForm)
      clear()
      dialogAddVisible.value = false;
      ElMessage.success("添加预约成功");
    } else {
      console.log("error submit!!");
    }
  });
};


const disableSelect = ref([])

const handleSelectChange = async () => {
  const res = await axios.get(`/admin/api/select/books/${currentItem.value.ID}`,
  {
    params: {
      lab_id: currentItem.value.ID,
      book_time: addForm.book_time,
    }
  });
  disableSelect.value = res.data.data.map(item => item.ClassID)
}
</script>
