package resource

import (
	"log"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("Create with current model: %+v, previous model %+v\n", currentModel, prevModel)
	ssmClient := ssm.New(req.Session)
	tags := make([]*ssm.Tag, 0)
	for _, t := range currentModel.Tags {
		tags = append(tags, &ssm.Tag{
			Key:   t.Key.Value(),
			Value: t.Value.Value(),
		})
	}

	_, err := ssmClient.PutParameter(&ssm.PutParameterInput{
		AllowedPattern: currentModel.AllowedPattern.Value(),
		Description:    currentModel.Description.Value(),
		KeyId:          currentModel.KeyId.Value(),
		Name:           currentModel.Name.Value(),
		Overwrite:      aws.Bool(false),
		Policies:       currentModel.Policies.Value(),
		Tags:           tags,
		Tier:           currentModel.Tier.Value(),
		Type:           aws.String(ssm.ParameterTypeSecureString),
		Value:          currentModel.Value.Value(),
	})

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	if currentModel.Name.Value() != prevModel.Name.Value() {
		_, _ = ssmClient.DeleteParameter(&ssm.DeleteParameterInput{Name: prevModel.Name.Value()})
	}

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("Read with current model: %+v, previous model %+v\n", currentModel, prevModel)
	client := ssm.New(req.Session)

	parameter, err := client.GetParameter(&ssm.GetParameterInput{
		Name:           currentModel.Name.Value(),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Assign the value
	currentModel.Value = encoding.NewString(*parameter.Parameter.Value)

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("Update with current model: %+v, previous model %+v\n", currentModel, prevModel)
	client := ssm.New(req.Session)

	_, err := client.PutParameter(&ssm.PutParameterInput{
		AllowedPattern: currentModel.AllowedPattern.Value(),
		Description:    currentModel.Description.Value(),
		KeyId:          currentModel.KeyId.Value(),
		Name:           currentModel.Name.Value(),
		Overwrite:      aws.Bool(true),
		Policies:       currentModel.Policies.Value(),
		Tier:           currentModel.Tier.Value(),
		Type:           aws.String(ssm.ParameterTypeSecureString),
		Value:          currentModel.Value.Value(),
	})

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("Delete with current model: %+v, previous model %+v\n", currentModel, prevModel)
	client := ssm.New(req.Session)

	_, err := client.DeleteParameter(&ssm.DeleteParameterInput{
		Name: currentModel.Name.Value(),
	})

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	log.Printf("List with current model: %+v, previous model %+v\n", currentModel, prevModel)
	client := ssm.New(req.Session)

	parameter, err := client.GetParameter(&ssm.GetParameterInput{
		Name:           currentModel.Name.Value(),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Assign the value
	currentModel.Value = encoding.NewString(*parameter.Parameter.Value)

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil

}
