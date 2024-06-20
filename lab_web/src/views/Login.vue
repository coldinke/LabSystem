<template>
  <vue-particles
    id="tsparticles"
    @particles-loaded="particlesLoaded"
    :options="config"
  />
  <div class="formContainer">
    <h3>实验室管理系统</h3>
    <el-form
      ref="ruleFormRef"
      style="max-width: 600px"
      :model="ruleForm"
      :rules="rules"
      label-width="auto"
      class="ruleForm"
      status-icon
    >
      <el-form-item label="用户名" prop="username">
        <el-input v-model="ruleForm.username" />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="ruleForm.password" type="password" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm(ruleFormRef)">
          登陆
        </el-button>
        <el-button @click="resetForm(ruleFormRef)">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { useUserStore } from "../store/useUserStore";
import { config } from "../util/config";
import { useRouter } from "vue-router";
import { ref, reactive } from "vue";
import axios from "axios";
import { ElMessage } from "element-plus";

const { changeUser } = useUserStore();
const router = useRouter();
const ruleFormRef = ref();
const ruleForm = reactive({
  username: "",
  password: "",
});
const rules = reactive({
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
});

const submitForm = async (formEl) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      const res = await axios
        .post("/admin/api/login", ruleForm)
        .then((res) => {
          const { code, data } = res.data;
          changeUser(data.user);
          ElMessage.success("登陆成功");
        })
        .catch((err) => {
          ElMessage.error("用户名或密码错误");
          return false;
        });
      router.push("/");
    } else  {
      console.log("error submit!!");
    }
  });
};

const resetForm = (formEl) => {
  if (!formEl) return;
  formEl.resetFields();
};

const particlesLoaded = async (container) => {
  console.log("Particles container loaded", container);
};
</script>

<style lang="scss" scoped>
.formContainer {
  width: 500px;
  height: 300px;
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  color: white;
  text-shadow: 2px 2px 5px black;
  text-align: center;

  .ruleForm {
    margin-top: 50px;
  }

  h3 {
    font-size: 40px;
  }

  :deep(.el-form-item__label) {
    color: white;
    font-size: 16px;
    font-weight: 700;
  }
}
</style>
