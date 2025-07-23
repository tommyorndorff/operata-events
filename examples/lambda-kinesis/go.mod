module lambda-kinesis-example

go 1.24.5

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/tommyorndorff/operata-events v0.0.0
)

replace github.com/tommyorndorff/operata-events => ../../
