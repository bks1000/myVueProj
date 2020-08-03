<template>
    <div id="applist">
        <el-button type="primary" icon="el-icon-search" v-on:click="query">查询</el-button>
        <el-button type="primary" icon="el-icon-info" v-on:click="info">说明</el-button>
        <el-table
            :data="lstData"
            stripe
            border
            height="800"
            style="width: 100%">
            <el-table-column
                prop="name"
                label="APP"
                width="180">
            </el-table-column>
            <el-table-column
                prop="description"
                label="描述"
                width="180">
            </el-table-column>
            <el-table-column
                prop="appIcon"
                label="Icon"
                width="180">
            </el-table-column>
            <el-table-column prop="codeName" label="codeName" width="180"></el-table-column>
            <el-table-column label="分组名称" prop="groupName" width="180"></el-table-column>
            <el-table-column
                prop="appDllName"
                label="DllName" width="180">
            </el-table-column>
            <el-table-column
                prop="url"
                label="url" width="380">
            </el-table-column>
            <el-table-column label="创建时间" prop="createTime" width="180"></el-table-column>
        </el-table>
    </div>
</template>

<script>
//import axios from 'axios'; //在main.js中通过Vue.prototype.axios = axios定义。
export default {
    name: 'AppList',
    data(){
        return {
            lstData:[]
        }
    },
    methods: {
        query(){
            this.axios.get('http://localhost:8080/apis/java')   //vue.config.js中定义了代理
                .then(resp=>{
                    console.log(resp.data);
                    this.lstData =  JSON.parse(resp.data);
                })
                .catch(error => console.log(error));
        },
        info(){
            const h = this.$createElement;

            this.$notify({
            title: '技术说明',
            message: h('i', { style: 'color: teal'}, '本功能展示了:  1、如何通过代理访问服务。2、VUE定义全局变量axios。3、go请求Hessian并处理返回JSON【对于数据量大服务器分包传输的，还存在问题】')
            });
        }
    },
}
</script>