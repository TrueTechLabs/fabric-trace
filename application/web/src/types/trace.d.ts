// 溯源相关类型定义

/**
 * 农产品输入信息
 */
export interface FarmerInput {
  fa_fruitName: string
  fa_origin: string
  fa_plantTime: string
  fa_pickingTime: string
  fa_farmerName: string
  fa_img?: string
  fa_txid: string
  fa_timestamp: string
}

/**
 * 工厂输入信息
 */
export interface FactoryInput {
  fac_productName: string
  fac_productionbatch: string
  fac_productionTime: string
  fac_factoryName: string
  fac_contactNumber: string
  fac_img?: string
  fac_txid: string
  fac_timestamp: string
}

/**
 * 司机输入信息
 */
export interface DriverInput {
  dr_name: string
  dr_age: string
  dr_phone: string
  dr_carNumber: string
  dr_transport: string
  dr_img?: string
  dr_txid: string
  dr_timestamp: string
}

/**
 * 商店输入信息
 */
export interface ShopInput {
  sh_storeTime: string
  sh_sellTime: string
  sh_shopName: string
  sh_shopAddress: string
  sh_shopPhone: string
  sh_img?: string
  sh_txid: string
  sh_timestamp: string
}

/**
 * 完整溯源数据
 */
export interface TraceData {
  traceability_code: string
  farmer_input: FarmerInput
  factory_input: FactoryInput
  driver_input: DriverInput
  shop_input: ShopInput
}

/**
 * API响应格式
 */
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
  txid?: string
  traceability_code?: string
}
