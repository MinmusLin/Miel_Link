<!--suppress HtmlUnknownTag-->
<!--suppress JSUnresolvedReference-->
<template>
  <div class='trace-container'>
    <el-input v-model='input' placeholder='请输入溯源码查询' style='width: 300px'/>
    <el-button type='primary' plain @click='ProductInfo'>查询</el-button>
    <el-button type='success' plain @click='AllProductInfo'>获取所有产品信息</el-button>
    <el-table :data='tracedata' style='width: 100%'>
      <el-table-column type='expand'>
        <template slot-scope='props'>
          <el-form label-position='left' inline class='demo-table-expand'>
            <div>
              <span class='trace-text'>养蜂场信息</span>
            </div>
            <el-form-item label='养蜂场名称'>
              <span>{{ props.row.beeFarmInput.beeFarmName }}</span>
            </el-form-item>
            <el-form-item label='养蜂场地点'>
              <span>{{ props.row.beeFarmInput.beeFarmLocation }}</span>
            </el-form-item>
            <el-form-item label='养蜂箱编号'>
              <span>{{ props.row.beeFarmInput.beeBoxId }}</span>
            </el-form-item>
            <el-form-item label='蜂蜜种类'>
              <span>{{ props.row.beeFarmInput.honeyVariety }}</span>
            </el-form-item>
            <el-form-item label='花卉种类'>
              <span>{{ props.row.beeFarmInput.flowerVariety }}</span>
            </el-form-item>
            <el-form-item v-show='props.row.beeFarmInput.beeFarmIPFSCID' label='附件'>
              <el-link type='primary'
                       @click='GetIPFSFile(props.row.beeFarmInput.beeFarmIPFSCID, props.row.beeFarmInput.beeFarmIPFSFileName)'>
                {{ props.row.beeFarmInput.beeFarmIPFSCID }}
              </el-link>
            </el-form-item>
            <el-form-item label='区块链交易 ID'>
              <span>{{ props.row.beeFarmInput.beeFarmTxid }}</span>
            </el-form-item>
            <el-form-item label='区块链交易时间'>
              <span>{{ props.row.beeFarmInput.beeFarmTimestamp }}</span>
            </el-form-item>
            <div>
              <span class='trace-text'>加工厂信息</span>
            </div>
            <el-form-item label='加工厂名称'>
              <span>{{ props.row.processingPlantInput.processingPlantName }}</span>
            </el-form-item>
            <el-form-item label='加工厂地点'>
              <span>{{ props.row.processingPlantInput.processingPlantLocation }}</span>
            </el-form-item>
            <el-form-item label='加工批次'>
              <span>{{ props.row.processingPlantInput.processingBatchId }}</span>
            </el-form-item>
            <el-form-item label='包装规格'>
              <span>{{ props.row.processingPlantInput.packagingSpecification }}</span>
            </el-form-item>
            <el-form-item label='保质期'>
              <span>{{ props.row.processingPlantInput.shelfLife }}</span>
            </el-form-item>
            <el-form-item v-show='props.row.processingPlantInput.processingPlantIPFSCID' label='附件'>
              <el-link type='primary'
                       @click='GetIPFSFile(props.row.processingPlantInput.processingPlantIPFSCID, props.row.processingPlantInput.processingPlantIPFSFileName)'>
                {{ props.row.processingPlantInput.processingPlantIPFSCID }}
              </el-link>
            </el-form-item>
            <el-form-item label='区块链交易 ID'>
              <span>{{ props.row.processingPlantInput.processingPlantTxid }}</span>
            </el-form-item>
            <el-form-item label='区块链交易时间'>
              <span>{{ props.row.processingPlantInput.processingPlantTimestamp }}</span>
            </el-form-item>
            <div>
              <span class='trace-text'>批发商信息</span>
            </div>
            <el-form-item label='仓库名称'>
              <span>{{ props.row.wholesalerInput.warehouseName }}</span>
            </el-form-item>
            <el-form-item label='仓库地点'>
              <span>{{ props.row.wholesalerInput.warehouseLocation }}</span>
            </el-form-item>
            <el-form-item label='进货批次'>
              <span>{{ props.row.wholesalerInput.wholesalerBatchId }}</span>
            </el-form-item>
            <el-form-item label='运输方式'>
              <span>{{ props.row.wholesalerInput.transportationMethod }}</span>
            </el-form-item>
            <el-form-item label='交通方式'>
              <span>{{ props.row.wholesalerInput.transportMode }}</span>
            </el-form-item>
            <el-form-item v-show='props.row.wholesalerInput.wholesalerIPFSCID' label='附件'>
              <el-link type='primary'
                       @click='GetIPFSFile(props.row.wholesalerInput.wholesalerIPFSCID, props.row.wholesalerInput.wholesalerIPFSFileName)'>
                {{ props.row.wholesalerInput.wholesalerIPFSCID }}
              </el-link>
            </el-form-item>
            <el-form-item label='区块链交易 ID'>
              <span>{{ props.row.wholesalerInput.wholesalerTxid }}</span>
            </el-form-item>
            <el-form-item label='区块链交易时间'>
              <span>{{ props.row.wholesalerInput.wholesalerTimestamp }}</span>
            </el-form-item>
            <div>
              <span class='trace-text'>零售商信息</span>
            </div>
            <el-form-item label='商店名称'>
              <span>{{ props.row.retailerInput.storeName }}</span>
            </el-form-item>
            <el-form-item label='商店地点'>
              <span>{{ props.row.retailerInput.storeLocation }}</span>
            </el-form-item>
            <el-form-item label='采购批次'>
              <span>{{ props.row.retailerInput.retailerBatchId }}</span>
            </el-form-item>
            <el-form-item label='销售渠道'>
              <span>{{ props.row.retailerInput.salesChannel }}</span>
            </el-form-item>
            <el-form-item label='销售价格'>
              <span>{{ props.row.retailerInput.salesPrice }}</span>
            </el-form-item>
            <el-form-item v-show='props.row.retailerInput.retailerIPFSCID' label='附件'>
              <el-link type='primary'
                       @click='GetIPFSFile(props.row.retailerInput.retailerIPFSCID, props.row.retailerInput.retailerIPFSFileName)'>
                {{ props.row.retailerInput.retailerIPFSCID }}
              </el-link>
            </el-form-item>
            <el-form-item label='区块链交易 ID'>
              <span>{{ props.row.retailerInput.retailerTxid }}</span>
            </el-form-item>
            <el-form-item label='区块链交易时间'>
              <span>{{ props.row.retailerInput.retailerTimestamp }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label='溯源码' prop='traceabilityCode'/>
      <el-table-column label='养蜂场' prop='beeFarmInput.beeFarmTxid'>
        <template slot-scope='scope'>
          <span>{{ getStatus(scope.row, 'beeFarm', 1) }}</span>
        </template>
      </el-table-column>
      <el-table-column label='加工厂' prop='processingPlantInput.processingPlantTxid'>
        <template slot-scope='scope'>
          <span>{{ getStatus(scope.row, 'processingPlant', 2) }}</span>
        </template>
      </el-table-column>
      <el-table-column label='批发商' prop='wholesalerInput.wholesalerTxid'>
        <template slot-scope='scope'>
          <span>{{ getStatus(scope.row, 'wholesaler', 3) }}</span>
        </template>
      </el-table-column>
      <el-table-column label='零售商' prop='retailerInput.retailerTxid'>
        <template slot-scope='scope'>
          <span>{{ getStatus(scope.row, 'retailer', 4) }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import {getProductInfo, getProductList, getAllProductInfo, ipfsDownload} from '@/api'

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
    getProductList().then(res => {
      this.tracedata = JSON.parse(res.data).filter(item => item.traceabilityCode !== '')
    })
  },
  mounted() {
    this.AllProductInfo()
  },
  methods: {
    getStatus(row, stage, index) {
      const stages = ['beeFarm', 'processingPlant', 'wholesaler', 'retailer']
      let firstMissingIndex = null
      for (let i = 0; i < stages.length; i++) {
        const txid = row[`${stages[i]}Input`][`${stages[i]}Txid`]
        if (!txid) {
          firstMissingIndex = i + 1
          break
        }
      }
      if (firstMissingIndex === null) {
        return '已完成'
      }
      if (index === firstMissingIndex) {
        return '进行中'
      } else if (index < firstMissingIndex) {
        return '已完成'
      } else {
        return '等待中'
      }
    },
    GetIPFSFile(cid, filename) {
      const formData = new FormData()
      formData.append('cid', cid)
      formData.append('filename', filename)
      ipfsDownload(formData).then(res => {
        window.open('http://43.156.142.179:9090/download?filename=' + res.data)
      })
    },
    AllProductInfo() {
      getAllProductInfo().then(res => {
        this.tracedata = JSON.parse(res.data).filter(item => item.traceabilityCode !== '')
      })
    },
    ProductInfo() {
      const formData = new FormData()
      formData.append('traceabilityCode', this.input)
      getProductInfo(formData).then(res => {
        if (res.code === 200) {
          this.tracedata = []
          this.tracedata[0] = JSON.parse(res.data)
        }
      })
    }
  }
}
</script>

<style lang='scss' scoped>
.demo-table-expand {
  font-size: 0;
}

.demo-table-expand label {
  width: 90px;
  color: #99A9BF;
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
