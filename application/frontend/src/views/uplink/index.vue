<!--suppress HtmlUnknownTag-->
<template>
  <div class="uplink-container">
    <div style="color:#909399;margin-bottom: 30px">
      当前用户：{{ name }};
      用户角色: {{ userType }}
    </div>
    <div>
      <el-form ref="form" :model="tracedata" label-width="80px" size="mini" style="">
        <el-form-item v-show="userType!=='养蜂场'&userType!=='消费者'" label="溯源码:" style="width: 300px" label-width="120px">
          <el-input v-model="tracedata.traceabilityCode" />
        </el-form-item>
        <div v-show="userType==='养蜂场'">
          <el-form-item label="养蜂场名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.BeeFarmInput.BeeFarmName" />
          </el-form-item>
          <el-form-item label="养蜂场地点:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.BeeFarmInput.BeeFarmLocation" />
          </el-form-item>
          <el-form-item label="养蜂箱编号:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.BeeFarmInput.BeeBoxId" />
          </el-form-item>
          <el-form-item label="蜂蜜种类:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.BeeFarmInput.HoneyVariety" />
          </el-form-item>
          <el-form-item label="花卉种类:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.BeeFarmInput.FlowerVariety" />
          </el-form-item>
        </div>
        <div v-show="userType==='加工厂'">
          <el-form-item label="加工厂名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.ProcessingPlantInput.ProcessingPlantName" />
          </el-form-item>
          <el-form-item label="加工厂地点:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.ProcessingPlantInput.ProcessingPlantLocation" />
          </el-form-item>
          <el-form-item label="加工批次:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.ProcessingPlantInput.ProcessingBatchId" />
          </el-form-item>
          <el-form-item label="包装规格:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.ProcessingPlantInput.PackagingSpecification" />
          </el-form-item>
          <el-form-item label="保质期:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.ProcessingPlantInput.ShelfLife" />
          </el-form-item>
        </div>
        <div v-show="userType==='批发商'">
          <el-form-item label="仓库名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.WholesalerInput.WarehouseName" />
          </el-form-item>
          <el-form-item label="仓库地点:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.WholesalerInput.WarehouseLocation" />
          </el-form-item>
          <el-form-item label="进货批次:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.WholesalerInput.WholesalerBatchId" />
          </el-form-item>
          <el-form-item label="运输方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.WholesalerInput.TransportationMethod" />
          </el-form-item>
          <el-form-item label="交通方式:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.WholesalerInput.TransportMode" />
          </el-form-item>
        </div>
        <div v-show="userType==='零售商'">
          <el-form-item label="商店名称:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.RetailerInput.StoreName" />
          </el-form-item>
          <el-form-item label="商店地点:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.RetailerInput.StoreLocation" />
          </el-form-item>
          <el-form-item label="采购批次:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.RetailerInput.RetailerBatchId" />
          </el-form-item>
          <el-form-item label="销售渠道:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.RetailerInput.SalesChannel" />
          </el-form-item>
          <el-form-item label="销售价格:" style="width: 300px" label-width="120px">
            <el-input v-model="tracedata.RetailerInput.SalesPrice" />
          </el-form-item>
        </div>
      </el-form>
      <el-form>
        <el-form-item label="IPFS数据:" style="width: 300px" label-width="120px">
          <el-upload
            ref="upload"
            action="#"
            :auto-upload="false"
            class="upload-demo"
          >
            <el-button size="small" type="primary">点击上传</el-button>
            <div slot="tip" class="el-upload__tip">提示：将上传至IPFS网络</div>
          </el-upload>
        </el-form-item>
      </el-form>
      <span slot="footer" style="color: gray;" class="dialog-footer">
        <el-button v-show="userType !== '消费者'" type="primary" plain style="margin-left: 220px;" @click="submittracedata()">提 交</el-button>
      </span>
      <span v-show="userType === '消费者'" slot="footer" style="color: gray;" class="dialog-footer">
        消费者没有权限录入！请使用溯源功能!
      </span>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { uplink, ipfsUpload } from '@/api/trace'

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
    uploadIPFSFile() {
      return new Promise((resolve, reject) => {
        var formData = new FormData()
        formData.append('file', this.$refs.upload._data.uploadFiles[0].raw)
        ipfsUpload(formData).then(res => {
          if (res.code === 200) {
            this.tracedata.ipfsFileCID = res.cid
            resolve()
          } else {
            reject('IPFS文件上传失败')
          }
        }).catch(err => {
          reject('IPFS文件上传失败')
        })
      })
    },

    submittracedata() {
      if (this.$refs.upload._data.uploadFiles.length !== 0) {
        this.uploadIPFSFile().then(() => {
          this.tracedata.ipfsFileName = this.$refs.upload._data.uploadFiles[0].name

          const loading = this.$loading({
            lock: true,
            text: '数据上链中...',
            spinner: 'el-icon-loading',
            background: 'rgba(0, 0, 0, 0.7)'
          })

          var formData = new FormData()
          formData.append('traceabilityCode', this.tracedata.traceabilityCode)
          // 根据不同的用户给arg1、arg2、arg3..赋值,
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
              this.$message({
                message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceabilityCode,
                type: 'success'
              })
            } else {
              loading.close()
              this.$message({
                message: '上链失败',
                type: 'error'
              })
            }
          }).catch(() => {
          })
        }).catch(error => {
          this.$message({
            message: error,
            type: 'error'
          })
        })
      }
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
