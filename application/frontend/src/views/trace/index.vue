<!--suppress HtmlUnknownTag-->
<template>
  <div class="trace-container">
    <el-input v-model="input" placeholder="请输入溯源码查询" style="width: 300px;margin-right: 15px;" />
    <el-button type="primary" plain @click="ProductInfo"> 查询 </el-button>
    <el-button type="success" plain @click="AllProductInfo"> 获取所有农产品信息 </el-button>
    <el-table
      :data="tracedata"
      style="width: 100%"
    >
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <div><span class="trace-text" style="color: #67C23A;">农产品信息</span></div>
            <el-form-item label="农产品名称：">
              <span>{{ props.row.beeFarmInput.beeFarmName }}</span>
            </el-form-item>
            <el-form-item label="产地：">
              <span>{{ props.row.beeFarmInput.beeFarmLocation }}</span>
            </el-form-item>
            <el-form-item label="种植时间：">
              <span>{{ props.row.beeFarmInput.beeBoxId }}</span>
            </el-form-item>
            <el-form-item label="采摘时间：">
              <span>{{ props.row.beeFarmInput.honeyVariety }}</span>
            </el-form-item>
            <el-form-item label="养蜂场名称：">
              <span>{{ props.row.beeFarmInput.flowerVariety }}</span>
            </el-form-item>
            <el-form-item v-show="props.row.beeFarmInput.fa_ipfscid" label="附件IPFSCID：">
              <el-link type="success" @click="GetIPFSFile(props.row.beeFarmInput.fa_ipfscid,props.row.beeFarmInput.fa_ipfsfilename)">{{ props.row.beeFarmInput.fa_ipfscid }}</el-link>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.beeFarmInput.beeFarmTxid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.beeFarmInput.beeFarmTimestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #409EFF;">加工厂信息</span></div>
            <el-form-item label="加工厂名称：">
              <span>{{ props.row.processingPlantInput.processingPlantName }}</span>
            </el-form-item>
            <el-form-item label="加工厂地点：">
              <span>{{ props.row.processingPlantInput.processingPlantLocation }}</span>
            </el-form-item>
            <el-form-item label="加工批次：">
              <span>{{ props.row.processingPlantInput.processingBatchId }}</span>
            </el-form-item>
            <el-form-item label="包装规格：">
              <span>{{ props.row.processingPlantInput.packagingSpecification }}</span>
            </el-form-item>
            <el-form-item label="保质期：">
              <span>{{ props.row.processingPlantInput.shelfLife }}</span>
            </el-form-item>
            <el-form-item v-show="props.row.processingPlantInput.fac_ipfscid" label="附件IPFSCID：">
              <el-link type="success" @click="GetIPFSFile(props.row.processingPlantInput.fac_ipfscid,props.row.processingPlantInput.fac_ipfsfilename)">{{ props.row.beeFarmInput.fac_ipfscid }}</el-link>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.processingPlantInput.processingPlantTxid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.processingPlantInput.processingPlantTimestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #E6A23C;">物流轨迹信息</span></div>
            <el-form-item label="仓库名称：">
              <span>{{ props.row.wholesalerInput.warehouseName }}</span>
            </el-form-item>
            <el-form-item label="仓库地点：">
              <span>{{ props.row.wholesalerInput.warehouseLocation }}</span>
            </el-form-item>
            <el-form-item label="进货批次：">
              <span>{{ props.row.wholesalerInput.wholesalerBatchId }}</span>
            </el-form-item>
            <el-form-item label="运输方式：">
              <span>{{ props.row.wholesalerInput.transportationMethod }}</span>
            </el-form-item>
            <el-form-item label="交通方式：">
              <span>{{ props.row.wholesalerInput.transportMode }}</span>
            </el-form-item>
            <el-form-item v-show="props.row.wholesalerInput.dr_ipfscid" label="附件IPFSCID：">
              <el-link type="success" @click="GetIPFSFile(props.row.wholesalerInput.dr_ipfscid,props.row.wholesalerInput.dr_ipfsfilename)">{{ props.row.wholesalerInput.dr_ipfscid }}</el-link>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.wholesalerInput.wholesalerTxid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.wholesalerInput.wholesalerTimestamp }}</span>
            </el-form-item>
            <div><span class="trace-text" style="color: #909399;">零售商信息</span></div>
            <el-form-item label="商店名称：">
              <span>{{ props.row.retailerInput.storeName }}</span>
            </el-form-item>
            <el-form-item label="商店地点：">
              <span>{{ props.row.retailerInput.storeLocation }}</span>
            </el-form-item>
            <el-form-item label="采购批次：">
              <span>{{ props.row.retailerInput.retailerBatchId }}</span>
            </el-form-item>
            <el-form-item label="销售渠道：">
              <span>{{ props.row.retailerInput.salesChannel }}</span>
            </el-form-item>
            <el-form-item label="销售价格：">
              <span>{{ props.row.retailerInput.salesPrice }}</span>
            </el-form-item>
            <el-form-item v-show="props.row.retailerInput.sh_ipfscid" label="附件IPFSCID：">
              <el-link type="success" @click="GetIPFSFile(props.row.retailerInput.sh_ipfscid,props.row.retailerInput.sh_ipfsfilename)">{{ props.row.retailerInput.sh_ipfscid }}</el-link>
            </el-form-item>
            <el-form-item label="区块链交易ID：">
              <span>{{ props.row.retailerInput.retailerTxid }}</span>
            </el-form-item>
            <el-form-item label="区块链交易时间：">
              <span>{{ props.row.retailerInput.retailerTimestamp }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        label="溯源码"
        prop="traceabilityCode"
      />
      <el-table-column
        label="农产品名称"
        prop="beeFarmInput.beeFarmName"
      />
      <el-table-column
        label="农产品采摘时间"
        prop="beeFarmInput.honeyVariety"
      />
    </el-table>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getProductInfo, getProductList, getAllProductInfo, getProductHistory, ipfsDownload } from '@/api/trace'

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
  methods: {
    GetIPFSFile(cid, filename) {
      var formData = new FormData()
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
      var formData = new FormData()
      formData.append('traceabilityCode', this.input)
      getProductInfo(formData).then(res => {
        if (res.code === 200) {
          this.tracedata = []
          this.tracedata[0] = JSON.parse(res.data)
        } else {
          this.$message.error('查询失败')
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
