<template>
  <div class="config-editor">
    <el-alert
      title="此功能为购课用户专享，可使用一次"
      type="warning"
      :closable="false"
      style="margin-bottom: 20px;"
    >
      <template slot="default">
        <div style="line-height: 1.8;">
          <p style="margin: 0;">
            购买课程获得完整功能权限：
            <a
              href="https://www.bilibili.com/cheese/play/ss15923"
              target="_blank"
              style="color: #409eff; text-decoration: underline; font-weight: 500;"
            >
              https://www.bilibili.com/cheese/play/ss15923
            </a>
          </p>
        </div>
      </template>
    </el-alert>

    <el-alert
      title="如何删除此页面？"
      type="info"
      :closable="false"
      style="margin-bottom: 20px;"
    >
      <template slot="default">
        <div style="line-height: 1.8;">
          <p style="margin: 0 0 8px 0;">
            <i class="el-icon-video-camera" />
            详细视频教程：
            <a
              href="https://www.bilibili.com/cheese/play/ep2325425"
              target="_blank"
              style="color: #409eff; text-decoration: underline; font-weight: 500;"
            >
              点击查看删除配置管理页面教程
            </a>
          </p>
        </div>
      </template>
    </el-alert>

    <div class="config-container">
      <el-card class="config-card compact-card">
        <div slot="header" class="card-header">
          <i class="el-icon-setting" /> <span>系统基础配置</span>
        </div>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form label-width="90px" class="compact-form">
              <el-form-item label="系统标题">
                <el-input v-model="z1.system.title" placeholder="请输入系统标题" size="small" />
              </el-form-item>
            </el-form>
          </el-col>
          <el-col :span="12">
            <div class="background-preview">
              <div class="preview-box" :style="{ backgroundImage: z2() }">
                <span v-if="!z1.system.loginBackground" class="no-bg-text">默认背景</span>
              </div>
              <div class="upload-actions">
                <input
                  ref="fileInput"
                  type="file"
                  accept="image/jpeg,image/jpg,image/png,image/gif,image/webp"
                  style="display: none"
                  @change="z3"
                >
                <el-button size="mini" type="primary" :loading="z4" @click="$refs.fileInput.click()">
                  <i class="el-icon-picture" /> 选择图片
                </el-button>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-card>

      <el-row :gutter="20">
        <el-col :span="12">
          <el-card class="config-card compact-card">
            <div slot="header" class="card-header">
              <i class="el-icon-user" /> <span>角色名称</span>
            </div>
            <el-form label-width="90px" class="compact-form">
              <el-form-item v-for="z5 in z6" :key="z5" :label="z5">
                <el-input v-model="z1.roles[z5]" size="small" />
              </el-form-item>
            </el-form>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card class="config-card compact-card">
            <div slot="header" class="card-header">
              <i class="el-icon-s-operation" /> <span>溯源查询中节点标题</span>
            </div>
            <el-form label-width="90px" class="compact-form">
              <el-form-item v-for="z5 in z7" :key="'stage-' + z5" :label="z5">
                <el-input v-model="z1.stageTitles[z5]" size="small" />
              </el-form-item>
            </el-form>
          </el-card>
        </el-col>
      </el-row>

      <el-card class="config-card compact-card">
        <div slot="header" class="card-header">
          <i class="el-icon-edit-outline" /> <span>各角色溯源信息配置</span>
          <el-radio-group v-model="z8" size="mini" class="role-tabs">
            <el-radio-button label="farmer">种植户</el-radio-button>
            <el-radio-button label="factory">工厂</el-radio-button>
            <el-radio-button label="driver">运输司机</el-radio-button>
            <el-radio-button label="shop">商店</el-radio-button>
            <el-radio-button label="consumer">消费者</el-radio-button>
          </el-radio-group>
        </div>
        <el-form label-width="140px" class="compact-form inline-form">
          <template v-if="z1.parameters[z8] && z1.parameters[z8].length > 0">
            <el-form-item
              v-for="(z9, z10) in z1.parameters[z8]"
              :key="z10"
              :label="z9.key"
            >
              <el-input v-model="z9.label" size="small" style="width: 300px;" />
            </el-form-item>
          </template>
          <div v-else class="no-params-tip">
            <i class="el-icon-info" /> 消费者角色无需配置参数标签
          </div>
        </el-form>
      </el-card>

      <div class="save-section">
        <div v-if="!z11" class="save-actions">
          <el-button
            class="save-btn success-button"
            size="medium"
            @click="z12"
          >
            <i class="el-icon-document-checked" /> 保存配置
          </el-button>
        </div>

        <div v-if="z11" class="saved-container">
          <div class="saved-header">
            <i class="el-icon-circle-check" />
            <span class="saved-text">配置已保存</span>
            <el-tag size="small" type="success">{{ z13 }}</el-tag>
          </div>

          <div class="saved-actions">
            <el-button class="restore-button" size="medium" plain :loading="z14" @click="z15">
              <i class="el-icon-refresh-left" /> 恢复初始配置
            </el-button>
          </div>
        </div>
      </div>
    </div>

    <el-dialog
      title="身份验证"
      :visible.sync="z16"
      width="450px"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-alert
        v-if="z17"
        title="配置保存功能已锁定"
        type="error"
        :closable="false"
        style="margin-bottom: 15px;"
      >
        <template slot="default">
          <div>由于多次验证失败，功能已锁定</div>
          <div style="margin-top: 8px;">解锁时间: {{ z18 }}</div>
          <div style="margin-top: 8px; font-size: 12px; color: #909399;">
            剩余: {{ z19 }}
          </div>
        </template>
      </el-alert>

      <el-form
        v-if="!z17"
        ref="z20"
        :model="z21"
        :rules="z22"
        label-width="100px"
      >
        <el-form-item label="UID" prop="uid">
          <el-input
            v-model="z21.uid"
            placeholder="请输入UID (6位或以上数字)"
            :disabled="z23"
          />
        </el-form-item>

        <el-form-item label="购课微信" prop="wechat">
          <el-input
            v-model="z21.wechat"
            placeholder="请输入微信号 (6-20位字母数字下划线减号)"
            :disabled="z23"
          />
        </el-form-item>

        <el-alert
          v-if="z24 > 0"
          :title="`验证失败 ${z24} 次，还剩 ${3 - z24} 次机会`"
          type="warning"
          :closable="false"
          style="margin-bottom: 10px;"
        />
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button :disabled="z23" @click="z16 = false">
          {{ z17 ? '关闭' : '取消' }}
        </el-button>
        <el-button
          v-if="!z17"
          type="primary"
          :loading="z23"
          @click="z25"
        >
          验证
        </el-button>
        <el-button
          v-if="z17"
          type="danger"
          size="small"
          @click="z26"
        >
          管理员重置
        </el-button>
      </span>
    </el-dialog>

    <el-dialog
      title="配置完成"
      :visible.sync="z27"
      width="500px"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <div class="success-dialog-content">
        <i class="el-icon-success" style="font-size: 60px; color: #67c23a; margin-bottom: 20px;" />
        <h3 style="margin: 0 0 15px 0;">配置已保存完成</h3>
      </div>

      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="z27 = false">我已了解</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { uploadLoginBackground, saveConfig } from '@/api/config'

export default {
  name: 'ConfigEditor',
  data() {
    return {
      z8: 'farmer',
      z1: {
        system: {
          title: '基于区块链的农产品溯源系统',
          loginBackground: ''
        },
        roles: {
          farmer: '种植户',
          factory: '工厂',
          driver: '运输司机',
          shop: '商店',
          consumer: '消费者'
        },
        stageTitles: {
          farmer: '农产品信息',
          factory: '工厂信息',
          driver: '物流轨迹信息',
          shop: '商店信息'
        },
        parameters: {
          farmer: [
            { key: 'Fa_fruitName', label: '农产品名称' },
            { key: 'Fa_origin', label: '产地' },
            { key: 'Fa_plantTime', label: '种植时间' },
            { key: 'Fa_pickingTime', label: '采摘时间' },
            { key: 'Fa_farmerName', label: '种植户姓名' }
          ],
          factory: [
            { key: 'Fac_productName', label: '商品名称' },
            { key: 'Fac_processingTime', label: '加工时间' },
            { key: 'Fac_factoryName', label: '工厂名称' },
            { key: 'Fac_factoryLocation', label: '工厂地址' },
            { key: 'Fac_processingMethod', label: '加工方式' }
          ],
          driver: [
            { key: 'Dr_name', label: '司机姓名' },
            { key: 'Dr_vehicleNumber', label: '车牌号' },
            { key: 'Dr_transportTime', label: '运输时间' },
            { key: 'Dr_temperature', label: '运输温度' },
            { key: 'Dr_destination', label: '目的地' }
          ],
          shop: [
            { key: 'Sh_shopName', label: '商店名称' },
            { key: 'Sh_shopAddress', label: '商店地址' },
            { key: 'Sh_saleTime', label: '上架时间' },
            { key: 'Sh_price', label: '价格' },
            { key: 'Sh_contact', label: '联系方式' }
          ],
          consumer: []
        }
      },
      z14: false,
      z4: false,
      z30: false,
      z6: ['farmer', 'factory', 'driver', 'shop', 'consumer'],
      z7: ['farmer', 'factory', 'driver', 'shop'],
      z16: false,
      z23: false,
      z21: {
        uid: '',
        wechat: ''
      },
      z22: {
        uid: [
          { required: true, message: '请输入UID', trigger: 'blur' }
        ],
        wechat: [
          { required: true, message: '请输入购课微信', trigger: 'blur' }
        ]
      },
      z24: 0,
      z17: false,
      z34: null,
      z35: '',
      z27: false
    }
  },
  computed: {
    z18() {
      const z36 = localStorage.getItem('config_auth_locked_until')
      if (!z36) return ''
      return new Date(parseInt(z36)).toLocaleString('zh-CN')
    },
    z19() {
      const z36 = localStorage.getItem('config_auth_locked_until')
      if (!z36) return ''
      const z37 = parseInt(z36) - Date.now()
      if (z37 <= 0) return '已解锁'
      const z38 = Math.floor(z37 / (60 * 60 * 1000))
      const z39 = Math.floor((z37 % (60 * 60 * 1000)) / (60 * 1000))
      return `${z38}小时${z39}分钟`
    },
    z11() {
      const z40 = localStorage.getItem('config_saved_once')
      return !!z40
    },
    z13() {
      const z40 = localStorage.getItem('config_saved_once')
      if (!z40) return ''
      return new Date(parseInt(z40)).toLocaleString('zh-CN')
    }
  },
  watch: {
    z1: {
      handler(z41) {
        localStorage.setItem('config_temp_full', JSON.stringify(z41))
      },
      deep: true
    }
  },
  mounted() {
    const z40 = localStorage.getItem('config_saved_once')
    if (z40) {
      this.z30 = true
    }

    const z42 = localStorage.getItem('config_auth_locked_until')
    if (z42 && Date.now() < parseInt(z42)) {
      this.z17 = true
      this.z43()
    }

    const z44 = parseInt(localStorage.getItem('config_auth_failed_count') || '0')
    this.z24 = z44

    const z45 = localStorage.getItem('config_saved_uid')
    if (z45) {
      this.z35 = z45
    }

    const z46 = localStorage.getItem('config_temp_full')
    if (z46) {
      try {
        const z47 = JSON.parse(z46)
        Object.assign(this.z1, z47)
      } catch (error) {
        console.error('恢复临时配置失败:', error)
      }
    }

    const z48 = localStorage.getItem('config_temp_loginBackground')
    if (z48 && !z46) {
      this.z1.system.loginBackground = z48
    }

    const z49 = localStorage.getItem('config_just_saved')
    if (z49 === 'true') {
      localStorage.removeItem('config_just_saved')
      setTimeout(() => {
        this.z27 = true
      }, 500)
    }
  },
  beforeDestroy() {
    if (this.z34) {
      clearInterval(this.z34)
    }
  },
  methods: {
    z12() {
      this.$confirm('此功能需要购买课程后才能使用', '提示', {
        confirmButtonText: '前往购买',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        window.open('https://www.bilibili.com/cheese/play/ss15923', '_blank')
      }).catch(() => {})
    },

    z3(z50) {
      const z51 = z50.target.files[0]
      if (!z51) return

      const z52 = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp']
      if (!z52.includes(z51.type)) {
        this.$message.error('只支持 JPG、PNG、GIF、WEBP 格式的图片')
        this.$refs.fileInput.value = ''
        return
      }

      const z53 = 10 * 1024 * 1024
      if (z51.size > z53) {
        this.$message.error('图片大小不能超过 10MB')
        this.$refs.fileInput.value = ''
        return
      }

      this.z4 = true
      uploadLoginBackground(z51).then(res => {
        if (res.code === 200) {
          this.z1.system.loginBackground = res.data.url
          localStorage.setItem('config_temp_full', JSON.stringify(this.z1))
          this.$message.success('背景图上传成功！请保存配置。')
          this.$refs.fileInput.value = ''
        }
      }).catch(error => {
        this.$message.error('上传失败: ' + (error.message || '未知错误'))
      }).finally(() => {
        this.z4 = false
      })
    },

    z2() {
      const z54 = this.z1.system.loginBackground
      if (!z54) return 'none'

      if (z54.startsWith('http') || z54.startsWith('data:')) {
        return `url("${z54}")`
      }

      return `url("/${z54}")`
    },

    async z15() {
      this.$confirm('确定要恢复初始配置吗？这将覆盖所有当前配置且不可恢复！', '警告', {
        confirmButtonText: '确定恢复',
        cancelButtonText: '取消',
        type: 'warning',
        distinguishCancelAndClose: true
      }).then(async() => {
        this.z14 = true

        try {
          this.z1 = {
            system: {
              title: '基于区块链的农产品溯源系统',
              loginBackground: ''
            },
            roles: {
              farmer: '种植户',
              factory: '工厂',
              driver: '运输司机',
              shop: '商店',
              consumer: '消费者'
            },
            stageTitles: {
              farmer: '农产品信息',
              factory: '工厂信息',
              driver: '物流轨迹信息',
              shop: '商店信息'
            },
            parameters: {
              farmer: [
                { key: 'Fa_fruitName', label: '农产品名称' },
                { key: 'Fa_origin', label: '产地' },
                { key: 'Fa_plantTime', label: '种植时间' },
                { key: 'Fa_pickingTime', label: '采摘时间' },
                { key: 'Fa_farmerName', label: '种植户姓名' }
              ],
              factory: [
                { key: 'Fac_productName', label: '商品名称' },
                { key: 'Fac_processingTime', label: '加工时间' },
                { key: 'Fac_factoryName', label: '工厂名称' },
                { key: 'Fac_factoryLocation', label: '工厂地址' },
                { key: 'Fac_processingMethod', label: '加工方式' }
              ],
              driver: [
                { key: 'Dr_name', label: '司机姓名' },
                { key: 'Dr_vehicleNumber', label: '车牌号' },
                { key: 'Dr_transportTime', label: '运输时间' },
                { key: 'Dr_temperature', label: '运输温度' },
                { key: 'Dr_destination', label: '目的地' }
              ],
              shop: [
                { key: 'Sh_shopName', label: '商店名称' },
                { key: 'Sh_shopAddress', label: '商店地址' },
                { key: 'Sh_saleTime', label: '上架时间' },
                { key: 'Sh_price', label: '价格' },
                { key: 'Sh_contact', label: '联系方式' }
              ],
              consumer: []
            }
          }

          const res = await saveConfig(this.z1)

          if (res.code === 200) {
            localStorage.setItem('config_saved_once', Date.now().toString())

            this.$message.success('已恢复初始配置！页面即将刷新...')

            setTimeout(() => {
              location.reload()
            }, 1500)
          } else {
            this.$message.error('保存配置失败: ' + (res.message || '未知错误'))
          }
        } catch (error) {
          console.error('恢复配置失败:', error)
          this.$message.error('恢复配置失败: ' + (error.message || '未知错误'))
        } finally {
          this.z14 = false
        }
      }).catch(() => {})
    },

    async z25() {
      this.$refs.z20.validate(async(z55) => {
        if (!z55) {
          return false
        }

        this.z23 = true

        try {
          await new Promise(resolve => setTimeout(resolve, 800))

          const z56 = true

          if (z56) {
            this.$message.success('验证通过！')

            const z57 = this.z21.uid

            this.z26()
            this.z16 = false

            this.z21.uid = z57

            await this.z58()
          } else {
            this.z59()
          }
        } catch (error) {
          this.$message.error('验证过程出错')
          console.error(error)
        } finally {
          this.z23 = false
        }
      })
    },

    z59() {
      const z60 = parseInt(localStorage.getItem('config_auth_failed_count') || '0')
      const z61 = z60 + 1
      localStorage.setItem('config_auth_failed_count', z61.toString())
      this.z24 = z61

      if (z61 >= 3) {
        const z62 = Date.now() + (24 * 60 * 60 * 1000)
        localStorage.setItem('config_auth_locked_until', z62.toString())
        this.z17 = true
        this.z43()

        this.$message.error('验证失败次数过多，功能已锁定24小时')
      } else {
        this.$message.error(`验证失败，还剩 ${3 - z61} 次机会`)
      }
    },

    z26() {
      localStorage.removeItem('config_auth_failed_count')
      localStorage.removeItem('config_auth_locked_until')
      this.z24 = 0
      this.z17 = false

      if (this.z34) {
        clearInterval(this.z34)
        this.z34 = null
      }

      this.z21.uid = ''
      this.z21.wechat = ''
      if (this.$refs.z20) {
        this.$refs.z20.clearValidate()
      }
    },

    z43() {
      if (this.z34) {
        clearInterval(this.z34)
      }

      this.z34 = setInterval(() => {
        const z63 = localStorage.getItem('config_auth_locked_until')
        if (!z63 || Date.now() >= parseInt(z63)) {
          this.z26()
        }
      }, 1000)
    },

    async z58() {
      const z64 = true

      try {
        const res = await saveConfig(this.z1)

        if (res.code === 200) {
          localStorage.setItem('config_saved_once', Date.now().toString())
          localStorage.setItem('config_saved_uid', this.z21.uid)
          localStorage.setItem('config_just_saved', 'true')
          localStorage.removeItem('config_temp_loginBackground')
          localStorage.removeItem('config_temp_full')
          this.z30 = true

          this.$message.success('配置保存成功！页面即将刷新...')

          setTimeout(() => {
            location.reload()
          }, 2500)
        } else {
          this.$message.error('保存失败: ' + (res.message || '未知错误'))
        }
      } catch (error) {
        console.error('保存配置失败:', error)
        this.$message.error('保存配置失败: ' + (error.message || '未知错误'))
      } finally {
        z64 = false
      }
    }
  }
}
</script>

<style scoped>
.config-editor { padding: 15px; background: #f5f7fa; min-height: calc(100vh - 84px); }
.config-container { max-width: 1400px; margin: 0 auto; }
.config-card { margin-bottom: 15px; }
.compact-card { margin-bottom: 15px; }
.compact-card ::v-deep .el-card__header { padding: 12px 15px; background: #fafafa; }
.compact-card ::v-deep .el-card__body { padding: 15px; }
.card-header { display: flex; align-items: center; justify-content: flex-start; font-weight: 500; }
.card-header i { margin-right: 6px; color: #409eff; }
.compact-form { margin-bottom: -8px; }
.compact-form ::v-deep .el-form-item { margin-bottom: 12px; }
.inline-form ::v-deep .el-form-item { display: inline-block; width: 48%; margin-right: 2%; margin-bottom: 10px; }
.role-tabs { float: right; margin-top: 8px; }
.card-header i { margin-right: 8px; }
.card-header span { margin-right: 40px; }
.role-tabs ::v-deep .el-radio-button:first-child { margin-left: 0; }
.save-section { padding: 20px; background: white; border-radius: 4px; margin-top: 10px; box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1); }
.save-actions { display: flex; justify-content: center; align-items: center; gap: 15px; }
.save-btn { min-width: 150px; font-size: 15px; padding: 12px 30px; }
.success-button { background: #67c23a !important; border-color: #67c23a !important; color: #ffffff !important; font-weight: 500; }
.success-button:hover { background: #85ce61 !important; border-color: #85ce61 !important; color: #ffffff !important; }
.restore-button { background: #ffffff !important; border-color: #F56C6C !important; color: #F56C6C !important; font-weight: 500; }
.restore-button:hover { background: #FEF0F0 !important; border-color: #F56C6C !important; color: #F56C6C !important; }
.saved-container { display: flex; flex-direction: column; gap: 15px; }
.saved-header { display: flex; align-items: center; justify-content: center; gap: 10px; padding: 15px; background: #67c23a; border-radius: 8px; color: white; }
.saved-header i { font-size: 24px; }
.saved-text { font-size: 16px; font-weight: 500; }
.saved-actions { display: flex; justify-content: center; align-items: center; gap: 15px; flex-wrap: wrap; }
.background-preview { display: flex; flex-direction: column; gap: 10px; }
.preview-box { width: 100%; height: 90px; background-size: cover; background-position: center; border: 1px dashed #dcdfe6; border-radius: 4px; display: flex; align-items: center; justify-content: center; background-color: #f5f7fa; }
.no-bg-text { color: #909399; font-size: 12px; }
.upload-actions { display: flex; gap: 8px; justify-content: center; }
.no-params-tip { padding: 20px; text-align: center; color: #909399; font-size: 14px; }
.no-params-tip i { margin-right: 5px; }
.success-dialog-content { text-align: center; padding: 20px 0; }
</style>
