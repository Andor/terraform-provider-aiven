// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func kafkaUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Kafka user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"additional_backup_regions": {
				Deprecated:  "This property is deprecated.",
				Description: "Additional Cloud Regions for Backup Replication.",
				Elem: &schema.Schema{
					Description: "Target cloud.",
					Type:        schema.TypeString,
				},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"aiven_kafka_topic_messages": {
				Description: "Allow access to read Kafka topic messages in the Aiven Console and REST API.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"custom_domain": {
				Description: "Serve the web frontend using a custom CNAME pointing to the Aiven DNS name.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"ip_filter": {
				Deprecated:  "Deprecated. Use `ip_filter_string` instead.",
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"ip_filter_object": {
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"description": {
						Description: "Description for IP filter list entry.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"network": {
						Description: "CIDR address block.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"ip_filter_string": {
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"kafka": {
				Description: "Kafka broker configuration values",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"auto_create_topics_enable": {
						Description: "Enable auto creation of topics.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"compression_type": {
						Description:  "Specify the final compression type for a given topic. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'uncompressed' which is equivalent to no compression; and 'producer' which means retain the original compression codec set by the producer.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"gzip", "snappy", "lz4", "zstd", "uncompressed", "producer"}, false),
					},
					"connections_max_idle_ms": {
						Description: "Idle connections timeout: the server socket processor threads close the connections that idle for longer than this.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"default_replication_factor": {
						Description: "Replication factor for autocreated topics.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"group_initial_rebalance_delay_ms": {
						Description: "The amount of time, in milliseconds, the group coordinator will wait for more consumers to join a new group before performing the first rebalance. A longer delay means potentially fewer rebalances, but increases the time until processing begins. The default value for this is 3 seconds. During development and testing it might be desirable to set this to 0 in order to not delay test execution time.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"group_max_session_timeout_ms": {
						Description: "The maximum allowed session timeout for registered consumers. Longer timeouts give consumers more time to process messages in between heartbeats at the cost of a longer time to detect failures.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"group_min_session_timeout_ms": {
						Description: "The minimum allowed session timeout for registered consumers. Longer timeouts give consumers more time to process messages in between heartbeats at the cost of a longer time to detect failures.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_cleaner_delete_retention_ms": {
						Description: "How long are delete records retained?",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_cleaner_max_compaction_lag_ms": {
						Description: "The maximum amount of time message will remain uncompacted. Only applicable for logs that are being compacted.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_cleaner_min_cleanable_ratio": {
						Description: "Controls log compactor frequency. Larger value means more frequent compactions but also more space wasted for logs. Consider setting log.cleaner.max.compaction.lag.ms to enforce compactions sooner, instead of setting a very high value for this option.",
						Optional:    true,
						Type:        schema.TypeFloat,
					},
					"log_cleaner_min_compaction_lag_ms": {
						Description: "The minimum time a message will remain uncompacted in the log. Only applicable for logs that are being compacted.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_cleanup_policy": {
						Description:  "The default cleanup policy for segments beyond the retention window.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"delete", "compact", "compact,delete"}, false),
					},
					"log_flush_interval_messages": {
						Description: "The number of messages accumulated on a log partition before messages are flushed to disk.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_flush_interval_ms": {
						Description: "The maximum time in ms that a message in any topic is kept in memory before flushed to disk. If not set, the value in log.flush.scheduler.interval.ms is used.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_index_interval_bytes": {
						Description: "The interval with which Kafka adds an entry to the offset index.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_index_size_max_bytes": {
						Description: "The maximum size in bytes of the offset index.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_local_retention_bytes": {
						Description: "The maximum size of local log segments that can grow for a partition before it gets eligible for deletion. If set to -2, the value of log.retention.bytes is used. The effective value should always be less than or equal to log.retention.bytes value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_local_retention_ms": {
						Description: "The number of milliseconds to keep the local log segments before it gets eligible for deletion. If set to -2, the value of log.retention.ms is used. The effective value should always be less than or equal to log.retention.ms value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_message_downconversion_enable": {
						Description: "This configuration controls whether down-conversion of message formats is enabled to satisfy consume requests. .",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"log_message_timestamp_difference_max_ms": {
						Description: "The maximum difference allowed between the timestamp when a broker receives a message and the timestamp specified in the message.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_message_timestamp_type": {
						Description:  "Define whether the timestamp in the message is message create time or log append time.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"CreateTime", "LogAppendTime"}, false),
					},
					"log_preallocate": {
						Description: "Should pre allocate file when create new segment?",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"log_retention_bytes": {
						Description: "The maximum size of the log before deleting messages.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_retention_hours": {
						Description: "The number of hours to keep a log file before deleting it.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_retention_ms": {
						Description: "The number of milliseconds to keep a log file before deleting it (in milliseconds), If not set, the value in log.retention.minutes is used. If set to -1, no time limit is applied.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_roll_jitter_ms": {
						Description: "The maximum jitter to subtract from logRollTimeMillis (in milliseconds). If not set, the value in log.roll.jitter.hours is used.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_roll_ms": {
						Description: "The maximum time before a new log segment is rolled out (in milliseconds).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_segment_bytes": {
						Description: "The maximum size of a single log file.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_segment_delete_delay_ms": {
						Description: "The amount of time to wait before deleting a file from the filesystem.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_connections_per_ip": {
						Description: "The maximum number of connections allowed from each ip address (defaults to 2147483647).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_incremental_fetch_session_cache_slots": {
						Description: "The maximum number of incremental fetch sessions that the broker will maintain.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"message_max_bytes": {
						Description: "The maximum size of message that the server can receive.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"min_insync_replicas": {
						Description: "When a producer sets acks to 'all' (or '-1'), min.insync.replicas specifies the minimum number of replicas that must acknowledge a write for the write to be considered successful.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"num_partitions": {
						Description: "Number of partitions for autocreated topics.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"offsets_retention_minutes": {
						Description: "Log retention window in minutes for offsets topic.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_purgatory_purge_interval_requests": {
						Description: "The purge interval (in number of requests) of the producer request purgatory(defaults to 1000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"replica_fetch_max_bytes": {
						Description: "The number of bytes of messages to attempt to fetch for each partition (defaults to 1048576). This is not an absolute maximum, if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that progress can be made.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"replica_fetch_response_max_bytes": {
						Description: "Maximum bytes expected for the entire fetch response (defaults to 10485760). Records are fetched in batches, and if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that progress can be made. As such, this is not an absolute maximum.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"sasl_oauthbearer_expected_audience": {
						Description: "The (optional) comma-delimited setting for the broker to use to verify that the JWT was issued for one of the expected audiences.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"sasl_oauthbearer_expected_issuer": {
						Description: "Optional setting for the broker to use to verify that the JWT was created by the expected issuer.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"sasl_oauthbearer_jwks_endpoint_url": {
						Description: "OIDC JWKS endpoint URL. By setting this the SASL SSL OAuth2/OIDC authentication is enabled. See also other options for SASL OAuth2/OIDC. .",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"sasl_oauthbearer_sub_claim_name": {
						Description: "Name of the scope from which to extract the subject claim from the JWT. Defaults to sub.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"socket_request_max_bytes": {
						Description: "The maximum number of bytes in a socket request (defaults to 104857600).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"transaction_partition_verification_enable": {
						Description: "Enable verification that checks that the partition has been added to the transaction before writing transactional records to the partition.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"transaction_remove_expired_transaction_cleanup_interval_ms": {
						Description: "The interval at which to remove transactions that have expired due to transactional.id.expiration.ms passing (defaults to 3600000 (1 hour)).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"transaction_state_log_segment_bytes": {
						Description: "The transaction topic segment bytes should be kept relatively small in order to facilitate faster log compaction and cache loads (defaults to 104857600 (100 mebibytes)).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"kafka_authentication_methods": {
				Description: "Kafka authentication methods",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"certificate": {
						Description: "Enable certificate/SSL authentication. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"sasl": {
						Description: "Enable SASL authentication. The default value is `false`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"kafka_connect": {
				Description: "Enable Kafka Connect service. The default value is `false`.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"kafka_connect_config": {
				Description: "Kafka Connect configuration values",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"connector_client_config_override_policy": {
						Description:  "Defines what client configurations can be overridden by the connector. Default is None.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"None", "All"}, false),
					},
					"consumer_auto_offset_reset": {
						Description:  "What to do when there is no initial offset in Kafka or if the current offset does not exist any more on the server. Default is earliest.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"earliest", "latest"}, false),
					},
					"consumer_fetch_max_bytes": {
						Description: "Records are fetched in batches by the consumer, and if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that the consumer can make progress. As such, this is not a absolute maximum.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"consumer_isolation_level": {
						Description:  "Transaction read isolation level. read_uncommitted is the default, but read_committed can be used if consume-exactly-once behavior is desired.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"read_uncommitted", "read_committed"}, false),
					},
					"consumer_max_partition_fetch_bytes": {
						Description: "Records are fetched in batches by the consumer.If the first record batch in the first non-empty partition of the fetch is larger than this limit, the batch will still be returned to ensure that the consumer can make progress. .",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"consumer_max_poll_interval_ms": {
						Description: "The maximum delay in milliseconds between invocations of poll() when using consumer group management (defaults to 300000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"consumer_max_poll_records": {
						Description: "The maximum number of records returned in a single call to poll() (defaults to 500).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"offset_flush_interval_ms": {
						Description: "The interval at which to try committing offsets for tasks (defaults to 60000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"offset_flush_timeout_ms": {
						Description: "Maximum number of milliseconds to wait for records to flush and partition offset data to be committed to offset storage before cancelling the process and restoring the offset data to be committed in a future attempt (defaults to 5000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_batch_size": {
						Description: "This setting gives the upper bound of the batch size to be sent. If there are fewer than this many bytes accumulated for this partition, the producer will 'linger' for the linger.ms time waiting for more records to show up. A batch size of zero will disable batching entirely (defaults to 16384).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_buffer_memory": {
						Description: "The total bytes of memory the producer can use to buffer records waiting to be sent to the broker (defaults to 33554432).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_compression_type": {
						Description:  "Specify the default compression type for producers. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'none' which is the default and equivalent to no compression.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"gzip", "snappy", "lz4", "zstd", "none"}, false),
					},
					"producer_linger_ms": {
						Description: "This setting gives the upper bound on the delay for batching: once there is batch.size worth of records for a partition it will be sent immediately regardless of this setting, however if there are fewer than this many bytes accumulated for this partition the producer will 'linger' for the specified time waiting for more records to show up. Defaults to 0.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_max_request_size": {
						Description: "This setting will limit the number of record batches the producer will send in a single request to avoid sending huge requests.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"scheduled_rebalance_max_delay_ms": {
						Description: "The maximum delay that is scheduled in order to wait for the return of one or more departed workers before rebalancing and reassigning their connectors and tasks to the group. During this period the connectors and tasks of the departed workers remain unassigned. Defaults to 5 minutes.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"session_timeout_ms": {
						Description: "The timeout in milliseconds used to detect failures when using Kafka’s group management facilities (defaults to 10000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"kafka_rest": {
				Description: "Enable Kafka-REST service. The default value is `false`.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"kafka_rest_authorization": {
				Description: "Enable authorization in Kafka-REST service.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"kafka_rest_config": {
				Description: "Kafka REST configuration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"consumer_enable_auto_commit": {
						Description: "If true the consumer's offset will be periodically committed to Kafka in the background. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"consumer_request_max_bytes": {
						Description: "Maximum number of bytes in unencoded message keys and values by a single request. The default value is `67108864`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"consumer_request_timeout_ms": {
						Description:  "The maximum total time to wait for messages for a request if the maximum number of messages has not yet been reached. The default value is `1000`.",
						Optional:     true,
						Type:         schema.TypeInt,
						ValidateFunc: validation.IntInSlice([]int{1000, 15000, 30000}),
					},
					"name_strategy": {
						Description:  "Name strategy to use when selecting subject for storing schemas. The default value is `topic_name`.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"topic_name", "record_name", "topic_record_name"}, false),
					},
					"name_strategy_validation": {
						Description: "If true, validate that given schema is registered under expected subject name by the used name strategy when producing messages. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"producer_acks": {
						Description:  "The number of acknowledgments the producer requires the leader to have received before considering a request complete. If set to 'all' or '-1', the leader will wait for the full set of in-sync replicas to acknowledge the record. The default value is `1`.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"all", "-1", "0", "1"}, false),
					},
					"producer_compression_type": {
						Description:  "Specify the default compression type for producers. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'none' which is the default and equivalent to no compression.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"gzip", "snappy", "lz4", "zstd", "none"}, false),
					},
					"producer_linger_ms": {
						Description: "Wait for up to the given delay to allow batching records together. The default value is `0`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_max_request_size": {
						Description: "The maximum size of a request in bytes. Note that Kafka broker can also cap the record batch size. The default value is `1048576`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"simpleconsumer_pool_size_max": {
						Description: "Maximum number of SimpleConsumers that can be instantiated per broker. The default value is `25`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"kafka_version": {
				Description:  "Kafka major version.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"3.2", "3.3", "3.1", "3.4", "3.5", "3.6"}, false),
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"kafka": {
						Description: "Allow clients to connect to kafka with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka_connect": {
						Description: "Allow clients to connect to kafka_connect with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka_rest": {
						Description: "Allow clients to connect to kafka_rest with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"schema_registry": {
						Description: "Allow clients to connect to schema_registry with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"privatelink_access": {
				Description: "Allow access to selected service components through Privatelink",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"jolokia": {
						Description: "Enable jolokia.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka": {
						Description: "Enable kafka.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka_connect": {
						Description: "Enable kafka_connect.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka_rest": {
						Description: "Enable kafka_rest.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Enable prometheus.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"schema_registry": {
						Description: "Enable schema_registry.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"kafka": {
						Description: "Allow clients to connect to kafka from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka_connect": {
						Description: "Allow clients to connect to kafka_connect from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka_rest": {
						Description: "Allow clients to connect to kafka_rest from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"schema_registry": {
						Description: "Allow clients to connect to schema_registry from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"schema_registry": {
				Description: "Enable Schema-Registry service. The default value is `false`.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"schema_registry_config": {
				Description: "Schema Registry configuration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"leader_eligibility": {
						Description: "If true, Karapace / Schema Registry on the service nodes can participate in leader election. It might be needed to disable this when the schemas topic is replicated to a secondary cluster and Karapace / Schema Registry there must not participate in leader election. Defaults to `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"topic_name": {
						Description: "The durable single partition topic that acts as the durable log for the data. This topic must be compacted to avoid losing data due to retention policy. Please note that changing this configuration in an existing Schema Registry / Karapace setup leads to previous schemas being inaccessible, data encoded with them potentially unreadable and schema ID sequence put out of order. It's only possible to do the switch while Schema Registry / Karapace is disabled. Defaults to `_schemas`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"service_log": {
				Description: "Store logs for the service so that they are available in the HTTP API and console.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"static_ips": {
				Description: "Use static public IP addresses.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"tiered_storage": {
				Description: "Tiered storage configuration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"enabled": {
						Description: "Whether to enable the tiered storage functionality.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"local_cache": {
						Deprecated:  "This property is deprecated.",
						Description: "Local cache configuration",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{"size": {
							Deprecated:  "This property is deprecated.",
							Description: "Local cache size in bytes.",
							Optional:    true,
							Type:        schema.TypeInt,
						}}},
						MaxItems: 1,
						Optional: true,
						Type:     schema.TypeList,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
