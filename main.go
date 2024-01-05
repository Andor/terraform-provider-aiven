package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"

	"github.com/aiven/terraform-provider-aiven/internal/server"
)

// registryPrefix is the registry prefix for the provider.
//
//go:generate go test -tags userconfig ./internal/schemautil/userconfig
//go:generate go run ./ucgenerator/... --services cassandra,clickhouse,flink,grafana,influxdb,kafka,kafka_connect,kafka_mirrormaker,m3aggregator,m3db,mysql,opensearch,pg,redis
//go:generate go run ./ucgenerator/... --integrations clickhouse_kafka,clickhouse_postgresql,datadog,external_aws_cloudwatch_logs,external_aws_cloudwatch_metrics,external_elasticsearch_logs,external_opensearch_logs,kafka_connect,kafka_logs,kafka_mirrormaker,logs,metrics,prometheus
//go:generate go run ./ucgenerator/... --integration-endpoints datadog,external_aws_cloudwatch_logs,external_aws_cloudwatch_metrics,external_elasticsearch_logs,external_google_cloud_bigquery,external_google_cloud_logging,external_kafka,external_opensearch_logs,external_postgresql,external_schema_registry,jolokia,prometheus,rsyslog
const registryPrefix = "registry.terraform.io/"

// version is the version of the provider.
var version = "dev"

func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")

	flag.Parse()
	ctx := context.Background()
	muxServer, err := server.NewMuxServer(ctx, version)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt
	if *debugFlag {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	name := registryPrefix + "aiven/aiven"

	//goland:noinspection GoBoolExpressions
	if version == "dev" {
		name = registryPrefix + "aiven-dev/aiven"
	}

	err = tf6server.Serve(
		name,
		func() tfprotov6.ProviderServer {
			return muxServer
		},
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
