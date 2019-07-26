# event-kafka
Event Kafka

cli producer producerName version


{
    "id":"20181214150036-0.90816200-121.35.3.218-19310-1",
    "eventKey":"GOODS_BUY",
    "version": "v2.0.0",
    "forward": false,
    "broadcast": false, 
    "data":{
        "uid": 168,
        "goods_id":1,
        "num":121129,
        "operation_uid":0,
        "operation_time":"2018-12-14 15:00:36"
    },
    "time":1544770836,
    "ip":"121.35.3.218",
    "additional": {
        "forward": ["转发的topic"]
        "broadcast": ["广播的topic"],
    }
}

 v0_t_default_bbs_user_log
 v0_g_default_bbs_user_limit
 v0_p_default_bbs_user_log
 p0_producerName_
 
 
 'v0_g_default_bbs_user_limit' => [// 消费者群组名 group id: 必须唯一
         'cluster'   => 'default',// 集群名: zwei-kafka-cluster.php中配置
         'class'     => '',// 使用那种类型消费者消费消息
         'topics'    => [// 主题列表
             "test6",
         ],
         'options'   => [// kafka配置选项
             "offset.store.method" => "broker",// offset保存在broker中
         ],
         'timeout-ms' => 3000,// 消费者超时时间
         'client-ids' => "",
         'broadcast' => 'v0_b_default_common_target',
         'events'    => [
             // 事件名 => 事件回调函数(必须是静态方法)
             // 初始化事件 CRM_ZNTK_INIT
             'EVENT_INIT' => '\Zwei\Kafka\AppEventConsumer::testEventCallback',
         ],
     ],