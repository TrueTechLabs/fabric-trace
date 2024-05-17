<template>
  <div class="uplink-container">
    <div style="color:#909399;margin-bottom: 30px">
      当前用户：{{ name }};
      用户角色: {{ userType }}
    </div>
    <div>
      <el-form ref="form" :model="tracedata" label-width="80px" size="mini" style="">
        <el-form-item v-show="userType!='种植户'&userType!='消费者'" label="溯源码:" style="width: 300px" label-width="120px">
          <el-input v-model="tracedata.traceability_code" />
        </el-form-item>

        <div v-show="userType=='种植户'">
          <el-form-item label="农产品名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_fruitName" />
          </el-form-item>
          <el-form-item label="产地:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_origin" />
          </el-form-item>
          <el-form-item label="种植时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_plantTime" />
          </el-form-item>
          <el-form-item label="采摘时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_pickingTime" />
          </el-form-item>
          <el-form-item label="种植户名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Farmer_input.Fa_farmerName" />
          </el-form-item>
        </div>
        <div v-show="userType=='工厂'">
          <el-form-item label="商品名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_productName" />
          </el-form-item>
          <el-form-item label="生产批次:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_productionbatch" />
          </el-form-item>
          <el-form-item label="生产时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_productionTime" />
          </el-form-item>
          <el-form-item label="工厂名称与厂址:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_factoryName" />
          </el-form-item>
          <el-form-item label="联系电话:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Factory_input.Fac_contactNumber" />
          </el-form-item>
        </div>
        <div v-show="userType=='运输司机'">
          <el-form-item label="姓名:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_name" />
          </el-form-item>
          <el-form-item label="年龄:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_age" />
          </el-form-item>
          <el-form-item label="联系电话:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_phone" />
          </el-form-item>
          <el-form-item label="车牌号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_carNumber" />
          </el-form-item>
          <el-form-item label="运输工具:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Driver_input.Dr_transport" />
          </el-form-item>
        </div>
        <div v-show="userType=='商店'">
          <el-form-item label="存入时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_storeTime" />
          </el-form-item>
          <el-form-item label="销售时间:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_sellTime" />
          </el-form-item>
          <el-form-item label="商店名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_shopName" />
          </el-form-item>
          <el-form-item label="商店位置:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_shopAddress" />
          </el-form-item>
          <el-form-item label="商店电话:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.Shop_input.Sh_shopPhone" />
          </el-form-item>
        </div>
      </el-form>
      <span slot="footer" style="color: gray;" class="dialog-footer">
        <el-button v-show="userType != '消费者'" type="primary" plain style="margin-left: 220px;" @click="submittracedata()">提 交</el-button>
      </span>
      <span v-show="userType == '消费者'" slot="footer" style="color: gray;" class="dialog-footer">
        消费者没有权限录入！请使用溯源功能!
      </span>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { uplink } from '@/api/trace'

export default {
  name: 'Uplink',
  data() {
    return {
      tracedata: {
        traceability_code: '',
        Farmer_input: {
          Fa_fruitName: '',
          Fa_origin: '',
          Fa_plantTime: '',
          Fa_pickingTime: '',
          Fa_farmerName: ''
        },
        Factory_input: {
          Fac_productName: '',
          Fac_productionbatch: '',
          Fac_productionTime: '',
          Fac_factoryName: '',
          Fac_contactNumber: ''
        },
        Driver_input: {
          Dr_name: '',
          Dr_age: '',
          Dr_phone: '',
          Dr_carNumber: '',
          Dr_transport: ''
        },
        Shop_input: {
          Sh_storeTime: '',
          Sh_sellTime: '',
          Sh_shopName: '',
          Sh_shopAddress: '',
          Sh_shopPhone: ''
        }
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  methods: {
    submittracedata() {
      console.log(this.tracedata)
      const loading = this.$loading({
        lock: true,
        text: '数据上链中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      var formData = new FormData()
      formData.append('traceability_code', this.tracedata.traceability_code)
      // 根据不同的用户给arg1、arg2、arg3..赋值,
      switch (this.userType) {
        case '种植户':
          formData.append('arg1', this.tracedata.Farmer_input.Fa_fruitName)
          formData.append('arg2', this.tracedata.Farmer_input.Fa_origin)
          formData.append('arg3', this.tracedata.Farmer_input.Fa_plantTime)
          formData.append('arg4', this.tracedata.Farmer_input.Fa_pickingTime)
          formData.append('arg5', this.tracedata.Farmer_input.Fa_farmerName)
          break
        case '工厂':
          formData.append('arg1', this.tracedata.Factory_input.Fac_productName)
          formData.append('arg2', this.tracedata.Factory_input.Fac_productionbatch)
          formData.append('arg3', this.tracedata.Factory_input.Fac_productionTime)
          formData.append('arg4', this.tracedata.Factory_input.Fac_factoryName)
          formData.append('arg5', this.tracedata.Factory_input.Fac_contactNumber)
          break
        case '运输司机':
          formData.append('arg1', this.tracedata.Driver_input.Dr_name)
          formData.append('arg2', this.tracedata.Driver_input.Dr_age)
          formData.append('arg3', this.tracedata.Driver_input.Dr_phone)
          formData.append('arg4', this.tracedata.Driver_input.Dr_carNumber)
          formData.append('arg5', this.tracedata.Driver_input.Dr_transport)
          break
        case '商店':
          formData.append('arg1', this.tracedata.Shop_input.Sh_storeTime)
          formData.append('arg2', this.tracedata.Shop_input.Sh_sellTime)
          formData.append('arg3', this.tracedata.Shop_input.Sh_shopName)
          formData.append('arg4', this.tracedata.Shop_input.Sh_shopAddress)
          formData.append('arg5', this.tracedata.Shop_input.Sh_shopPhone)
          break
      }
      uplink(formData).then(res => {
        if (res.code === 200) {
          loading.close()
          this.$message({
            message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceability_code,
            type: 'success'
          })
        } else {
          loading.close()
          this.$message({
            message: '上链失败',
            type: 'error'
          })
        }
      }).catch(err => {
        loading.close()
        console.log(err)
      })
    }
  }
}

</script>

<style lang="scss" scoped>
.uplink {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
