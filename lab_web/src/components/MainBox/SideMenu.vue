<template>
    <el-aside width="200px">
        <el-scrollbar>
            <el-menu :default-active="route.fullPath" class="el-menu-vertical-demo" style="height:100vh;" :router="true">
                <template v-for="data in rightData" :key="data.path">
                    <el-sub-menu :index="data.path" v-if="data.children.length && checkAuth(data.path)">
                        <template #title>
                            <el-icon>
                                <component :is="mapIcons[data.icon]" />
                            </el-icon>
                            <span>{{ data.title }}</span>
                        </template>
                        <template  v-for="item in data.children" :key="item.path">
                            <el-menu-item :index="item.path" v-if="checkAuth(item.path)">
                            <el-icon>
                                <component :is="mapIcons[item.icon]" />
                            </el-icon>
                            <span>{{ item.title }}</span>
                        </el-menu-item>
                        </template>
                    </el-sub-menu>
                    <el-menu-item :index="data.path" v-else-if="checkAuth(data.path)" >
                        <el-icon>
                            <component :is="mapIcons[data.icon]" />
                        </el-icon>
                        <span>{{ data.title }}</span>
                    </el-menu-item>
                </template>
            </el-menu>
        </el-scrollbar>
    </el-aside>
</template>

<script setup>
import { mapIcons } from '../../util/icons'
import { onMounted, ref } from 'vue';
import axios from 'axios';
import { useRoute } from 'vue-router';
import {useUserStore} from '../../store/useUserStore'

const route = useRoute()
onMounted(() => {
    getrightList()
})

const rightData = ref([])
const getrightList = async () => {
    var res = await axios.get('admin/api/rights').then(res => {
        rightData.value = res.data.data
    })
}

const {user} = useUserStore()
const checkAuth = (path) => {
    return user.role.rights.some(right => right.path === path);
}

</script>