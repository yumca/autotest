# autotest
说明
title 标题
typ 类型 post|get
url 请求地址
header 头部内容
param 请求内容 支持两种格式 post&get
result 返回内容

内容格式说明  每个内容字段下有 type为内容类型限制  根据type在type同级下寻找type所对应的数据内容
{
    内容字段：{
        "type":"内容类型",
        "内容类型":"数据内容"
    }
}
内容类型
default  默认类型
string  字符串类型
int 数字类型
datetime 格式化时间  默认Y-m-d H:i:s
timestamp 10位时间戳
object 对象类型
array 数组类型
objectjson 对象json化类型
objectjson 数组json化类型

内容类型限制
len@10  限制长度为10
max@3 限制最大长度为3
min@3  限制最小长度为3

{
    "title": "",
    "typ": "post",
    "url": "",
    "header": {
        "token": {
            "type": "default|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp",
            "default": ""
        }
    },
    "param": {
        "post": {
            "flow_code": {
                "type": "default|object|array|objectjson|arrayjson|string|int|datetime|timestamp",
                "string": "len@10|max@10",
                "int": "len@10|max@10",
                "datetime": "Y-m-d H:i:s"
            }
        },
        "get": {
            "flow_code": {
                "type": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp",
                "json": {
                    "a": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp"
                }
            }
        }
    },
    "result": {
        "code": {
            "type": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp"
        },
        "message": {
            "type": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp"
        },
        "time": {
            "type": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp"
        },
        "data": {
            "type": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp",
            "array": {
                "a": {
                    "type": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp",
                    "object": {
                        "b": {
                            "type": "default|object|array|objectjson|arrayjson|string#len@10#max@10|int#len@5|datetime#Y-m-d H:i:s|timestamp"
                        }
                    }
                }
            }
        }
    }
}
