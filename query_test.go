package main

import (
	"reflect"
	"testing"
)

func TestGetItem(t *testing.T) {
	id := "99"
	want := &Item{
		Id:    "99",
		Title: "寻找峰值",
		Question: struct {
			QuestionId string
			Title      string
			TitleSlug  string
		}{
			"162",
			"Find Peak Element",
			"find-peak-element",
		},
	}
	got := getItem(id)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("with %s, got %+v, want %+v\n", id, want, got)
	}
}

func TestGetQuestion(t *testing.T) {
	titleSlug := "3sum"
	want := &Question{
		QuestionId: "15",
	}
	got := getQuestion(titleSlug)
	if want.QuestionId != got.QuestionId {
		t.Errorf("with %s, got %+v, want %+v\n", titleSlug, want, got)
	}
}

func TestGetCard(t *testing.T) {
	cardSlug := "top-interview-questions-medium"
	want := &Card{
		[]struct{
			Id string
	}{
			{"29"},
			{"31"},
			{"32"},
			{"49"},
			{"50"},
			{"51"},
			{"52"},
			{"53"},
			{"54"},
		},
	}
	got := getCard(cardSlug)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("with %s, got %+v, want %+v\n", cardSlug, want, got)
	}
}

func TestGetChapter(t *testing.T) {
	chapterId := "29"
	cardSlug := "top-interview-questions-medium"
	want := &Chapter{
		"array-and-strings",
		[]struct{
			Id string
		} {
			{"75"},
			{"76"},
			{"77"},
			{"78"},
			{"79"},
			{"80"},
		},
	}
	got := getChapter(chapterId, cardSlug)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("with %s,%s, got %+v, want %+v\n", chapterId, cardSlug, want, got)
	}
}
