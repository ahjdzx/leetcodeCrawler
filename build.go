package main

import (
	"fmt"
	"path"
)

func buildProblem(id, basePath string) {
	item := getItem(id)
	titleSlug := item.Question.TitleSlug
	question := getQuestion(titleSlug)
	err := createQuestion(basePath, question)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func buildAllProblems() {
	// 一键创建所有题目，按难度和类型分文件夹保存
	cardSlugs := []string{
		"top-interview-questions-easy",
		"top-interview-questions-medium",
		"top-interview-questions-hard",
	}
	for _, cardSlug := range cardSlugs {
		cardFolder := path.Join("./" + cardSlug)
		err := createFolder(cardFolder)
		if err != nil {
			fmt.Println(err)
			return
		}
		card := getCard(cardSlug)
		for _, c := range card.Chapters {
			chapter := getChapter(c.Id, cardSlug)
			chapterFolder := path.Join(cardFolder + "/" + chapter.Slug)
			err := createFolder(chapterFolder)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, item := range chapter.Items {
				buildProblem(item.Id, chapterFolder)
			}
		}
	}
}
