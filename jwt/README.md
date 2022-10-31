JWT 构成：
* HEAD
  * 由 typ 和 alg 组成，指明类型和加密算法
* PAYLOAD
  * 负载信息，如：ExpiresAt 失效时间，Subject 主题...
* VERIFY SIGNATURE
  * 加密算法(base64UrlEncode(HEAD).base64UrlEncode(PAYLOAD), 密钥)