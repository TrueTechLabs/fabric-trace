(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-a6a6d8c8"],{"0825":function(t,a,e){},4939:function(t,a,e){"use strict";e.r(a);var r=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"uplink-container"},[e("div",{staticStyle:{color:"#909399","margin-bottom":"30px"}},[t._v(" 当前用户："+t._s(t.name)+"; 用户角色: "+t._s(t.userType)+" ")]),e("div",[e("el-form",{ref:"form",attrs:{model:t.tracedata,"label-width":"80px",size:"mini"}},[e("el-form-item",{directives:[{name:"show",rawName:"v-show",value:"种植户"!=t.userType&"消费者"!=t.userType,expression:"userType!='种植户'&userType!='消费者'"}],staticStyle:{width:"300px"},attrs:{label:"溯源码:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.traceability_code,callback:function(a){t.$set(t.tracedata,"traceability_code",a)},expression:"tracedata.traceability_code"}})],1),e("div",{directives:[{name:"show",rawName:"v-show",value:"种植户"==t.userType,expression:"userType=='种植户'"}]},[e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"农产品名称:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Farmer_input.Fa_fruitName,callback:function(a){t.$set(t.tracedata.Farmer_input,"Fa_fruitName",a)},expression:"tracedata.Farmer_input.Fa_fruitName"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"产地:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Farmer_input.Fa_origin,callback:function(a){t.$set(t.tracedata.Farmer_input,"Fa_origin",a)},expression:"tracedata.Farmer_input.Fa_origin"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"种植时间:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Farmer_input.Fa_plantTime,callback:function(a){t.$set(t.tracedata.Farmer_input,"Fa_plantTime",a)},expression:"tracedata.Farmer_input.Fa_plantTime"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"采摘时间:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Farmer_input.Fa_pickingTime,callback:function(a){t.$set(t.tracedata.Farmer_input,"Fa_pickingTime",a)},expression:"tracedata.Farmer_input.Fa_pickingTime"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"种植户名称:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Farmer_input.Fa_farmerName,callback:function(a){t.$set(t.tracedata.Farmer_input,"Fa_farmerName",a)},expression:"tracedata.Farmer_input.Fa_farmerName"}})],1)],1),e("div",{directives:[{name:"show",rawName:"v-show",value:"工厂"==t.userType,expression:"userType=='工厂'"}]},[e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"商品名称:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Factory_input.Fac_productName,callback:function(a){t.$set(t.tracedata.Factory_input,"Fac_productName",a)},expression:"tracedata.Factory_input.Fac_productName"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"生产批次:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Factory_input.Fac_productionbatch,callback:function(a){t.$set(t.tracedata.Factory_input,"Fac_productionbatch",a)},expression:"tracedata.Factory_input.Fac_productionbatch"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"生产时间:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Factory_input.Fac_productionTime,callback:function(a){t.$set(t.tracedata.Factory_input,"Fac_productionTime",a)},expression:"tracedata.Factory_input.Fac_productionTime"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"工厂名称与厂址:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Factory_input.Fac_factoryName,callback:function(a){t.$set(t.tracedata.Factory_input,"Fac_factoryName",a)},expression:"tracedata.Factory_input.Fac_factoryName"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"工厂电话:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Factory_input.Fac_contactNumber,callback:function(a){t.$set(t.tracedata.Factory_input,"Fac_contactNumber",a)},expression:"tracedata.Factory_input.Fac_contactNumber"}})],1)],1),e("div",{directives:[{name:"show",rawName:"v-show",value:"运输司机"==t.userType,expression:"userType=='运输司机'"}]},[e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"姓名:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Driver_input.Dr_name,callback:function(a){t.$set(t.tracedata.Driver_input,"Dr_name",a)},expression:"tracedata.Driver_input.Dr_name"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"年龄:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Driver_input.Dr_age,callback:function(a){t.$set(t.tracedata.Driver_input,"Dr_age",a)},expression:"tracedata.Driver_input.Dr_age"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"联系电话:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Driver_input.Dr_phone,callback:function(a){t.$set(t.tracedata.Driver_input,"Dr_phone",a)},expression:"tracedata.Driver_input.Dr_phone"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"车牌号:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Driver_input.Dr_carNumber,callback:function(a){t.$set(t.tracedata.Driver_input,"Dr_carNumber",a)},expression:"tracedata.Driver_input.Dr_carNumber"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"运输工具:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Driver_input.Dr_transport,callback:function(a){t.$set(t.tracedata.Driver_input,"Dr_transport",a)},expression:"tracedata.Driver_input.Dr_transport"}})],1)],1),e("div",{directives:[{name:"show",rawName:"v-show",value:"商店"==t.userType,expression:"userType=='商店'"}]},[e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"存入时间:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Shop_input.Sh_storeTime,callback:function(a){t.$set(t.tracedata.Shop_input,"Sh_storeTime",a)},expression:"tracedata.Shop_input.Sh_storeTime"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"销售时间:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Shop_input.Sh_sellTime,callback:function(a){t.$set(t.tracedata.Shop_input,"Sh_sellTime",a)},expression:"tracedata.Shop_input.Sh_sellTime"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"商店名称:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Shop_input.Sh_shopName,callback:function(a){t.$set(t.tracedata.Shop_input,"Sh_shopName",a)},expression:"tracedata.Shop_input.Sh_shopName"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"商店位置:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Shop_input.Sh_shopAddress,callback:function(a){t.$set(t.tracedata.Shop_input,"Sh_shopAddress",a)},expression:"tracedata.Shop_input.Sh_shopAddress"}})],1),e("el-form-item",{staticStyle:{width:"300px"},attrs:{label:"商店电话:","label-width":"120px"}},[e("el-input",{model:{value:t.tracedata.Shop_input.Sh_shopPhone,callback:function(a){t.$set(t.tracedata.Shop_input,"Sh_shopPhone",a)},expression:"tracedata.Shop_input.Sh_shopPhone"}})],1)],1)],1),e("span",{staticClass:"dialog-footer",staticStyle:{color:"gray"},attrs:{slot:"footer"},slot:"footer"},[e("el-button",{directives:[{name:"show",rawName:"v-show",value:"消费者"!=t.userType,expression:"userType != '消费者'"}],staticStyle:{"margin-left":"220px"},attrs:{type:"primary",plain:""},on:{click:function(a){return t.submittracedata()}}},[t._v("提 交")])],1),e("span",{directives:[{name:"show",rawName:"v-show",value:"消费者"==t.userType,expression:"userType == '消费者'"}],staticClass:"dialog-footer",staticStyle:{color:"gray"},attrs:{slot:"footer"},slot:"footer"},[t._v(" 消费者没有权限录入！请使用溯源功能! ")])],1)])},i=[],c=e("5530"),n=e("2f62"),p=e("89ff"),o={name:"Uplink",data:function(){return{tracedata:{traceability_code:"",Farmer_input:{Fa_fruitName:"",Fa_origin:"",Fa_plantTime:"",Fa_pickingTime:"",Fa_farmerName:""},Factory_input:{Fac_productName:"",Fac_productionbatch:"",Fac_productionTime:"",Fac_factoryName:"",Fac_contactNumber:""},Driver_input:{Dr_name:"",Dr_age:"",Dr_phone:"",Dr_carNumber:"",Dr_transport:""},Shop_input:{Sh_storeTime:"",Sh_sellTime:"",Sh_shopName:"",Sh_shopAddress:"",Sh_shopPhone:""}},loading:!1}},computed:Object(c["a"])({},Object(n["b"])(["name","userType"])),methods:{submittracedata:function(){var t=this;console.log(this.tracedata);var a=this.$loading({lock:!0,text:"数据上链中...",spinner:"el-icon-loading",background:"rgba(0, 0, 0, 0.7)"}),e=new FormData;switch(e.append("traceability_code",this.tracedata.traceability_code),this.userType){case"种植户":e.append("arg1",this.tracedata.Farmer_input.Fa_fruitName),e.append("arg2",this.tracedata.Farmer_input.Fa_origin),e.append("arg3",this.tracedata.Farmer_input.Fa_plantTime),e.append("arg4",this.tracedata.Farmer_input.Fa_pickingTime),e.append("arg5",this.tracedata.Farmer_input.Fa_farmerName);break;case"工厂":e.append("arg1",this.tracedata.Factory_input.Fac_productName),e.append("arg2",this.tracedata.Factory_input.Fac_productionbatch),e.append("arg3",this.tracedata.Factory_input.Fac_productionTime),e.append("arg4",this.tracedata.Factory_input.Fac_factoryName),e.append("arg5",this.tracedata.Factory_input.Fac_contactNumber);break;case"运输司机":e.append("arg1",this.tracedata.Driver_input.Dr_name),e.append("arg2",this.tracedata.Driver_input.Dr_age),e.append("arg3",this.tracedata.Driver_input.Dr_phone),e.append("arg4",this.tracedata.Driver_input.Dr_carNumber),e.append("arg5",this.tracedata.Driver_input.Dr_transport);break;case"商店":e.append("arg1",this.tracedata.Shop_input.Sh_storeTime),e.append("arg2",this.tracedata.Shop_input.Sh_sellTime),e.append("arg3",this.tracedata.Shop_input.Sh_shopName),e.append("arg4",this.tracedata.Shop_input.Sh_shopAddress),e.append("arg5",this.tracedata.Shop_input.Sh_shopPhone);break}Object(p["e"])(e).then((function(e){200===e.code?(a.close(),t.$message({message:"上链成功，交易ID："+e.txid+"\n溯源码："+e.traceability_code,type:"success"})):(a.close(),t.$message({message:"上链失败",type:"error"}))})).catch((function(t){a.close(),console.log(t)}))}}},l=o,s=(e("87d4"),e("2877")),u=Object(s["a"])(l,r,i,!1,null,"288d5ca1",null);a["default"]=u.exports},"87d4":function(t,a,e){"use strict";e("0825")},"89ff":function(t,a,e){"use strict";e.d(a,"e",(function(){return i})),e.d(a,"c",(function(){return c})),e.d(a,"d",(function(){return n})),e.d(a,"a",(function(){return p})),e.d(a,"b",(function(){return o}));var r=e("b775");function i(t){return Object(r["a"])({url:"/uplink",method:"post",headers:{"Content-Type":"multipart/form-data"},data:t})}function c(t){return Object(r["a"])({url:"/getFruitInfo",method:"post",headers:{"Content-Type":"multipart/form-data"},data:t})}function n(t){return Object(r["a"])({url:"/getFruitList",method:"post",headers:{"Content-Type":"multipart/form-data"},data:t})}function p(t){return Object(r["a"])({url:"/getAllFruitInfo",method:"post",headers:{"Content-Type":"multipart/form-data"},data:t})}function o(t){return Object(r["a"])({url:"/getFruitHistory",method:"post",headers:{"Content-Type":"multipart/form-data"},data:t})}}}]);