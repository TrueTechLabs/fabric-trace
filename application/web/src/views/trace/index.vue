<template>
  <div class="trace-container">
    <el-input v-model="input" placeholder="请输入溯源码查询" style="width: 300px;margin-right: 15px;" />
    <el-button type="primary" plain @click="FruitInfo"> 查询 </el-button>
    <el-button type="success" plain @click="AllFruitInfo"> 获取所有农产品信息 </el-button>
    <el-table
      :data="tracedata"
      style="width: 100%"
    >
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <div><span class="trace-text" style="color: #67C23A;">农产品信息</span></div>
            <el-form-item label="农产品名称：">
              <span>{{ props.row.farmer_input.fa_fruitName }}</span>
            </el-form-item>
            <el-form-item label="产地：">
              <span>{{ props.row.farmer_input.fa_origin }}</span>
            </el-form-item>
            <el-form-item label="种植时间：">
              <span>{{ props.row.farmer_input.fa_plantTime }}</span>
            </el-form-item>
            <el-form-item label="采摘时间：">
              <span>{{ props.row.farmer_input.fa_pickingTime }}</span>
            </el-form-item>
            <el-form-item label="种植户名称：">
              <span>{{ props.row.farmer_input.fa_farmerName }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.farmer_input.fa_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.farmer_input.fa_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #409EFF;">工厂信息</span></div>
            <el-form-item label="商品名称：">
              <span>{{ props.row.factory_input.fac_productName }}</span>
            </el-form-item>
            <el-form-item label="生产批次：">
              <span>{{ props.row.factory_input.fac_productionbatch }}</span>
            </el-form-item>
            <el-form-item label="生产时间：">
              <span>{{ props.row.factory_input.fac_productionTime }}</span>
            </el-form-item>
            <el-form-item label="工厂名称与厂址：">
              <span>{{ props.row.factory_input.fac_factoryName }}</span>
            </el-form-item>
            <el-form-item label="联系电话：">
              <span>{{ props.row.factory_input.fac_contactNumber }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.factory_input.fac_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.factory_input.fac_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #E6A23C;">物流轨迹信息</span></div>
            <el-form-item label="姓名：">
              <span>{{ props.row.driver_input.dr_name }}</span>
            </el-form-item>
            <el-form-item label="年龄：">
              <span>{{ props.row.driver_input.dr_age }}</span>
            </el-form-item>
            <el-form-item label="联系电话：">
              <span>{{ props.row.driver_input.dr_phone }}</span>
            </el-form-item>
            <el-form-item label="车牌号：">
              <span>{{ props.row.driver_input.dr_carNumber }}</span>
            </el-form-item>
            <el-form-item label="运输记录：">
              <span>{{ props.row.driver_input.dr_transport }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.driver_input.dr_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.driver_input.dr_timestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #909399;">商店信息</span></div>
            <el-form-item label="入库时间：">
              <span>{{ props.row.shop_input.sh_storeTime }}</span>
            </el-form-item>
            <el-form-item label="销售时间：">
              <span>{{ props.row.shop_input.sh_sellTime }}</span>
            </el-form-item>
            <el-form-item label="商店名称：">
              <span>{{ props.row.shop_input.sh_shopName }}</span>
            </el-form-item>
            <el-form-item label="商店位置：">
              <span>{{ props.row.shop_input.sh_shopAddress }}</span>
            </el-form-item>
            <el-form-item label="商店电话：">
              <span>{{ props.row.shop_input.sh_shopPhone }}</span>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.shop_input.sh_txid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.shop_input.sh_timestamp }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        label="溯源码"
        prop="traceability_code"
      />
      <el-table-column
        label="农产品名称"
        prop="farmer_input.fa_fruitName"
      />
      <el-table-column
        label="农产品采摘时间"
        prop="farmer_input.fa_pickingTime"
      />
    </el-table>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getFruitInfo, getFruitList, getAllFruitInfo, getFruitHistory } from '@/api/trace'

export default {
  name: 'Trace',
  data() {
    return {
      tracedata: [],
      loading: false,
      input: ''
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  created() {
    getFruitList().then(res => {
      this.tracedata = JSON.parse(res.data)
    })
  },
  methods: {
    AllFruitInfo() {
      getAllFruitInfo().then(res => {
        this.tracedata = JSON.parse(res.data)
      })
    },
    FruitHistory() {
      getFruitHistory().then(res => {
        // console.log(res)
      })
    },
    FruitInfo() {
      var formData = new FormData()
      formData.append('traceability_code', this.input)
      getFruitInfo(formData).then(res => {
        if (res.code === 200) {
          // console.log(res)
          this.tracedata = []
          this.tracedata[0] = JSON.parse(res.data)
          return
        } else {
          this.$message.error(res.message)
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>

.demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
.trace {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
