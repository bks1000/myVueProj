<template>
    <div id='ele'>
       <Tools @add='add' @setting='setting'/>
       <el-dialog title="在Dialog对话框中使用组件" :visible.sync="AddVisible">
        <Hello/>
      </el-dialog>
      <el-dialog title="系统设置" :visible.sync="SettingVisible">
        <Setting/>
      </el-dialog>
    </div>
</template>

<script>
  // @ is an alias to /src
import Tools from '@/components/Tools.vue'
import Hello from '@/components/HelloWorld.vue'
import Setting from '@/views/System/Setting.vue'

export default {
  name: 'Ele',
  data(){
      return{
          SettingVisible:false,
          AddVisible:false
      }
  },
  components: {
    Tools,Hello,Setting
  },
  props: {

  },
  methods:{
      //在父组件处理子组件中的事件
      add(){
          console.log("父组件接收事件");
          alert('ele add!');
          this.AddVisible=true;
      },
      setting(){
          this.SettingVisible=true;
      }
  }
}
</script>

<style scoped>

</style>