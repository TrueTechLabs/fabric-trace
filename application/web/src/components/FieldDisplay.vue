<template>
  <el-form-item :label="label" :label-width="labelWidth">
    <span v-if="!link" class="field-value">{{ value }}</span>
    <a v-else :href="value" target="_blank" class="field-link">{{ value }}</a>
  </el-form-item>
</template>

<script>
export default {
  name: 'FieldDisplay',
  props: {
    label: {
      type: String,
      required: true
    },
    value: {
      type: [String, Number],
      default: ''
    },
    link: {
      type: Boolean,
      default: false
    },
    labelWidth: {
      type: String,
      default: '120px'
    }
  }
}
</script>

<style lang="scss" scoped>
.field-value {
  color: #606266;
}

.field-link {
  color: #409eff;
  text-decoration: none;

  &:hover {
    text-decoration: underline;
  }
}
</style>
