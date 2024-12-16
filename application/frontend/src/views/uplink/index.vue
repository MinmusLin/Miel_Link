<!--suppress JSUnresolvedReference-->
<!--suppress HtmlUnknownTag-->
<template>
  <div class='container'>
    <div class='header'>
      <div class='avatar-container'>
        <img class='avatar' :src='getAvatarImage(userType)' alt='头像'>
      </div>
      <div class='user-info'>
        <div class='user-column'>
          <span>当前用户</span>
          <div class='user-type'>{{ name }}</div>
        </div>
        <div class='user-column'>
          <span>用户角色</span>
          <div class='user-type'>{{ userType }}</div>
        </div>
        <div class='user-column'>
          <span>用户职责</span>
          <div v-if="userType === '养蜂场'" class='user-type'>蜜蜂养殖与蜂蜜生产</div>
          <div v-else-if="userType === '加工厂'" class='user-type'>蜂蜜加工与产品包装</div>
          <div v-else-if="userType === '批发商'" class='user-type'>蜂蜜批量采购与分销</div>
          <div v-else-if="userType === '零售商'" class='user-type'>蜂蜜零售与市场推广</div>
          <div v-else-if="userType === '消费者'" class='user-type'>购买和使用蜂蜜产品</div>
        </div>
      </div>
    </div>
    <div class='bee-form-container'>
      <form>
        <el-form ref='form' :model='tracedata' label-width='80px' size='medium' style=''>
          <el-form-item v-show="userType !== '养蜂场' && userType !== '消费者'" style='width: 500px'>
            <div>
              <div class='bee-label'>溯源码<span style='color: red; font-size: 16px'> * </span></div>
              <div style='text-align: left'>
                <el-input v-model='tracedata.traceabilityCode'/>
              </div>
            </div>
          </el-form-item>
          <div v-show="userType==='养蜂场'">
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>养蜂场名称<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.BeeFarmInput.BeeFarmName'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>养蜂场地点<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.BeeFarmInput.BeeFarmLocation'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>养蜂箱编号<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.BeeFarmInput.BeeBoxId'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>蜂蜜种类<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.BeeFarmInput.HoneyVariety'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>花卉种类<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.BeeFarmInput.FlowerVariety'/>
                </div>
              </div>
            </el-form-item>
          </div>
          <div v-show="userType==='加工厂'">
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>加工厂名称<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.ProcessingPlantInput.ProcessingPlantName'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>加工厂地点<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.ProcessingPlantInput.ProcessingPlantLocation'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>加工批次<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.ProcessingPlantInput.ProcessingBatchId'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>包装规格<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.ProcessingPlantInput.PackagingSpecification'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>保质期<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.ProcessingPlantInput.ShelfLife'/>
                </div>
              </div>
            </el-form-item>
          </div>
          <div v-show="userType==='批发商'">
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>仓库名称<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.WholesalerInput.WarehouseName'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>仓库地点<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.WholesalerInput.WarehouseLocation'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>进货批次<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.WholesalerInput.WholesalerBatchId'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>运输方式<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.WholesalerInput.TransportationMethod'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>交通方式<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.WholesalerInput.TransportMode'/>
                </div>
              </div>
            </el-form-item>
          </div>
          <div v-show="userType==='零售商'">
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>商店名称<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.RetailerInput.StoreName'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>商店地点<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.RetailerInput.StoreLocation'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>采购批次<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.RetailerInput.RetailerBatchId'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>销售渠道<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.RetailerInput.SalesChannel'/>
                </div>
              </div>
            </el-form-item>
            <el-form-item style='width: 500px'>
              <div>
                <div class='bee-label'>销售价格<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left'>
                  <el-input v-model='tracedata.RetailerInput.SalesPrice'/>
                </div>
              </div>
            </el-form-item>
          </div>
        </el-form>
        <div class='ipfs-container' v-show="userType !== '消费者'">
          <el-form>
            <el-form-item style='width: 420px; margin-left: 80px; margin-right: 80px'>
              <div>
                <div class='bee-label'>质检报告<span style='color: red; font-size: 16px'> * </span></div>
                <div style='text-align: left; position: relative'>
                  <el-upload ref='upload' action='#' :auto-upload='false'>
                    <el-button type='warning' plain>
                      上传质检报告
                    </el-button>
                  </el-upload>
                  <el-button type='danger'
                             class='recall-button'
                             v-show="userType !== '养蜂场'"
                             @click="goToPage('/trace?recall=true')">
                    紧急召回
                  </el-button>
                </div>
              </div>
            </el-form-item>
          </el-form>
        </div>
        <div class='consumer-container'>
          <span slot='footer' style='color: gray'>
           <el-button v-show="userType !== '消费者'"
                      type='warning'
                      @click='submitTracedata()'
                      style='width: 420px; margin-top: 30px'>
             提交
           </el-button>
          </span>
          <span v-show="userType === '消费者'" slot='footer' style='margin-top: 50px; font-size: 18px'>
            消费者没有信息录入权限
          </span>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import {uplink, ipfsUpload} from '@/api'

export default {
  name: 'Uplink',
  data() {
    return {
      tracedata: {
        traceabilityCode: '',
        ipfsFileCID: '',
        ipfsFileName: '',
        BeeFarmInput: {
          BeeFarmName: '',
          BeeFarmLocation: '',
          BeeBoxId: '',
          HoneyVariety: '',
          FlowerVariety: ''
        },
        ProcessingPlantInput: {
          ProcessingPlantName: '',
          ProcessingPlantLocation: '',
          ProcessingBatchId: '',
          PackagingSpecification: '',
          ShelfLife: ''
        },
        WholesalerInput: {
          WarehouseName: '',
          WarehouseLocation: '',
          WholesalerBatchId: '',
          TransportationMethod: '',
          TransportMode: ''
        },
        RetailerInput: {
          StoreName: '',
          StoreLocation: '',
          RetailerBatchId: '',
          SalesChannel: '',
          SalesPrice: ''
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
    submitTracedata() {
      /*
       * This section currently lacks input validation for user submissions, which compromises data security.
       * Additionally, there is no verification to ensure the traceability code exists, which may lead to potential
       * issues. The quality inspection report upload does not restrict users to a single file, but according to IPFS
       * requirements, only a single file should be uploaded. We aim to make the project more robust and complete. Feel
       * free to open an issue or submit a pull request! :)
       */
      if (this.$refs.upload._data.uploadFiles.length !== 0) {
        this.uploadIPFSFile().then(() => {
          this.tracedata.ipfsFileName = this.$refs.upload._data.uploadFiles[0].name
          const loading = this.$loading({
            lock: true,
            text: '数据上链中...',
            spinner: 'el-icon-loading',
            background: 'rgba(0, 0, 0, 0.7)'
          })
          const formData = new FormData()
          formData.append('traceabilityCode', this.tracedata.traceabilityCode)
          switch (this.userType) {
            case '养蜂场':
              formData.append('arg1', this.tracedata.BeeFarmInput.BeeFarmName)
              formData.append('arg2', this.tracedata.BeeFarmInput.BeeFarmLocation)
              formData.append('arg3', this.tracedata.BeeFarmInput.BeeBoxId)
              formData.append('arg4', this.tracedata.BeeFarmInput.HoneyVariety)
              formData.append('arg5', this.tracedata.BeeFarmInput.FlowerVariety)
              break
            case '加工厂':
              formData.append('arg1', this.tracedata.ProcessingPlantInput.ProcessingPlantName)
              formData.append('arg2', this.tracedata.ProcessingPlantInput.ProcessingPlantLocation)
              formData.append('arg3', this.tracedata.ProcessingPlantInput.ProcessingBatchId)
              formData.append('arg4', this.tracedata.ProcessingPlantInput.PackagingSpecification)
              formData.append('arg5', this.tracedata.ProcessingPlantInput.ShelfLife)
              break
            case '批发商':
              formData.append('arg1', this.tracedata.WholesalerInput.WarehouseName)
              formData.append('arg2', this.tracedata.WholesalerInput.WarehouseLocation)
              formData.append('arg3', this.tracedata.WholesalerInput.WholesalerBatchId)
              formData.append('arg4', this.tracedata.WholesalerInput.TransportationMethod)
              formData.append('arg5', this.tracedata.WholesalerInput.TransportMode)
              break
            case '零售商':
              formData.append('arg1', this.tracedata.RetailerInput.StoreName)
              formData.append('arg2', this.tracedata.RetailerInput.StoreLocation)
              formData.append('arg3', this.tracedata.RetailerInput.RetailerBatchId)
              formData.append('arg4', this.tracedata.RetailerInput.SalesChannel)
              formData.append('arg5', this.tracedata.RetailerInput.SalesPrice)
              break
          }
          formData.append('arg6', this.tracedata.ipfsFileCID)
          formData.append('arg7', this.tracedata.ipfsFileName)
          uplink(formData).then(res => {
            if (res.code === 200) {
              loading.close()
            } else {
              loading.close()
              this.$message.error('数据上链失败，请检查网络连接情况。')
            }
          }).catch(() => {
            this.$message.error('数据上链失败，请检查网络连接情况。')
          })
        }).catch(() => {
          this.$message.error('数据上链失败，请检查网络连接情况。')
        })
      }
    },
    goToPage(path) {
      this.$router.push(path)
    },
    uploadIPFSFile() {
      return new Promise((resolve, reject) => {
        const formData = new FormData()
        formData.append('file', this.$refs.upload._data.uploadFiles[0].raw)
        ipfsUpload(formData).then(res => {
          if (res.code === 200) {
            this.tracedata.ipfsFileCID = res.cid
            resolve()
          } else {
            reject('Upload IPFS file failed')
          }
        }).catch(() => {
          reject('Upload IPFS file failed')
        })
      })
    },
    getAvatarImage(userType) {
      switch (userType) {
        case '养蜂场':
          return '/images/apiary.png'
        case '加工厂':
          return '/images/factory.png'
        case '批发商':
          return '/images/wholesaler.png'
        case '零售商':
          return '/images/retailer.png'
        default:
          return '/images/consumer.png'
      }
    }
  }
}
</script>

<style scoped lang='scss'>
.container {
  width: 100%;
  height: calc(100vh - 50px);
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #FFFCF0;
  font-family: Arial, sans-serif;
  overflow-y: auto;
}

.header {
  width: 100%;
  background-color: #F8E99B;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 0;
  box-sizing: border-box;
}

.avatar-container {
  display: flex;
  justify-content: center;
  margin-bottom: 15px;
}

.avatar {
  width: 110px;
  height: 110px;
  border-radius: 50%;
  background-color: #D9D9D9;
  display: flex;
  justify-content: center;
  align-items: center;
  font-weight: bold;
  color: #555555;
  margin-bottom: 20px;
}

.user-info {
  width: 60%;
  display: flex;
  justify-content: space-between;
  padding: 0 30px;
  box-sizing: border-box;
  font-size: 24px;
  color: #333333;
}

.user-info .user-column {
  font-weight: bold;
  font-size: 18px;
}

.user-info .user-type {
  font-weight: bold;
  font-size: 36px;
  margin-top: 5px;
}

.el-form-item {
  display: block;
  text-align: left;
  margin-bottom: 20px;
}

.bee-form-container {
  display: flex;
  justify-content: center;
  padding: 60px 0;
}

.bee-label {
  width: 120px;
  text-align: left;
  font-size: 18px;
  margin-bottom: 5px;
}

::v-deep .el-form-item label {
  display: block;
  font-size: 18px;
  margin-bottom: 10px;
}

.el-input {
  font-size: 16px;
}

.ipfs-container {
  display: flex;
  justify-content: center;
}

.consumer-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.recall-button {
  position: absolute;
  top: 0;
  right: 0;
}
</style>
