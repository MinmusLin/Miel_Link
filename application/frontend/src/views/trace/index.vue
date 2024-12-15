<!--suppress HtmlUnknownTag-->
<!--suppress JSUnresolvedReference-->
<template>
  <div class='trace-container'>
    <div>
      <el-input v-model='input' placeholder='请输入溯源码查询' style='width: 300px'/>
      <el-button type='primary' plain @click='ProductInfo'>查询</el-button>
    </div>
    <el-table :data='tracedata'>
      <el-table-column type='expand'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='props'>
          <el-form label-position='left' inline>
            <div class='container'>
              <div class='card'>
                <!--suppress HtmlDeprecatedAttribute-->
                <Timeline :value='generateEventsFromRow(props.row)' align='alternate'>
                  <template #marker='slotProps'>
                    <span class='custom-marker' :style='{ backgroundColor: slotProps.item.color }'>
                      <i :class='slotProps.item.icon' style='font-size: 22px'/>
                    </span>
                  </template>
                  <template #content='slotProps'>
                    <Card>
                      <template #title>
                        {{ slotProps.item.name }}
                      </template>
                      <template #subtitle>
                        {{ slotProps.item.time }}
                      </template>
                      <template #content>
                        <div v-if="slotProps.item.name === '养蜂场'">
                          <div>养蜂场名称: {{ slotProps.item.beeFarmName }}</div>
                          <div>养蜂场地点: {{ slotProps.item.beeFarmLocation }}</div>
                          <div>养蜂箱编号: {{ slotProps.item.beeBoxId }}</div>
                          <div>蜂蜜种类: {{ slotProps.item.honeyVariety }}</div>
                          <div>花卉种类: {{ slotProps.item.flowerVariety }}</div>
                          <div v-if='slotProps.item.fileid'>
                            附件:
                            <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                     type='primary'>
                              {{ slotProps.item.filename }}
                            </el-link>
                          </div>
                        </div>
                        <div v-if="slotProps.item.name === '加工厂'">
                          <div>加工厂名称: {{ slotProps.item.processingPlantName }}</div>
                          <div>加工厂地点: {{ slotProps.item.processingPlantLocation }}</div>
                          <div>加工批次: {{ slotProps.item.processingBatchId }}</div>
                          <div>包装规格: {{ slotProps.item.packagingSpecification }}</div>
                          <div>保质期: {{ slotProps.item.shelfLife }}</div>
                          <div v-if='slotProps.item.fileid'>
                            附件:
                            <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                     type='primary'>
                              {{ slotProps.item.filename }}
                            </el-link>
                          </div>
                        </div>
                        <div v-if="slotProps.item.name === '批发商'">
                          <div>仓库名称: {{ slotProps.item.warehouseName }}</div>
                          <div>仓库地点: {{ slotProps.item.warehouseLocation }}</div>
                          <div>进货批次: {{ slotProps.item.wholesalerBatchId }}</div>
                          <div>运输方式: {{ slotProps.item.transportationMethod }}</div>
                          <div>交通方式: {{ slotProps.item.transportMode }}</div>
                          <div v-if='slotProps.item.fileid'>
                            附件:
                            <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                     type='primary'>
                              {{ slotProps.item.filename }}
                            </el-link>
                          </div>
                        </div>
                        <div v-if="slotProps.item.name === '零售商'">
                          <div>商店名称: {{ slotProps.item['商店名称'] }}</div>
                          <div>商店地点: {{ slotProps.item['商店地点'] }}</div>
                          <div>采购批次: {{ slotProps.item['采购批次'] }}</div>
                          <div>销售渠道: {{ slotProps.item['销售渠道'] }}</div>
                          <div>销售价格: {{ slotProps.item['销售价格'] }}</div>
                          <div v-if='slotProps.item.fileid'>
                            附件:
                            <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                     type='primary'>
                              {{ slotProps.item.filename }}
                            </el-link>
                          </div>
                          <div v-if="slotProps.item.name === '消费者'">
                            <p>消费者</p>
                          </div>
                        </div>
                      </template>
                    </Card>
                  </template>
                </Timeline>
              </div>
            </div>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label='溯源码' prop='traceabilityCode'/>
      <el-table-column label='养蜂场' prop='beeFarmInput.beeFarmTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, ' beeFarm', 1).class">
            {{ getStatus(scope.row, 'beeFarm', 1).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='加工厂' prop='processingPlantInput.processingPlantTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, ' processingPlant', 2).class">
            {{ getStatus(scope.row, 'processingPlant', 2).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='批发商' prop='wholesalerInput.wholesalerTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, ' wholesaler', 3).class">
            {{ getStatus(scope.row, 'wholesaler', 3).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='零售商' prop='retailerInput.retailerTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, ' retailer', 4).class">
            {{ getStatus(scope.row, 'retailer', 4).status }}
          </span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import Timeline from 'primevue/timeline'
import Card from 'primevue/card'
import {getProductInfo, getProductList, getAllProductInfo, ipfsDownload} from '@/api'

export default {
  name: 'Trace',
  components: {
    Timeline,
    Card
  },
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
    generateEventsFromRow(row) {
      let events = []
      events.push({
        name: '养蜂场',
        icon: 'pi pi-building-columns',
        color: '#673AB7',
        time: row.beeFarmInput.beeFarmTimestamp || '',
        id: row.beeFarmInput.beeFarmTxid || '',
        beeFarmName: row.beeFarmInput.beeFarmName || '',
        beeFarmLocation: row.beeFarmInput.beeFarmLocation || '',
        beeBoxId: row.beeFarmInput.beeBoxId || '',
        honeyVariety: row.beeFarmInput.honeyVariety || '',
        flowerVariety: row.beeFarmInput.flowerVariety || '',
        fileid: row.beeFarmInput.beeFarmIPFSCID || '',
        filename: row.beeFarmInput.beeFarmIPFSFileName || ''
      })
      events.push({
        name: '加工厂',
        icon: 'pi pi-warehouse',
        color: '#673AB7',
        time: row.processingPlantInput.processingPlantTimestamp || '',
        id: row.processingPlantInput.processingPlantTxid || '',
        processingPlantName: row.processingPlantInput.processingPlantName || '',
        processingPlantLocation: row.processingPlantInput.processingPlantLocation || '',
        processingBatchId: row.processingPlantInput.processingBatchId || '',
        packagingSpecification: row.processingPlantInput.packagingSpecification || '',
        shelfLife: row.processingPlantInput.shelfLife || '',
        fileid: row.processingPlantInput.processingPlantIPFSCID || '',
        filename: row.processingPlantInput.processingPlantIPFSFileName || ''
      })
      events.push({
        name: '批发商',
        icon: 'pi pi-truck',
        color: '#673AB7',
        time: row.wholesalerInput.wholesalerTimestamp || '',
        id: row.wholesalerInput.wholesalerTxid || '',
        warehouseName: row.wholesalerInput.warehouseName || '',
        warehouseLocation: row.wholesalerInput.warehouseLocation || '',
        wholesalerBatchId: row.wholesalerInput.wholesalerBatchId || '',
        transportationMethod: row.wholesalerInput.transportationMethod || '',
        transportMode: row.wholesalerInput.transportMode || '',
        fileid: row.wholesalerInput.wholesalerIPFSCID || '',
        filename: row.wholesalerInput.wholesalerIPFSFileName || ''
      })
      events.push({
        name: '零售商',
        icon: 'pi pi-shop',
        color: '#673AB7',
        time: row.retailerInput.retailerTimestamp || '',
        id: row.retailerInput.retailerTxid || '',
        storeName: row.retailerInput.storeName || '',
        storeLocation: row.retailerInput.storeLocation || '',
        retailerBatchId: row.retailerInput.retailerBatchId || '',
        salesChannel: row.retailerInput.salesChannel || '',
        salesPrice: row.retailerInput.salesPrice || '',
        fileid: row.retailerInput.retailerIPFSCID || '',
        filename: row.retailerInput.retailerIPFSFileName || ''
      })
      events.push({
        name: '消费者',
        icon: 'pi pi-shopping-bag',
        color: '#673AB7',
      })
      return events
    },
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
        return {status: '已完成', class: 'completed'}
      }
      if (index === firstMissingIndex) {
        return {status: '进行中', class: 'in-progress'}
      } else if (index < firstMissingIndex) {
        return {status: '已完成', class: 'completed'}
      } else {
        return {status: '未开始', class: 'pending'}
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

<style scoped lang='scss'>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
}

.completed {
  background-color: #92D050;
  padding: 2px 28px;
  display: inline-block;
  border-radius: 14px;
  font-size: 14px;
  color: white;
}

.in-progress {
  background-color: #5B9BD5;
  padding: 2px 28px;
  display: inline-block;
  border-radius: 14px;
  font-size: 14px;
  color: white;
}

.pending {
  background-color: #FFC000;
  padding: 2px 28px;
  display: inline-block;
  border-radius: 14px;
  font-size: 14px;
  color: white;
}

.custom-marker {
  display: flex;
  width: 3rem;
  height: 3rem;
  align-items: center;
  justify-content: center;
  color: white;
  border-radius: 50%;
  z-index: 1;
}
</style>
