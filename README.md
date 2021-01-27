### 说明
使用golang语言colly框架爬取barcodeLookup以及中国商品信息平台网站的条码信息并用excel作为存储载体记录下来。
### 使用
根据需要修改自己的excel字段，修改barcodeLookup cookie 条码不存在状态码为2，中国商品信息平台只爬取状态码为2的条码，无匹配数据状态码修改为3
### cookie
barcodeLookup 使用前需要手动更新cookie 否则无法爬取（报403错误）
