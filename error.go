package ek

import "github.com/pkg/errors"

// kafka集群配置不存在
var ErrClusterNotExist = errors.New("package:ek.cluster.notExist")
// kafka生产者配置不存在
var ErrConsumerGroupNotExist = errors.New("package:ek.consumerGroup.notExist")
// kafka生产者配置不存在
var ErrProducerNotExist = errors.New("package:ek.producer.notExist")