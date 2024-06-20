<template>
    <div class="common-layout">
        <el-container style="height: 100vh;">
            <SideMenu></SideMenu>
            <el-container direction="vertical">
                <TopHeader></TopHeader> 
                <el-main>
                    <el-scrollbar>
                        <router-view></router-view>
                    </el-scrollbar>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<script setup>
import SideMenu from "../components/MainBox/SideMenu.vue";
import TopHeader from "../components/MainBox/TopHeader.vue";
</script>