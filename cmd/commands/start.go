package commands

import (
	"fmt"
	"strings"

	"github.com/dev-saw99/ioscript-cli/pkg/course"
	"github.com/dev-saw99/ioscript-cli/pkg/logger"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type CourseInput struct {
	CourseID   string // The ID of the course to start
	Level      string // The level of the course (e.g., Beginner, Intermediate, Advanced)
	CourseName string // The name of the course
}

var userInput CourseInput

func init() {
	startCmd.Flags().StringVarP(&userInput.CourseID, "course-id", "c", "", "ID of the course to start")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new course",
	Long:  `Start a new course by providing the course ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if userInput.CourseID != "" {
			// user has provided the course id, start the course
			courseName := course.GetCourseByCourseID(userInput.CourseID)
			if courseName == "" {
				logger.Error("Course not found", "courseID", userInput.CourseID, "courseName", courseName)
				fmt.Println("Uh-oh! Course not found. Please check the course ID and try again.")
				return
			}
			userInput.CourseName = courseName
			userInput.Level = strings.Split(courseName, ":")[0]
			startCourse(userInput)
			return
		}

		// user has not provided the course id, show the available courses
		levels := course.GetCoursesLevel()
		level, err := selectLevel(levels)
		if err != nil {
			logger.Error("Error selecting level", "error", err)
			fmt.Println("Uh-oh! Something went wrong. Please try again.")
			return
		}

		userInput.Level = level
		courses := course.GetCoursesByLevel(userInput.Level)
		if len(courses) == 0 {
			logger.Error("No courses found for the selected level", "level", userInput.Level)
			fmt.Println("Uh-oh! No courses found for the selected level. Please try again.")
			return
		}
		courseName, err := selectCourse(courses)
		if err != nil {
			logger.Error("Error selecting course", "error", err)
			fmt.Println("Uh-oh! Something went wrong. Please try again.")
			return
		}
		userInput.CourseName = courseName

		startCourse(userInput)
	},
}

func selectLevel(levels []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select your experience level",
		Items: levels,
	}

	_, level, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return level, nil
}

func selectCourse(courses []string) (string, error) {
	prompt := promptui.Select{
		Label: "Select a course",
		Items: courses,
	}

	_, courseName, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return courseName, nil
}

func startCourse(userInput CourseInput) {
	logger.Info("Starting course", "userInput", userInput)
	fmt.Printf("Starting course %s \n\n", userInput.CourseName)
	// Add  course starting logic here
}
