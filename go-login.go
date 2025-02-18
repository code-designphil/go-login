package main

import (
	"github.com/DataDog/datadog-cdk-constructs-go/ddcdkconstruct"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GoLoginStackProps struct {
	awscdk.StackProps
}

func NewGoLoginStack(scope constructs.Construct, id string, props *GoLoginStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	table := awsdynamodb.NewTable(stack, jsii.String("UserTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("userTable"),
	})

	datadog := ddcdkconstruct.NewDatadogLambda(stack, jsii.String("Datadog"), &ddcdkconstruct.DatadogLambdaProps{
		ApiKeySecretArn:       jsii.String("arn:aws:secretsmanager:eu-north-1:537461418918:secret:DdApiKeySecret-tyU7cYt920oQ-A7fcZ5"),
		Site:                  jsii.String("datadoghq.eu"),
		Service:               jsii.String("go-login"),
		EnableDatadogASM:      jsii.Bool(true),
		EnableDatadogLogs:     jsii.Bool(true),
		EnableDatadogTracing:  jsii.Bool(true),
		ExtensionLayerVersion: jsii.Number(1),
		InjectLogContext:      jsii.Bool(true),
		SourceCodeIntegration: jsii.Bool(true),
	})

	lambda := awslambda.NewFunction(stack, jsii.String("GoLoginFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	table.GrantReadWriteData(lambda)
	lambdaFunctions := []interface{}{lambda}
	datadog.AddLambdaFunctions(&lambdaFunctions, stack)

	api := awsapigateway.NewRestApi(stack, jsii.String("GoLoginApiGateway"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			AllowMethods: jsii.Strings("GET", "POST", "DELETE", "PUT", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
		DeployOptions: &awsapigateway.StageOptions{
			LoggingLevel: awsapigateway.MethodLoggingLevel_INFO,
		},
		CloudWatchRole: jsii.Bool(true),
	})

	integration := awsapigateway.NewLambdaIntegration(lambda, nil)

	registerResource := api.Root().AddResource(jsii.String("register"), nil)
	registerResource.AddMethod(jsii.String("POST"), integration, nil)
	loginResource := api.Root().AddResource(jsii.String("login"), nil)
	loginResource.AddMethod(jsii.String("POST"), integration, nil)
	protected := api.Root().AddResource(jsii.String("protected"), nil)
	protected.AddMethod(jsii.String("GET"), integration, nil)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGoLoginStack(app, "GoLoginStack", &GoLoginStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
