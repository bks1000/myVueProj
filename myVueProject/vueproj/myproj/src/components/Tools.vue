<template>
    <div id='tools'>
        <el-button-group>
            <el-button type="primary" icon="el-icon-arrow-left">上一页</el-button>
            <el-button type="primary">下一页<i class="el-icon-arrow-right el-icon--right"></i></el-button>
        </el-button-group> &nbsp;
        <el-button-group>
            <el-button type="primary" icon="el-icon-document-add" v-on:click="add">新增</el-button>
            <el-button type="primary" icon="el-icon-edit">修改</el-button>
            <el-button type="primary" icon="el-icon-delete">删除</el-button>
            <el-button type="primary" icon="el-icon-share">分享</el-button>
            <el-button type="primary" icon="el-icon-setting" @click="setting">设置</el-button>
        </el-button-group>
    </div>
</template>

<script>
  export default {
  name: 'Tools',
   created() {
    this.$on("add", function() {
      console.log("子组件接收自己发射的事件");
    });
  },
  methods: {
    add() {
      console.log("子组件发射事件");
      this.$emit("add");
    },
    setting(){
        this.$emit("setting");
    }
  }
}
</script>