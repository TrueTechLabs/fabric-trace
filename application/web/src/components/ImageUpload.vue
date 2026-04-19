<template>
  <div class="image-upload-container">
    <el-upload
      :action="action"
      :show-file-list="false"
      :on-change="handleChange"
      :before-upload="beforeUpload"
      :accept="accept"
      :disabled="disabled"
    >
      <el-button :disabled="disabled" type="primary" size="small">
        {{ buttonText }}
      </el-button>
    </el-upload>

    <div v-if="preview" class="image-preview">
      <img :src="preview" :alt="alt" />
      <div v-if="showRemove" class="remove-btn" @click="handleRemove">
        <i class="el-icon-delete" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ImageUpload',
  props: {
    value: {
      type: [String, File],
      default: null
    },
    action: {
      type: String,
      default: '#'
    },
    accept: {
      type: String,
      default: 'image/*'
    },
    maxSize: {
      type: Number,
      default: 2 // MB
    },
    disabled: {
      type: Boolean,
      default: false
    },
    buttonText: {
      type: String,
      default: '选择图片'
    },
    alt: {
      type: String,
      default: '图片预览'
    },
    showRemove: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      preview: null
    }
  },
  watch: {
    value: {
      immediate: true,
      handler(val) {
        if (val && typeof val === 'string') {
          this.preview = val
        }
      }
    }
  },
  methods: {
    handleChange(file) {
      this.preview = URL.createObjectURL(file.raw)
      this.$emit('input', file.raw)
      this.$emit('change', file)
    },
    beforeUpload(file) {
      const isImage = file.type.startsWith('image/')
      const isLtMaxSize = file.size / 1024 / 1024 < this.maxSize

      if (!isImage) {
        this.$message.error('只能上传图片文件!')
        return false
      }
      if (!isLtMaxSize) {
        this.$message.error(`图片大小不能超过 ${this.maxSize}MB!`)
        return false
      }
      // 禁止自动上传
      return false
    },
    handleRemove() {
      this.preview = null
      this.$emit('input', null)
      this.$emit('remove')
    }
  }
}
</script>

<style lang="scss" scoped>
.image-upload-container {
  display: inline-block;
}

.image-preview {
  position: relative;
  display: inline-block;
  margin-top: 10px;

  img {
    max-width: 200px;
    max-height: 200px;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
  }

  .remove-btn {
    position: absolute;
    top: 5px;
    right: 5px;
    width: 24px;
    height: 24px;
    background: rgba(0, 0, 0, 0.6);
    border-radius: 50%;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;
    font-size: 14px;

    &:hover {
      background: rgba(0, 0, 0, 0.8);
    }
  }
}
</style>
