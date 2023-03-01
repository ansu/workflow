package orchestrator

import (
	"context"
	"fmt"
	"log"

	"example.com/WORKFLOW/tasks"
	// "github.com/your-organization/workflow-service/models"
	// "github.com/your-organization/workflow-service/services"
)

type Orchestrator struct {
	// userWorkflowTaskRepo     repository.UserWorkflowTaskRepository
	// userWorkflowRunTimeRepo  repository.UserWorkflowRunTimeDataRepository
	// workflowRepo             repository.WorkflowRepository
	serviceTaskExecutor tasks.ServiceTaskExecutor
}

// func NewOrchestrator(userWorkflowTaskRepo repository.UserWorkflowTaskRepository, userWorkflowRunTimeRepo repository.UserWorkflowRunTimeDataRepository, workflowRepo repository.WorkflowRepository, serviceTaskExecutor tasks.ServiceTaskExecutor) *Orchestrator {
// 	return &Orchestrator{
// 		userWorkflowTaskRepo:     userWorkflowTaskRepo,
// 		userWorkflowRunTimeRepo:  userWorkflowRunTimeRepo,
// 		workflowRepo:             workflowRepo,
// 		serviceTaskExecutor:      serviceTaskExecutor,

// 	}
// }

func NewOrchestrator(serviceTaskExecutor tasks.ServiceTaskExecutor) *Orchestrator {
	return &Orchestrator{

		serviceTaskExecutor: serviceTaskExecutor,
	}
}

func (o *Orchestrator) Dummy() {
	inputData := []interface{}{1, "test"}
	// Example 1: Call the "sum" method with input parameters 2 and 3
	result, err := o.serviceTaskExecutor.Execute("Digio", inputData)
	if err != nil {
		log.Fatalf("Error calling service task: %s", err.Error())
	}
	fmt.Println(result)

}

// HandleUserScreen handles the user's screen submission and returns the next screen data or service task data
func (o *Orchestrator) HandleUserScreen(ctx context.Context, userId string, screenData []byte) ([]byte, error) {

	return nil, nil
}

// func (o *Orchestrator) GetTaskForm(userId int, workflowId int) (*models.UserTask, error) {
// Check if the user has started the workflow
// workflowTask, err := o.userWorkflowTaskRepo.GetLastUserTaskByUserIdAndWorkflowId(userId, workflowId)
// if err != nil {
// 	return nil, errors.New("failed to get user workflow task")
// }
// if workflowTask == nil {
// 	return nil, errors.New("user has not started the workflow")
// }

// // Get the workflow definition
// workflow, err := o.workflowRepo.GetById(workflowId)
// if err != nil {
// 	return nil, errors.New("failed to get workflow")
// }

// // Find the task definition for the current user task
// taskDefinition, err := workflow.GetTaskDefinitionById(workflowTask.TaskId)
// if err != nil {
// 	return nil, errors.New("failed to get task definition")
// }

// // Build the user task form
// formFields := make([]*models.FormField, len(taskDefinition.Form.Fields))
// for i, field := range taskDefinition.Form.Fields {
// 	formFields[i] = &models.FormField{
// 		Name:     field.Name,
// 		Type:     field.Type,
// 		Required: field.Required,
// 		Value:    "",
// 	}
// }

// return &models.UserTask{
// 	TaskId:    workflowTask.TaskId,
// 	Form:      &models.UserTaskForm{Fields: formFields},
// 	CreatedAt: workflowTask.StartTime,
// }, nil
// }
