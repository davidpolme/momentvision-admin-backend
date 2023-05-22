package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendToSns(message string) {
	// Crear una nueva sesión de AWS
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Crear un cliente de SNS
	svc := sns.New(sess)

	// ID o ARN del tópico de SNS al que deseas enviar el mensaje
	topicARN := "arn:aws:sns:us-east-1:133440494183:ProgramsTopic"

	// Enviar el mensaje al tópico de SNS
	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topicARN),
	})
	if err != nil {
		fmt.Println("Error al enviar el mensaje:", err)
		return
	}

	// Imprimir el ID del mensaje publicado
	fmt.Println("Mensaje enviado con éxito. ID del mensaje:", *result.MessageId)
}
