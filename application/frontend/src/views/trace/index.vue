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
                        <span>{{ slotProps.item.name }}</span>
                      </template>
                      <template #subtitle>
                        <span v-if="slotProps.item.name === '消费者'"/>
                        <span v-else-if="slotProps.item.status === '已完成'">{{ slotProps.item.time }}</span>
                        <span v-else>此环节{{ slotProps.item.status }}</span>
                      </template>
                      <template #content>
                        <div v-if="slotProps.item.name !== '消费者'">
                          <div v-if="slotProps.item.name === '养蜂场' && slotProps.item.id" class='two-column-layout'>
                            <div class='item'>
                              <span class='item-label'>养蜂场名称</span>
                              <span>{{ slotProps.item.beeFarmName }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>养蜂场地点</span>
                              <span>{{ slotProps.item.beeFarmLocation }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>养蜂箱编号</span>
                              <span>{{ slotProps.item.beeBoxId }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>蜂蜜种类</span>
                              <span>{{ slotProps.item.honeyVariety }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>花卉种类</span>
                              <span>{{ slotProps.item.flowerVariety }}</span>
                            </div>
                            <div v-if='slotProps.item.fileid' class='item'>
                              <span class='item-label'>质检报告</span>
                              <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                       type='primary'>
                                {{ slotProps.item.filename }}
                              </el-link>
                            </div>
                          </div>
                          <div v-if="slotProps.item.name === '加工厂' && slotProps.item.id" class='two-column-layout'>
                            <div class='item'>
                              <span class='item-label'>加工厂名称</span>
                              <span>{{ slotProps.item.processingPlantName }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>加工厂地点</span>
                              <span>{{ slotProps.item.processingPlantLocation }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>加工批次</span>
                              <span>{{ slotProps.item.processingBatchId }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>包装规格</span>
                              <span>{{ slotProps.item.packagingSpecification }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>保质期</span>
                              <span>{{ slotProps.item.shelfLife }}</span>
                            </div>
                            <div v-if='slotProps.item.fileid' class='item'>
                              <span class='item-label'>质检报告</span>
                              <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                       type='primary'>
                                {{ slotProps.item.filename }}
                              </el-link>
                            </div>
                          </div>
                          <div v-if="slotProps.item.name === '批发商' && slotProps.item.id" class='two-column-layout'>
                            <div class='item'>
                              <span class='item-label'>仓库名称</span>
                              <span>{{ slotProps.item.warehouseName }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>仓库地点</span>
                              <span>{{ slotProps.item.warehouseLocation }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>进货批次</span>
                              <span>{{ slotProps.item.wholesalerBatchId }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>运输方式</span>
                              <span>{{ slotProps.item.transportationMethod }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>交通方式</span>
                              <span>{{ slotProps.item.transportMode }}</span>
                            </div>
                            <div v-if='slotProps.item.fileid' class='item'>
                              <span class='item-label'>质检报告</span>
                              <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                       type='primary'>
                                {{ slotProps.item.filename }}
                              </el-link>
                            </div>
                          </div>
                          <div v-if="slotProps.item.name === '零售商' && slotProps.item.id" class='two-column-layout'>
                            <div class='item'>
                              <span class='item-label'>商店名称</span>
                              <span>{{ slotProps.item.storeName }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>商店地点</span>
                              <span>{{ slotProps.item.storeLocation }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>采购批次</span>
                              <span>{{ slotProps.item.retailerBatchId }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>销售渠道</span>
                              <span>{{ slotProps.item.salesChannel }}</span>
                            </div>
                            <div class='item'>
                              <span class='item-label'>销售价格</span>
                              <span>{{ slotProps.item.salesPrice }}</span>
                            </div>
                            <div v-if='slotProps.item.fileid' class='item'>
                              <span class='item-label'>质检报告</span>
                              <el-link @click='GetIPFSFile(slotProps.item.fileid, slotProps.item.filename)'
                                       type='primary'>
                                {{ slotProps.item.filename }}
                              </el-link>
                            </div>
                          </div>
                          <!--
                              This section is intended to fix the issue of the PrimeVue Card component in the PrimeVue
                              Timeline component not displaying with the expected width. The current solution is not
                              elegant, and we hope to resolve this issue in a more elegant way. Feel free to open an
                              Issue or submit a Pull Request! :)
                          -->
                          <p style='line-height: 0; color: transparent; margin-bottom: -16px'>
                            test test test test test test test test test test test test test test test test test test
                            test test test test test test test test test test test test test test test test test test
                            test test test test test test test test test test test test test test
                          </p>
                          <img :src='`/images/${slotProps.item.image}`'
                               alt='status'
                               width='150px'
                               style='transform: translateY(20px)'/>
                        </div>
                        <div v-if="slotProps.item.name === '消费者'" class='two-column-layout'>
                          <div class='item'>
                            <span class='item-label'>消费者</span>
                            <span>消费者</span>
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
          <span :class="getStatus(scope.row, 'beeFarm', 1).class">
            {{ getStatus(scope.row, 'beeFarm', 1).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='加工厂' prop='processingPlantInput.processingPlantTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, 'processingPlant', 2).class">
            {{ getStatus(scope.row, 'processingPlant', 2).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='批发商' prop='wholesalerInput.wholesalerTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, 'wholesaler', 3).class">
            {{ getStatus(scope.row, 'wholesaler', 3).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='零售商' prop='retailerInput.retailerTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
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
        status: this.getStatus(row, 'beeFarm', 1).status,
        color: this.getStatus(row, 'beeFarm', 1).color,
        image: this.getStatus(row, 'beeFarm', 1).image,
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
        status: this.getStatus(row, 'processingPlant', 2).status,
        color: this.getStatus(row, 'processingPlant', 2).color,
        image: this.getStatus(row, 'processingPlant', 2).image,
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
        status: this.getStatus(row, 'wholesaler', 3).status,
        color: this.getStatus(row, 'wholesaler', 3).color,
        image: this.getStatus(row, 'wholesaler', 3).image,
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
        status: this.getStatus(row, 'retailer', 4).status,
        color: this.getStatus(row, 'retailer', 4).color,
        image: this.getStatus(row, 'retailer', 4).image,
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
        color: this.getStatus(row, 'retailer', 4).status === '进行中' ? '#FFC000' : this.getStatus(row, 'retailer', 4).color,
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
        return {status: '已完成', class: 'completed', color: '#92D050', image: 'completed.png'}
      }
      if (index === firstMissingIndex) {
        return {status: '进行中', class: 'in-progress', color: '#5B9BD5', image: 'inprogress.png'}
      } else if (index < firstMissingIndex) {
        return {status: '已完成', class: 'completed', color: '#92D050', image: 'completed.png'}
      } else {
        return {status: '未开始', class: 'pending', color: '#FFC000', image: 'notstarted.png'}
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

.two-column-layout {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
}

.item {
  display: flex;
  flex-basis: 48%;
}

.item-label {
  font-weight: bold;
  margin-right: 10px;
}
</style>
