package provider

import (
	"embed"

	"github.com/cloudquery/cq-provider-aws/resources/services/iot"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/resources/services/accessanalyzer"
	"github.com/cloudquery/cq-provider-aws/resources/services/acm"
	"github.com/cloudquery/cq-provider-aws/resources/services/apigateway"
	"github.com/cloudquery/cq-provider-aws/resources/services/apigatewayv2"
	"github.com/cloudquery/cq-provider-aws/resources/services/applicationautoscaling"
	"github.com/cloudquery/cq-provider-aws/resources/services/autoscaling"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudfront"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudtrail"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudwatch"
	"github.com/cloudquery/cq-provider-aws/resources/services/cloudwatchlogs"
	"github.com/cloudquery/cq-provider-aws/resources/services/codebuild"
	"github.com/cloudquery/cq-provider-aws/resources/services/cognito"
	"github.com/cloudquery/cq-provider-aws/resources/services/config"
	"github.com/cloudquery/cq-provider-aws/resources/services/dax"
	"github.com/cloudquery/cq-provider-aws/resources/services/directconnect"
	"github.com/cloudquery/cq-provider-aws/resources/services/dms"
	"github.com/cloudquery/cq-provider-aws/resources/services/dynamodb"
	"github.com/cloudquery/cq-provider-aws/resources/services/ec2"
	"github.com/cloudquery/cq-provider-aws/resources/services/ecr"
	"github.com/cloudquery/cq-provider-aws/resources/services/ecs"
	"github.com/cloudquery/cq-provider-aws/resources/services/efs"
	"github.com/cloudquery/cq-provider-aws/resources/services/eks"
	"github.com/cloudquery/cq-provider-aws/resources/services/elasticbeanstalk"
	"github.com/cloudquery/cq-provider-aws/resources/services/elasticsearch"
	"github.com/cloudquery/cq-provider-aws/resources/services/elbv1"
	"github.com/cloudquery/cq-provider-aws/resources/services/elbv2"
	"github.com/cloudquery/cq-provider-aws/resources/services/emr"
	"github.com/cloudquery/cq-provider-aws/resources/services/fsx"
	"github.com/cloudquery/cq-provider-aws/resources/services/guardduty"
	"github.com/cloudquery/cq-provider-aws/resources/services/iam"
	"github.com/cloudquery/cq-provider-aws/resources/services/kms"
	"github.com/cloudquery/cq-provider-aws/resources/services/lambda"
	"github.com/cloudquery/cq-provider-aws/resources/services/mq"
	"github.com/cloudquery/cq-provider-aws/resources/services/organizations"
	"github.com/cloudquery/cq-provider-aws/resources/services/rds"
	"github.com/cloudquery/cq-provider-aws/resources/services/redshift"
	"github.com/cloudquery/cq-provider-aws/resources/services/route53"
	"github.com/cloudquery/cq-provider-aws/resources/services/s3"
	"github.com/cloudquery/cq-provider-aws/resources/services/sagemaker"
	"github.com/cloudquery/cq-provider-aws/resources/services/secretsmanager"
	"github.com/cloudquery/cq-provider-aws/resources/services/sns"
	"github.com/cloudquery/cq-provider-aws/resources/services/sqs"
	"github.com/cloudquery/cq-provider-aws/resources/services/ssm"
	"github.com/cloudquery/cq-provider-aws/resources/services/waf"
	"github.com/cloudquery/cq-provider-aws/resources/services/wafv2"
)

var (
	//go:embed migrations/*/*.sql
	awsMigrations embed.FS
	Version       = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:            "aws",
		Version:         Version,
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		Migrations:      awsMigrations,
		ResourceMap: map[string]*schema.Table{
			"acm.certificates":                      acm.AcmCertificates(),
			"aws.regions":                           ec2.AwsRegions(),
			"accessanalyzer.analyzers":              accessanalyzer.AccessAnalyzerAnalyzer(),
			"apigateway.api_keys":                   apigateway.ApigatewayAPIKeys(),
			"apigateway.client_certificates":        apigateway.ApigatewayClientCertificates(),
			"apigateway.domain_names":               apigateway.ApigatewayDomainNames(),
			"apigateway.rest_apis":                  apigateway.ApigatewayRestApis(),
			"apigateway.usage_plans":                apigateway.ApigatewayUsagePlans(),
			"apigateway.vpc_links":                  apigateway.ApigatewayVpcLinks(),
			"apigatewayv2.apis":                     apigatewayv2.Apigatewayv2Apis(),
			"apigatewayv2.domain_names":             apigatewayv2.Apigatewayv2DomainNames(),
			"apigatewayv2.vpc_links":                apigatewayv2.Apigatewayv2VpcLinks(),
			"applicationautoscaling.policies":       applicationautoscaling.ApplicationautoscalingPolicies(),
			"autoscaling.launch_configurations":     autoscaling.AutoscalingLaunchConfigurations(),
			"autoscaling.groups":                    autoscaling.AutoscalingGroups(),
			"cloudfront.cache_policies":             cloudfront.CloudfrontCachePolicies(),
			"cloudfront.distributions":              cloudfront.CloudfrontDistributions(),
			"codebuild.projects":                    codebuild.CodebuildProjects(),
			"cloudtrail.trails":                     cloudtrail.CloudtrailTrails(),
			"cloudwatch.alarms":                     cloudwatch.CloudwatchAlarms(),
			"cloudwatchlogs.filters":                cloudwatchlogs.CloudwatchlogsFilters(),
			"cognito.identity_pools":                cognito.CognitoIdentityPools(),
			"cognito.user_pools":                    cognito.CognitoUserPools(),
			"config.configuration_recorders":        config.ConfigConfigurationRecorders(),
			"config.conformance_packs":              config.ConfigConformancePack(),
			"directconnect.connections":             directconnect.DirectconnectConnections(),
			"directconnect.gateways":                directconnect.DirectconnectGateways(),
			"directconnect.lags":                    directconnect.DirectconnectLags(),
			"directconnect.virtual_gateways":        directconnect.DirectconnectVirtualGateways(),
			"directconnect.virtual_interfaces":      directconnect.DirectconnectVirtualInterfaces(),
			"dax.clusters":                          dax.DaxClusters(),
			"dms.replication_instances":             dms.DmsReplicationInstances(),
			"dynamodb.tables":                       dynamodb.DynamodbTables(),
			"ec2.byoip_cidrs":                       ec2.Ec2ByoipCidrs(),
			"ec2.customer_gateways":                 ec2.Ec2CustomerGateways(),
			"ec2.ebs_volumes":                       ec2.Ec2EbsVolumes(),
			"ec2.ebs_snapshots":                     ec2.Ec2EbsSnapshots(),
			"ec2.flow_logs":                         ec2.Ec2FlowLogs(),
			"ec2.images":                            ec2.Ec2Images(),
			"ec2.instances":                         ec2.Ec2Instances(),
			"ec2.internet_gateways":                 ec2.Ec2InternetGateways(),
			"ec2.nat_gateways":                      ec2.Ec2NatGateways(),
			"ec2.network_acls":                      ec2.Ec2NetworkAcls(),
			"ec2.regional_config":                   ec2.Ec2RegionalConfig(),
			"ec2.route_tables":                      ec2.Ec2RouteTables(),
			"ec2.security_groups":                   ec2.Ec2SecurityGroups(),
			"ec2.subnets":                           ec2.Ec2Subnets(),
			"ec2.transit_gateways":                  ec2.Ec2TransitGateways(),
			"ec2.vpc_endpoints":                     ec2.Ec2VpcEndpoints(),
			"ec2.vpc_peering_connections":           ec2.Ec2VpcPeeringConnections(),
			"ec2.vpcs":                              ec2.Ec2Vpcs(),
			"ec2.eips":                              ec2.Ec2Eips(),
			"ec2.vpn_gateways":                      ec2.Ec2VpnGateways(),
			"ecr.repositories":                      ecr.EcrRepositories(),
			"ecs.clusters":                          ecs.EcsClusters(),
			"ecs.task_definitions":                  ecs.EcsTaskDefinitions(),
			"efs.filesystems":                       efs.EfsFilesystems(),
			"eks.clusters":                          eks.EksClusters(),
			"elasticbeanstalk.applications":         elasticbeanstalk.ElasticbeanstalkApplications(),
			"elasticbeanstalk.environments":         elasticbeanstalk.ElasticbeanstalkEnvironments(),
			"elasticsearch.domains":                 elasticsearch.ElasticsearchDomains(),
			"elbv1.load_balancers":                  elbv1.Elbv1LoadBalancers(),
			"elbv2.load_balancers":                  elbv2.Elbv2LoadBalancers(),
			"elbv2.target_groups":                   elbv2.Elbv2TargetGroups(),
			"emr.clusters":                          emr.EmrClusters(),
			"emr.block_public_access_configs":       emr.EmrBlockPublicAccessConfigs(),
			"fsx.backups":                           fsx.FsxBackups(),
			"guardduty.detectors":                   guardduty.GuarddutyDetectors(),
			"iam.accounts":                          iam.IamAccounts(),
			"iam.groups":                            iam.IamGroups(),
			"iam.openid_connect_identity_providers": iam.IamOpenidConnectIdentityProviders(),
			"iam.password_policies":                 iam.IamPasswordPolicies(),
			"iam.policies":                          iam.IamPolicies(),
			"iam.roles":                             iam.IamRoles(),
			"iam.saml_identity_providers":           iam.IamSamlIdentityProviders(),
			"iam.server_certificates":               iam.IamServerCertificates(),
			"iam.users":                             iam.IamUsers(),
			"iam.virtual_mfa_devices":               iam.IamVirtualMfaDevices(),
			"kms.keys":                              kms.KmsKeys(),
			"lambda.functions":                      lambda.LambdaFunctions(),
			"lambda.runtimes":                       lambda.LambdaRuntimes(),
			"lambda.layers":                         lambda.LambdaLayers(),
			"mq.brokers":                            mq.MqBrokers(),
			"organizations.accounts":                organizations.OrganizationsAccounts(),
			"rds.cluster_parameter_groups":          rds.RdsClusterParameterGroups(),
			"rds.db_parameter_groups":               rds.RdsDbParameterGroups(),
			"rds.db_security_groups":                rds.RdsDbSecurityGroups(),
			"rds.event_subscriptions":               rds.RdsEventSubscriptions(),
			"rds.certificates":                      rds.RdsCertificates(),
			"rds.clusters":                          rds.RdsClusters(),
			"rds.cluster_snapshots":                 rds.RdsClusterSnapshots(),
			"rds.db_snapshots":                      rds.RdsDbSnapshots(),
			"rds.db_subnet_groups":                  rds.RdsSubnetGroups(),
			"rds.instances":                         rds.RdsInstances(),
			"redshift.clusters":                     redshift.RedshiftClusters(),
			"redshift.subnet_groups":                redshift.RedshiftSubnetGroups(),
			"route53.domains":                       route53.Route53Domains(),
			"route53.health_checks":                 route53.Route53HealthChecks(),
			"route53.hosted_zones":                  route53.Route53HostedZones(),
			"route53.reusable_delegation_sets":      route53.Route53ReusableDelegationSets(),
			"route53.traffic_policies":              route53.Route53TrafficPolicies(),
			"s3.buckets":                            s3.S3Buckets(),
			"s3.accounts":                           s3.S3Accounts(),
			"sagemaker.notebook_instances":          sagemaker.SagemakerNotebookInstances(),
			"sagemaker.models":                      sagemaker.SagemakerModels(),
			"sagemaker.endpoint_configurations":     sagemaker.SagemakerEndpointConfigurations(),
			"sagemaker.training_jobs":               sagemaker.SagemakerTrainingJobs(),
			"secretsmanager.secrets":                secretsmanager.SecretsmanagerSecrets(),
			"sns.subscriptions":                     sns.SnsSubscriptions(),
			"sns.topics":                            sns.SnsTopics(),
			"sqs.queues":                            sqs.SQSQueues(),
			"ssm.documents":                         ssm.SsmDocuments(),
			"ssm.instances":                         ssm.SsmInstances(),
			"waf.rule_groups":                       waf.WafRuleGroups(),
			"waf.rules":                             waf.WafRules(),
			"waf.subscribed_rule_groups":            waf.WafSubscribedRuleGroups(),
			"waf.web_acls":                          waf.WafWebAcls(),
			"wafv2.managed_rule_groups":             wafv2.Wafv2ManagedRuleGroups(),
			"wafv2.rule_groups":                     wafv2.Wafv2RuleGroups(),
			"wafv2.web_acls":                        wafv2.Wafv2WebAcls(),
			"iot.things":                            iot.IotThings(),
			"iot.ca_certificates":                   iot.IotCaCertificates(),
			"iot.certificates":                      iot.IotCertificates(),
			"iot.streams":                           iot.IotStreams(),
			"iot.billing_groups":                    iot.IotBillingGroups(),
			"iot.thing_groups":                      iot.IotThingGroups(),
			"iot.thing_types":                       iot.IotThingTypes(),
			//todo disabled because of api error NotFoundException: No method found matching route security-profiles for http method GET.
			//"iot.security_profiles":                 iot.IotSecurityProfiles(),
			"iot.policies":    iot.IotPolicies(),
			"iot.topic_rules": iot.IotTopicRules(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
