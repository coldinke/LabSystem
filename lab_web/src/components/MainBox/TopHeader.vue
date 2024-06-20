<template>
    <el-header>
        <div>智慧实验室管理系统</div>
        <div>
            <span style="line-height: 40px; padding-right: 10px;">欢迎 {{user.username}} 回来</span>
            <el-dropdown>
                <span class="el-dropdown-link">
                    <el-avatar :size="40" :src="circleUrl" />
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item>{{ user.username }}</el-dropdown-item>
                        <el-dropdown-item>{{ roleTypeMap[user.role.roleType] }}</el-dropdown-item>
                        <el-dropdown-item @click="handleExit">退出</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>

        </div>
    </el-header>
</template>

<script setup>
import { useUserStore } from "../../store/useUserStore";
import { useRouterStore } from "../../store/useRouterStore";
import { useRouter } from "vue-router";
import { roleTypeMap } from "../../util/roles";

const circleUrl = "https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"

const { changeUser, user } = useUserStore()
const { changeRouter } = useRouterStore()
const router = useRouter()

console.log(user)

const handleExit = () => {

    changeUser({})
    changeRouter(false)

    router.push("/login")
}

</script>

<style lang="scss" scoped>
.el-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 20px;
    background-color: #409EFF;
    color: #fff;
}
</style>