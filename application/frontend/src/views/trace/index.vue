<!--suppress HtmlUnknownTag-->
<!--suppress JSUnresolvedReference-->
<template>
  <div class='trace-container'>
    <div style='margin-top: 10px; margin-left: 10px'>
      <el-input v-model='input' placeholder='请输入溯源码查询' style='width: 500px; margin-right: 10px'/>
      <el-button type='primary' plain @click='ProductInfo'>查询</el-button>
    </div>
    <el-table :data='tracedata' v-if="tracedata.length!==1 || tracedata[0].traceabilityCode!==''">
      <el-table-column type='expand'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='props'>
          <el-form label-position='left' inline>
            <div class='container'>
              <div class='card'>
                <!--suppress HtmlDeprecatedAttribute-->
                <Timeline :value='generateEventsFromRow(props.row, props.$index)' align='alternate'>
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
                          <div class='two-column-layout'
                               v-if="slotProps.item.name === '批发商' && slotProps.item.id && slotProps.item.status !== '进行中'">
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
                          <div style='display: flex; flex-direction: column; align-items: flex-start'
                               v-if="slotProps.item.name === '批发商' && !slotProps.item.id && slotProps.item.status === '进行中'">
                            <el-button plain type='primary' @click='showLocation = true'>
                              获取定位
                            </el-button>
                            <img :src='locationImageUrl' alt='locationImage' v-show='showLocation' style='margin-top: 16px'>
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
                        <div v-if="slotProps.item.name === '消费者'">
                          <p>通过我们的溯源系统，您可以轻松查看每一瓶蜂蜜的来源，享受透明、安全的购物体验。</p>
                          <p>每一步都经过严格监控，从生产到销售，确保您购买的每一份产品都是高质量的。</p>
                          <p>我们承诺 100% 可追溯，确保每一位消费者都能享受健康、安全的产品。</p>
                          <div v-if="slotProps.item.color==='#92D050'">
                            <el-button type='success' plain @click='generateQRCode(props.row)'>
                              生成二维码分享溯源报告
                            </el-button>
                            <div v-if='qrCodeDataUrl'>
                              <img :src='qrCodeDataUrl' alt='qrCode' style='width: 500px; transform: translateY(10px)'/>
                            </div>
                          </div>
                          <img src='/images/logo.png' alt='logo' style='transform: translateY(16px); width: 250px'/>
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
          <span :class="getStatus(scope.row, 'beeFarm', 1, scope.$index).class">
            {{ getStatus(scope.row, 'beeFarm', 1, scope.$index).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='加工厂' prop='processingPlantInput.processingPlantTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, 'processingPlant', 2, scope.$index).class">
            {{ getStatus(scope.row, 'processingPlant', 2, scope.$index).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='批发商' prop='wholesalerInput.wholesalerTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, 'wholesaler', 3, scope.$index).class">
            {{ getStatus(scope.row, 'wholesaler', 3, scope.$index).status }}
          </span>
        </template>
      </el-table-column>
      <el-table-column label='零售商' prop='retailerInput.retailerTxid'>
        <!--suppress HtmlDeprecatedAttribute-->
        <template slot-scope='scope'>
          <span :class="getStatus(scope.row, 'retailer', 4, scope.$index).class">
            {{ getStatus(scope.row, 'retailer', 4, scope.$index).status }}
          </span>
        </template>
      </el-table-column>
    </el-table>
    <el-table v-else>
      <el-table-column type='expand'/>
      <el-table-column label='溯源码'/>
      <el-table-column label='养蜂场'/>
      <el-table-column label='加工厂'/>
      <el-table-column label='批发商'/>
      <el-table-column label='零售商'/>
    </el-table>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import Timeline from 'primevue/timeline'
import Card from 'primevue/card'
import {getAllProductInfo, getProductInfo, getProductList, ipfsDownload} from '@/api'
import QRCode from 'qrcode'

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
      input: '',
      qrCodeDataUrl: '',
      recall: false,
      randomIndexes: [],
      showLocation: false,
      locationImageUrl: 'https://restapi.amap.com/v3/staticmap?location=121.21,31.28&zoom=10&size=500*300&markers=mid,,A:121.21,31.28&key=ee95e52bf08006f63fd29bcfbcf21df0'
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ])
  },
  created() {
    this.recall = this.$route.query.recall === 'true'
    getProductList().then(res => {
      this.tracedata = JSON.parse(res.data).filter(item => item.traceabilityCode !== '')
      if (this.recall && this.tracedata.length >= 2) {
        this.randomIndexes = this.generateTwoUniqueRandomIndexes(this.tracedata.length)
      }
    })
  },
  mounted() {
    this.AllProductInfo()
  },
  watch: {
    '$route.query.recall'(newVal) {
      this.recall = newVal === 'true'
    }
  },
  methods: {
    generateTwoUniqueRandomIndexes(max) {
      const indexes = new Set()
      while (indexes.size < 2) {
        const randomIndex = Math.floor(Math.random() * max)
        indexes.add(randomIndex)
      }
      return Array.from(indexes)
    },
    generateQRCode(data) {
      const formatData = (obj, indent = 0) => {
        let formattedString = ''
        for (const key in obj) {
          if (obj.hasOwnProperty(key)) {
            const value = obj[key]
            const indentSpace = '    '.repeat(indent)
            if (typeof value === 'object' && value !== null) {
              formattedString += `${indentSpace}${key}:\n` + formatData(value, indent + 1)
            } else {
              formattedString += `${indentSpace}${key}: ${value}\n`
            }
          }
        }
        return formattedString
      }
      let dataString = 'Miel Link 溯源报告\n\n'
      const currentTime = new Date().toLocaleString()
      dataString += `${currentTime} 生成\n\n`
      dataString += formatData(data)
      QRCode.toDataURL(dataString).then(url => {
        this.qrCodeDataUrl = url
      }).catch(() => {
        this.$message.error('二维码生成失败，请检查网络连接情况。')
      })
    },
    generateEventsFromRow(row, index = null) {
      let events = []
      events.push({
        name: '养蜂场',
        icon: 'pi pi-table',
        status: this.getStatus(row, 'beeFarm', 1, index).status,
        color: this.getStatus(row, 'beeFarm', 1, index).color,
        image: this.getStatus(row, 'beeFarm', 1, index).image,
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
        icon: 'pi pi-home',
        status: this.getStatus(row, 'processingPlant', 2, index).status,
        color: this.getStatus(row, 'processingPlant', 2, index).color,
        image: this.getStatus(row, 'processingPlant', 2, index).image,
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
        status: this.getStatus(row, 'wholesaler', 3, index).status,
        color: this.getStatus(row, 'wholesaler', 3, index).color,
        image: this.getStatus(row, 'wholesaler', 3, index).image,
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
        icon: 'pi pi-shopping-bag',
        status: this.getStatus(row, 'retailer', 4, index).status,
        color: this.getStatus(row, 'retailer', 4, index).color,
        image: this.getStatus(row, 'retailer', 4, index).image,
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
        icon: 'pi pi-wallet',
        color: this.getStatus(row, 'retailer', 4, index).status === '进行中' ? '#FFC000' : this.getStatus(row, 'retailer', 4, index).color,
      })
      return events
    },
    getStatus(row, stage, index, _index = null) {
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
        if (this.recall && ((_index != null && this.randomIndexes.includes(_index)) || _index == null)) {
          return {status: '召回中', class: 'in-recall', color: '#EE8944', image: 'inrecall.png'}
        } else {
          return {status: '已完成', class: 'completed', color: '#92D050', image: 'completed.png'}
        }
      }
      if (index === firstMissingIndex) {
        if (this.recall && ((_index != null && this.randomIndexes.includes(_index)) || _index == null)) {
          return {status: '已终止', class: 'terminated', color: '#C00000', image: 'terminated.png'}
        } else {
          return {status: '进行中', class: 'in-progress', color: '#5B9BD5', image: 'inprogress.png'}
        }
      } else if (index < firstMissingIndex) {
        if (this.recall && ((_index != null && this.randomIndexes.includes(_index)) || _index == null)) {
          return {status: '召回中', class: 'in-recall', color: '#EE8944', image: 'inrecall.png'}
        } else {
          return {status: '已完成', class: 'completed', color: '#92D050', image: 'completed.png'}
        }
      } else {
        if (this.recall && ((_index != null && this.randomIndexes.includes(_index)) || _index == null)) {
          return {status: '已终止', class: 'terminated', color: '#C00000', image: 'terminated.png'}
        } else {
          return {status: '未开始', class: 'pending', color: '#FFC000', image: 'notstarted.png'}
        }
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
      if (this.input) {
        const formData = new FormData()
        formData.append('traceabilityCode', this.input)
        getProductInfo(formData).then(res => {
          if (res.code === 200) {
            this.tracedata = []
            this.tracedata[0] = JSON.parse(res.data)
          }
        })
      } else {
        this.AllProductInfo()
      }
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

.in-recall {
  background-color: #EE8944;
  padding: 2px 28px;
  display: inline-block;
  border-radius: 14px;
  font-size: 14px;
  color: white;
}

.terminated {
  background-color: #C00000;
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
