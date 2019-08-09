package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
)

const problemBaseCN = "https://leetcode-cn.com/problems/"
const problemBase = "https://leetcode.com/problems/"

func createQuestion(basePath string, question *Question) error {
	qid, _ := strconv.Atoi(question.QuestionId)
	titleSlug := question.TitleSlug

	name := fmt.Sprintf("%04d_%s", qid, titleSlug)
	folder := path.Join(basePath, name)
	err := createFolder(folder)
	if err != nil {
		return err
	}
	defaultCode := parseDefaultCode(question.CodeSnippets)
	err = createSolutionFile(folder, defaultCode)
	if err != nil {
		return err
	}
	err = createReadmeFile(folder, question)
	if err != nil {
		return err
	}
	log.Printf("create problem <%s> successfully!\n", name)
	return nil
}

func createFolder(folder string) error {
	err := os.Mkdir(folder, os.ModePerm)
	if err != nil {
		if os.IsExist(err) {
			log.Println("folder existed: ", folder)
			return err
		}
	}
	return nil
}

func createSolutionFile(folder, defaultCode string) error {
	solutionPath := path.Join(folder, "solution.go")
	f, err := os.Create(solutionPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %s\n", solutionPath)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "package leetcode")
	fmt.Fprintln(w)
	fmt.Fprint(w, defaultCode)
	return w.Flush()
}

func createReadmeFile(folder string, question *Question) error {
	readmePath := path.Join(folder, "README.md")
	f, err := os.Create(readmePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %s\n", readmePath)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "# "+question.Title)
	fmt.Fprintln(w)
	fmt.Fprintln(w, question.TranslatedTitle)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "**"+question.Difficulty+"**")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "## 问题描述")
	fmt.Fprintln(w)
	fmt.Fprint(w, question.TranslatedContent)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "## 相关标签")
	fmt.Fprintln(w)
	for _, topicTag := range question.TopicTags {
		fmt.Fprintln(w, "* "+topicTag.TranslatedName)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "## 测试用例")
	fmt.Fprintln(w)
	fmt.Fprintln(w, question.SampleTestCase)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "-----")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "## 链接")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "中文："+problemBaseCN+question.TitleSlug)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "英文："+problemBase+question.TitleSlug)
	return w.Flush()
}

func parseDefaultCode(codeSnippets []CodeSnippet) string {
	for _, item := range codeSnippets {
		if item.Lang == "Go" {
			return item.Code
		}
	}
	return ""
}

type CodeDefinition struct {
	Value       string
	Text        string
	DefaultCode string
}
