<!--suppress HtmlUnknownTag-->
<!--suppress JSUnresolvedReference-->
<template>
  <div class='trace-container'>
    <div style="display: flex; justify-content: flex-end; margin-right:30px; align-items: center;">
      <el-input v-model='input' placeholder='请输入溯源码查询' style='width: 300px;'/>
      <el-button type='primary' plain @click='ProductInfo'>查询</el-button>
    </div>

    <!--<el-button type='success' plain @click='AllProductInfo'>获取所有产品信息</el-button>-->
    <el-table :data='tracedata' style='width: 100%'>
      <el-table-column type='expand'>
        <template slot-scope='props'>
          <el-form label-position='left' inline class='demo-table-expand'>
            <div class="container">

              <div class="card-container-left">
                <div>
                  <span class="trace-text">养蜂场信息</span>
                </div>
                <div class="info-line">
                  <el-form-item label="养蜂场名称">
                    <span>{{ props.row.beeFarmInput.beeFarmName }}</span>
                  </el-form-item>
                  <el-form-item label="养蜂场地点">
                    <span>{{ props.row.beeFarmInput.beeFarmLocation }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="养蜂箱编号">
                    <span>{{ props.row.beeFarmInput.beeBoxId }}</span>
                  </el-form-item>
                  <el-form-item label="蜂蜜种类">
                    <span>{{ props.row.beeFarmInput.honeyVariety }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="花卉种类">
                    <span>{{ props.row.beeFarmInput.flowerVariety }}</span>
                  </el-form-item>
                </div>
                <el-form-item v-show="props.row.beeFarmInput.beeFarmIPFSCID" label="附件">
                  <el-link type="primary" @click="GetIPFSFile(props.row.beeFarmInput.beeFarmIPFSCID, props.row.beeFarmInput.beeFarmIPFSFileName)">
                    {{ props.row.beeFarmInput.beeFarmIPFSCID }}
                  </el-link>
                </el-form-item>
                <el-form-item label="区块链交易 ID">
                  <span>{{ props.row.beeFarmInput.beeFarmTxid }}</span>
                </el-form-item>
                <el-form-item label="区块链交易时间">
                  <span>{{ props.row.beeFarmInput.beeFarmTimestamp }}</span>
                </el-form-item>
              </div>


              <div class="card-container-right">
                <div>
                  <span class="trace-text">加工厂信息</span>
                </div>
                <div class="info-line">
                  <el-form-item label="加工厂名称">
                    <span>{{ props.row.processingPlantInput.processingPlantName }}</span>
                  </el-form-item>
                  <el-form-item label="加工厂地点">
                    <span>{{ props.row.processingPlantInput.processingPlantLocation }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="加工批次">
                    <span>{{ props.row.processingPlantInput.processingBatchId }}</span>
                  </el-form-item>
                  <el-form-item label="包装规格">
                    <span>{{ props.row.processingPlantInput.packagingSpecification }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="保质期">
                    <span>{{ props.row.processingPlantInput.shelfLife }}</span>
                  </el-form-item>
                </div>
                <el-form-item v-show="props.row.processingPlantInput.processingPlantIPFSCID" label="附件">
                  <el-link type="primary" @click="GetIPFSFile(props.row.processingPlantInput.processingPlantIPFSCID, props.row.processingPlantInput.processingPlantIPFSFileName)">
                    {{ props.row.processingPlantInput.processingPlantIPFSCID }}
                  </el-link>
                </el-form-item>
                <el-form-item label="区块链交易 ID">
                  <span>{{ props.row.processingPlantInput.processingPlantTxid }}</span>
                </el-form-item>
                <el-form-item label="区块链交易时间">
                  <span>{{ props.row.processingPlantInput.processingPlantTimestamp }}</span>
                </el-form-item>
              </div>

              <div class="divider"></div>

              <div class="card-container-left">
                <div>
                  <span class="trace-text">批发商信息</span>
                </div>
                <div class="info-line">
                  <el-form-item label="仓库名称">
                    <span>{{ props.row.wholesalerInput.warehouseName }}</span>
                  </el-form-item>
                  <el-form-item label="仓库地点">
                    <span>{{ props.row.wholesalerInput.warehouseLocation }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="进货批次">
                    <span>{{ props.row.wholesalerInput.wholesalerBatchId }}</span>
                  </el-form-item>
                  <el-form-item label="运输方式">
                    <span>{{ props.row.wholesalerInput.transportationMethod }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="交通方式">
                    <span>{{ props.row.wholesalerInput.transportMode }}</span>
                  </el-form-item>
                </div>
                <el-form-item v-show="props.row.wholesalerInput.wholesalerIPFSCID" label="附件">
                  <el-link type="primary" @click="GetIPFSFile(props.row.wholesalerInput.wholesalerIPFSCID, props.row.wholesalerInput.wholesalerIPFSFileName)">
                    {{ props.row.wholesalerInput.wholesalerIPFSCID }}
                  </el-link>
                </el-form-item>
                <el-form-item label="区块链交易 ID">
                  <span>{{ props.row.wholesalerInput.wholesalerTxid }}</span>
                </el-form-item>
                <el-form-item label="区块链交易时间">
                  <span>{{ props.row.wholesalerInput.wholesalerTimestamp }}</span>
                </el-form-item>
              </div>

              <div class="card-container-right">
                <div>
                  <span class="trace-text">零售商信息</span>
                </div>
                <div class="info-line">
                  <el-form-item label="商店名称">
                    <span>{{ props.row.retailerInput.storeName }}</span>
                  </el-form-item>
                  <el-form-item label="商店地点">
                    <span>{{ props.row.retailerInput.storeLocation }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="采购批次">
                    <span>{{ props.row.retailerInput.retailerBatchId }}</span>
                  </el-form-item>
                  <el-form-item label="销售渠道">
                    <span>{{ props.row.retailerInput.salesChannel }}</span>
                  </el-form-item>
                </div>
                <div class="info-line">
                  <el-form-item label="销售价格">
                    <span>{{ props.row.retailerInput.salesPrice }}</span>
                  </el-form-item>
                </div>
                <el-form-item v-show="props.row.retailerInput.retailerIPFSCID" label="附件">
                  <el-link type="primary" @click="GetIPFSFile(props.row.retailerInput.retailerIPFSCID, props.row.retailerInput.retailerIPFSFileName)">
                    {{ props.row.retailerInput.retailerIPFSCID }}
                  </el-link>
                </el-form-item>
                <el-form-item label="区块链交易 ID">
                  <span>{{ props.row.retailerInput.retailerTxid }}</span>
                </el-form-item>
                <el-form-item label="区块链交易时间">
                  <span>{{ props.row.retailerInput.retailerTimestamp }}</span>
                </el-form-item>
              </div>
            </div>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label="溯源码" prop="traceabilityCode"/>

      <el-table-column label="养蜂场" prop="beeFarmInput.beeFarmTxid">
        <template slot-scope="scope">
    <span :class="getStatus(scope.row, 'beeFarm', 1).class">
      {{ getStatus(scope.row, 'beeFarm', 1).status }}
    </span>
        </template>
      </el-table-column>

      <el-table-column label="加工厂" prop="processingPlantInput.processingPlantTxid">
        <template slot-scope="scope">
    <span :class="getStatus(scope.row, 'processingPlant', 2).class">
      {{ getStatus(scope.row, 'processingPlant', 2).status }}
    </span>
        </template>
      </el-table-column>

      <el-table-column label="批发商" prop="wholesalerInput.wholesalerTxid">
        <template slot-scope="scope">
    <span :class="getStatus(scope.row, 'wholesaler', 3).class">
      {{ getStatus(scope.row, 'wholesaler', 3).status }}
    </span>
        </template>
      </el-table-column>

      <el-table-column label="零售商" prop="retailerInput.retailerTxid">
        <template slot-scope="scope">
    <span :class="getStatus(scope.row, 'retailer', 4).class">
      {{ getStatus(scope.row, 'retailer', 4).status }}
    </span>
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
      const stages = ["beeFarm", "processingPlant", "wholesaler", "retailer"];
      let firstMissingIndex = null;
      for (let i = 0; i < stages.length; i++) {
        const txid = row[`${stages[i]}Input`][`${stages[i]}Txid`];
        if (!txid) {
          firstMissingIndex = i + 1;
          break;
        }
      }

      if (firstMissingIndex === null) {
        return {status: "已完成", class: "completed"};
      }
      if (index === firstMissingIndex) {
        return {status: "进行中", class: "in-progress"};
      } else if (index < firstMissingIndex) {
        return {status: "已完成", class: "completed"};
      } else {
        return {status: "等待中", class: "pending"};
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

<style scoped lang='scss' >
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
}

.divider {
  height: 100%;
  width: 4px;
  border-left: 4px dashed #e0e0e0;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

.card-container-left {
  display: flex;
  flex-direction: column;
  padding: 16px;
  background: #f0f8ff;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  margin-bottom: -180px;
  width: 49%;
  margin-right: auto;
}

.card-container-right {
  display: flex;
  flex-direction: column;
  padding: 16px;
  background: #fffce8;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  margin-bottom: 5px;
  width: 49%;
  margin-left: auto;
}

.trace-text {
  font-weight: bold;
  margin-bottom: 12px;
}

.el-form-item {
  margin-bottom: 8px;
}

.el-form-item label {
  font-weight: 500;
}

.el-form-item span {
  display: block;
  font-size: 14px;
  color: #606266;
}
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
    font-size: 20px;
    line-height: 46px;
  }
}

.completed {
  background-color: #3d906c;
  padding: 4px 20px;
  border-radius: 20px;
  font-size: 14px;
  display: inline-block;
  color: #fff;
  margin-left: -10px;
}

.in-progress {
  background-color: #4475ff;
  padding: 4px 20px;
  border-radius: 20px;
  font-size: 14px;
  display: inline-block;
  color: #fff;
  margin-left: -10px;
}

.pending {
  background-color: #eed84c;
  padding: 4px 20px;
  border-radius: 20px;
  font-size: 14px;
  display: inline-block;
  color: #fff;
  margin-left: -10px;
}

.info-line {
  display: flex;
  justify-content: space-between;
}
</style>
