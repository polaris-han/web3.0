1.注册    
请求方式：post
URL：/register
参数类型：body
参数示例：
{
    "username":"",
    "email":"",
    "password":""
}

2.登录    
请求方式：post
URL：/login
参数类型：body
请求示例：
{
    "username":"",
    "password":""
}

3.文章列表
请求方式：get
URL：/posts

4.查询文章详情
请求方式：get
URL：/posts/:id

5.发表文章
请求方式：post
URL：/api/posts
token类型：Bearer <token>
参数类型：body
请求示例：
{
    "title":"",
    "content":""
}

6.修改文章
请求方式：put
token类型：Bearer <token>
URL：/api/posts
参数类型：body
请求示例：
{
    "id":"",
    "title":"",
    "content":""
}

7.删除文章
请求方式：delete
token类型：Bearer <token>
URL：/api/posts/:id
参数类型：path

8.发表评论
请求方式：post
token类型：Bearer <token>
URL：/api/comments
参数类型：body
请求示例：
{
    "content":"",
    "postId":""
}

9.查询文章所有评论
请求方式：get
token类型：Bearer <token>
URL：/api/posts/:postID/comments
参数类型：path

postman文件版本v2.1
postman版本v11.60.0


