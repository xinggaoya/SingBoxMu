package utils

import (
	"changeme/app/consts"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

/**
  @author: XingGao
  @date: 2024/8/9
**/

type TaskUtils struct{}

func NewTaskUtils() *TaskUtils {
	return &TaskUtils{}
}

func (t *TaskUtils) CreateTask() error {

	// Get the executable path of the current program
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error getting executable path: %v", err)
	}
	exePath, err = filepath.Abs(exePath)
	exePath = fmt.Sprintf(`"%s" -hide`, exePath)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	// Ensure the executable path is correctly quoted
	createTaskCmd := fmt.Sprintf(`schtasks /create /tn %s /tr %s /sc onlogon /rl HIGHEST`, consts.TaskName, exePath)
	cmd := exec.Command("cmd", "/C", createTaskCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creating task: %v, output: %s", err, output)
	}

	slog.Info("Task created successfully with highest privileges")
	return nil
}

// RunTask 运行
func (t *TaskUtils) RunTask() error {
	// Ensure the task name is correctly quoted
	runTaskCmd := fmt.Sprintf(`schtasks /run /tn %s`, consts.TaskName)
	cmd := exec.Command("cmd", "/C", runTaskCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running task: %v,output: %s", err, output)
	}
	slog.Info("Task running successfully")
	return nil
}

// CheckTaskExists checks if a scheduled task exists
func (t *TaskUtils) CheckTaskExists() (bool, error) {
	checkTaskCmd := fmt.Sprintf(`schtasks /query /tn "%s"`, consts.TaskName)
	cmd := exec.Command("cmd", "/C", checkTaskCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		// If the task does not exist, schtasks returns an error
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			if exitError.ExitCode() == 1 {
				return false, nil
			}
		}
		return false, fmt.Errorf("error checking task: %v, output: %s", err, output)
	}

	if output != nil {
		return true, nil
	}
	return true, nil
}

// StopTask 停止
func (t *TaskUtils) StopTask() error {
	// Ensure the task name is correctly quoted
	stopTaskCmd := fmt.Sprintf(`schtasks /end /tn %s`, consts.TaskName)
	cmd := exec.Command("cmd", "/C", stopTaskCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error stopping task: %v, output: %s", err, output)
	}

	fmt.Println("Task stopped successfully")
	return nil
}

// DeleteTask deletes a scheduled task
func (t *TaskUtils) DeleteTask() error {
	// Ensure the task name is correctly quoted
	deleteTaskCmd := fmt.Sprintf(`schtasks /delete /tn %s /f`, consts.TaskName)
	cmd := exec.Command("cmd", "/C", deleteTaskCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error deleting task: %v, output: %s", err, output)
	}

	fmt.Println("Task deleted successfully")
	return nil
}
