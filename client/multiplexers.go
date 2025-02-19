package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func AccountMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	client := meta.(*Client)
	for accountID := range client.ServicesManager.services {
		l = append(l, client.withAccountID(accountID))
	}
	return l
}

func ServiceAccountRegionMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for accountID := range client.ServicesManager.services {
			for region := range client.ServicesManager.services[accountID] {
				if !isSupportedServiceForRegion(service, region) {
					meta.Logger().Trace("region is not supported for service", "service", service, "region", region)
					continue
				}
				l = append(l, client.withAccountIDAndRegion(accountID, region))
			}
		}
		return l
	}
}

var AllNamespaces = []string{ // this is only used in applicationautoscaling
	"comprehend", "rds", "sagemaker", "appstream", "elasticmapreduce", "dynamodb", "lambda", "ecs", "cassandra", "ec2", "neptune", "kafka", "custom-resource", "elasticache",
}

func ServiceAccountRegionNamespaceMultiplexer(service string) func(meta schema.ClientMeta) []schema.ClientMeta {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		client := meta.(*Client)
		for accountID := range client.ServicesManager.services {
			for region := range client.ServicesManager.services[accountID] {
				for _, ns := range AllNamespaces {
					if !isSupportedServiceForRegion(service, region) {
						meta.Logger().Trace("region is not supported for service", "service", service, "region", region)
						continue
					}
					l = append(l, client.withAccountIDRegionAndNamespace(accountID, region, ns))
				}
			}
		}
		return l
	}
}
