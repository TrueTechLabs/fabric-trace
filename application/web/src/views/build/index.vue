<template>
  <div>
    <h1 style="color: #1f2f3d; text-align: center;">5分钟构建任意溯源系统</h1>
    <p style="color: #5e6d82; text-align: center;">请对比字段填写，生成个性化的溯源系统。激活码仅可使用一次，</p>
    <p style="color: #5e6d82; text-align: center;">提交前请认真对比，生成后请尽快下载并备份，源码在服务器保留一周后删除。</p>
    <div style="text-align: center; margin-bottom: 20px;">
      <a href="https://www.bilibili.com/video/BV1Ar421H7TK" style="color: #409EFF; text-decoration: underline;">
        B站：使用教程
      </a>
    </div>
    <el-button type="text" @click="dialog2Visible = true">购买激活码</el-button>
    <el-dialog
      title="删除此页面"
      :visible.sync="dialog1Visible"
      width="30%"
      :before-close="handleClose"
    >
      <span>1. 在fabric-trace/application/web目录下，运行./rmbuildsyspage.sh</span>
      <br>
      <span style="display: block;margin-top: 20px;">2.重新启动前端，npm run dev</span>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialog1Visible = false">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog
      title="开发不易，感谢支持！"
      :visible.sync="dialog2Visible"
      width="30%"
      :before-close="handleClose"
    >
      <span>激活码售价：269元</span>

      <br>
      <span style="display: block; margin-top: 20px;">请加QQ群776873343联系群主购买 （补差价可包搭建）</span>
      <br>
      <span style="display: block; margin-top: 20px;">购买课程后可99元购买激活码</span>
      <br>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialog2Visible = false">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog
      title="构建成功！"
      :visible.sync="dialog3Visible"
      width="30%"
      :before-close="handleClose"
    >
      <span>下载地址：</span>
      <br>
      <span style="display: block;margin-top: 20px;">{{ downloadUrl }} </span>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialog3Visible = false">确 定</el-button>
      </span>
    </el-dialog>

    <div class="form-container">
      <el-form ref="form" :model="form" label-width="150px" class="form">

        <el-form-item v-for="(value, key) in form" :key="key" :label="key">
          <el-input v-model="form[key]" :placeholder="'请输入 ' + key + ' 值'" />
        </el-form-item>
        <div style="text-align: center; margin-top: 20px;">
          <el-button v-loading="loading" type="primary" element-loading-text="构建中，稍等1分钟" @click="submitForm">开始构建</el-button>
        </div>
      </el-form>
      <el-button type="text" style="display: block;margin-top: 20px;" @click="dialog1Visible = true">如何删除此页面？</el-button>
    </div>
    <div style="height: 30px;" />
  </div>
</template>

<script>
export default {
  name: 'Build',
  data() {
    return {
      loading: false,
      dialog1Visible: false,
      dialog2Visible: false,
      form: {
        parm1: '基于区块链的农产品溯源系统',
        parm2: '农产品信息',
        parm3: '农产品名称',
        parm4: '产地',
        parm5: '种植时间',
        parm6: '采摘时间',
        parm7: '种植户名称',
        parm8: '种植户',
        parm9: '工厂信息',
        parm10: '商品名称',
        parm11: '生产批次',
        parm12: '生产时间',
        parm13: '工厂名称与厂址',
        parm14: '工厂电话',
        parm15: '工厂',
        parm16: '物流轨迹信息',
        parm17: '姓名',
        parm18: '年龄',
        parm19: '联系电话',
        parm20: '车牌号',
        parm21: '运输记录',
        parm22: '物流司机',
        parm23: '商店信息',
        parm24: '入库时间',
        parm25: '销售时间',
        parm26: '商店名称',
        parm27: '商店位置',
        parm28: '商店电话',
        parm29: '商店',
        parm30: '消费者',
        parm31: '云服务器IP,例如：32.12.243.30/192.168.1.20',
        activatecode: '激活码'
      },
      downloadUrl: '',
      dialog3Visible: false
    }
  },
  methods: {
    submitForm() {
      this.loading = true
      const params = new URLSearchParams()
      for (const key in this.form) {
        params.append(key, this.form[key])
      }
      fetch('http://realcool.top:8088/activate', {
      // fetch('http://127.0.0.1:8088/activate', {

        method: 'POST',
        body: params,
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
      })
        .then(response => response.json())
        .then(data => {
          console.log('响应数据：', data)
          this.loading = false
          this.downloadUrl = data.msg
          // eslint-disable-next-line eqeqeq
          if (data.code == 0) {
            // this.$message.success('构建成功！' + data.msg);
            this.dialog3Visible = true
          } else {
            this.$message.error('构建失败：' + data.msg)
            console.log('构建失败：', data.msg)
          }
        })
        .catch(error => {
          this.$message.error('提交失败，请重试！')
          console.error('提交失败：' + error)
          this.loading = false
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.form-container {
  max-width: 700px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f9f9f9;
  border: 1px solid #ebebeb;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.el-form-item {
  margin-bottom: 15px;
}

.el-button {
  width: 100%;
}
</style>
