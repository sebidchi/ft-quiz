package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sebidchi/ft-quiz/internal/pkg/infrastructure/question"
	"github.com/spf13/cobra"

	"github.com/AlecAivazis/survey/v2"
)

var username string

var rootCmd = &cobra.Command{
	Use:   "ft-quiz-cli",
	Short: "CLI for taking a quiz",
	Long:  `A command line interface for taking a quiz and checking your score.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Quiz CLI")
		fmt.Print("Enter your username: ")
		fmt.Scanln(&username)

		questions, err := getQuestions()
		if err != nil {
			fmt.Println("Error fetching questions:", err)
			os.Exit(1)
		}

		var answers []question.UserAnswer

		for i := 0; i < len(questions); i++ {
			qt := questions[i]
			answerIndex := 0
			opts := make([]string, 0)
			for j := 0; j < len(qt.Options); j++ {
				opts = append(opts, qt.Options[j].Text)
			}
			prompt := &survey.Select{
				Message: qt.Question,
				Options: opts,
			}
			survey.AskOne(prompt, &answerIndex)
			answers = append(answers, question.UserAnswer{QuestionID: qt.AnswerOptionId, Answer: qt.Options[answerIndex].OptionId})
		}

		err = postAnswers(answers)
		if err != nil {
			fmt.Println("Error submitting answers:", err)
			os.Exit(1)
		}

		results, err := getUserResults()
		if err != nil {
			fmt.Println("Error fetching results:", err)
			os.Exit(1)
		}

		fmt.Printf("You got %d out of %d correct. %.2f%% of success rate.\n", results.Total, len(questions), results.Percentage)
		fmt.Printf("You were better than %.2f%% of total of %d quizzers.\n", results.BetterThan, results.TotalUsers)
	},
}

func getQuestions() ([]question.Question, error) {
	resp, err := http.Get("http://localhost:8087/quiz")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	questions := make([]question.Question, 0)
	err = json.Unmarshal(body, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func postAnswers(answers []question.UserAnswer) error {
	answersPayload := question.AnswersPayload{
		Username: username,
		Answers:  answers,
	}
	answersJSON, err := json.Marshal(answersPayload)
	if err != nil {
		return err
	}

	_, err = http.Post("http://localhost:8087/answers", "application/json", bytes.NewBuffer(answersJSON))

	return err
}

func getUserResults() (question.UserResults, error) {
	var userResults question.UserResults
	resp, err := http.Get(fmt.Sprintf("http://localhost:8087/answers/%s", username))
	if err != nil {
		return userResults, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return userResults, err
	}

	if err = json.Unmarshal(body, &userResults); err != nil {
		return userResults, err
	}

	return userResults, nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for the quiz")
}
func main() {
	cobra.CheckErr(rootCmd.Execute())
}
