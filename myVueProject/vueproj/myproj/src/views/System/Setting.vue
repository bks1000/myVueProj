<template>
    <div id="setting">
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="密码" prop="pass">
                <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="checkPass">
                <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="年龄" prop="age">
                <el-input v-model.number="ruleForm.age"></el-input>
            </el-form-item>
            <el-form-item label="是否启用" prop="enable">
                <div style="text-align:left">
                    <el-switch
                    v-model="ruleForm.enable"
                    active-color="#13ce66"
                    inactive-color="#ff4949">
                    </el-switch>
                </div>
            </el-form-item>
            <el-form-item label="显示列数" prop="colcnt">
                <div style="text-align:left">
                    <el-slider
                    v-model="ruleForm.colcnt"
                    :step="1"
                    :max="5" :min="0"
                    show-stops>
                    </el-slider>
                </div>
            </el-form-item>
            <el-form-item label="" prop="colcnt">
                <el-form :inline="true" :model="ruleForm" class="demo-form-inline">
                    <el-form-item label="起始日期">
                     <el-date-picker
                        v-model="ruleForm.startdate"
                        type="date"
                        placeholder="选择日期">
                    </el-date-picker>
                    </el-form-item>
                    <el-form-item label="结束日期">
                        <el-date-picker
                        v-model="ruleForm.enddate"
                        type="date"
                        placeholder="选择日期">
                        </el-date-picker>
                    </el-form-item>
                </el-form>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name:"Setting",
        data() {
            return {
                 ruleForm: {
                    pass: '',
                    checkPass: '',
                    age: '',
                    enable:false,
                    colcnt:4,
                    startdate:'',
                    enddate:''
                }
            }
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                if (valid) {
                    alert('submit!');
                } else {
                    console.log('error submit!!');
                    return false;
                }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            }
        }
    }
</script>

<style lang="scss" scoped>

</style>