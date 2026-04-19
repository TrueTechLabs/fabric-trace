<template>
  <el-card :class="['trace-info-card', type]" shadow="hover">
    <div slot="header" class="card-header">
      <i :class="icon" />
      <span>{{ title }}</span>
    </div>
    <div class="card-body">
      <slot>
        <field-display
          v-for="field in fields"
          :key="field.key"
          :label="field.label"
          :value="getFieldValue(field)"
          :link="field.link"
        />
        <field-display
          v-if="image"
          label="相关图片"
          :link="true"
        >
          <template #default>
            <a :href="imageUrl" target="_blank">
              <el-image
                style="width: 100px; height: 100px;"
                :src="imageUrl"
                fit="cover"
              />
            </a>
          </template>
        </field-display>
      </slot>
    </div>
    <div v-if="txid" class="tx-info">
      <div class="tx-field">
        <span class="tx-label">交易ID：</span>
        <span class="tx-value">{{ txid }}</span>
      </div>
      <div class="tx-field">
        <span class="tx-label">交易时间：</span>
        <span class="tx-value">{{ timestamp }}</span>
      </div>
    </div>
  </el-card>
</template>

<script>
import FieldDisplay from './FieldDisplay.vue'

export default {
  name: 'TraceInfoCard',
  components: {
    FieldDisplay
  },
  props: {
    title: {
      type: String,
      required: true
    },
    type: {
      type: String,
      default: 'farmer',
      validator: value => ['farmer', 'factory', 'driver', 'shop'].includes(value)
    },
    fields: {
      type: Array,
      default: () => []
    },
    data: {
      type: Object,
      default: () => ({})
    },
    image: {
      type: String,
      default: ''
    },
    txid: {
      type: String,
      default: ''
    },
    timestamp: {
      type: String,
      default: ''
    }
  },
  computed: {
    icon() {
      const icons = {
        farmer: 'el-icon-s-opportunity',
        factory: 'el-icon-s-cooperation',
        driver: 'el-icon-van',
        shop: 'el-icon-shopping-cart-1'
      }
      return icons[this.type] || 'el-icon-document'
    },
    imageUrl() {
      if (!this.image) return ''
      const baseApi = process.env.VUE_APP_BASE_API
      return `${baseApi}getImg/${this.image}`
    }
  },
  methods: {
    getFieldValue(field) {
      if (typeof field === 'string') {
        return this.data[field]
      }
      if (field.key) {
        return this.data[field.key]
      }
      return field.value
    }
  }
}
</script>

<style lang="scss" scoped>
.trace-info-card {
  margin-bottom: 20px;

  &.farmer {
    border-left: 4px solid #67c23a;
  }

  &.factory {
    border-left: 4px solid #409eff;
  }

  &.driver {
    border-left: 4px solid #e6a23c;
  }

  &.shop {
    border-left: 4px solid #f56c6c;
  }

  .card-header {
    display: flex;
    align-items: center;
    font-weight: bold;
    font-size: 16px;

    i {
      margin-right: 8px;
      font-size: 18px;
    }
  }

  .card-body {
    ::v-deep .el-form-item {
      margin-bottom: 12px;
    }
  }

  .tx-info {
    margin-top: 15px;
    padding-top: 15px;
    border-top: 1px solid #ebeef5;

    .tx-field {
      display: flex;
      margin-bottom: 8px;
      font-size: 12px;

      &:last-child {
        margin-bottom: 0;
      }

      .tx-label {
        color: #909399;
        min-width: 80px;
      }

      .tx-value {
        color: #606266;
        word-break: break-all;
      }
    }
  }
}
</style>
