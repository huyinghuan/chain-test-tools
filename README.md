# NFT 接口服务



## 证书申请

```
字段	类型	含义	备注
orgId	string	组织ID	必填
uuid	string	用户ID	*选填
userType	string	用户类型	必填
certUsage	string	证书用途	必填
privateKeyPwd	string	密钥密码	选填
country	string	证书字段（国家）	必填
locality	string	证书字段（城市）	必填
province	string	证书字段（省份）	必填
token	string	token	选填
```

userType: 1.root , 2.ca , 3.admin , 4.client , 5.consensus , 6.common

certUsage: 1.sign , 2.tls , 3.tls-sign , 4.tls-enc

userId 只有在申请的用户类型是ca的类型时，可以填写为空。在申请节点证书时，需要保证链上节点ID唯一。
