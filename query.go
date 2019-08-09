package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const ApiUrl = "https://leetcode-cn.com/graphql"
const ContentType = "application/json"

const getItemTmpl = `{"operationName":"GetItem","variables":{"itemId":"%s"},"query":"query GetItem($itemId: String!) {\n  item(id: $itemId) {\n    id\n    title\n    type\n    paidOnly\n    lang\n    question {\n      questionId\n      title\n      titleSlug\n      __typename\n    }\n    article {\n      id\n      title\n      __typename\n    }\n    video {\n      id\n      __typename\n    }\n    htmlArticle {\n      id\n      __typename\n    }\n    webPage {\n      id\n      __typename\n    }\n    __typename\n  }\n  isCurrentUserAuthenticated\n}\n"}`

// const getQuestionTmpl = `{"operationName":"GetQuestion","variables":{"titleSlug":"%s"},"query":"query GetQuestion($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    sessionId\n    questionTitle\n    categoryTitle\n    submitUrl\n    interpretUrl\n    codeDefinition\n    sampleTestCase\n    enableTestMode\n    metaData\n    langToValidPlayground\n    enableRunCode\n    enableSubmit\n    judgerAvailable\n    infoVerified\n    envInfo\n    content\n    translatedContent\n    urlManager\n    __typename\n  }\n  isCurrentUserAuthenticated\n}\n"}`
const getCardDetailTmpl = `{"operationName":"GetExtendedCardDetail","variables":{"cardSlug":"%s"},"query":"query GetExtendedCardDetail($cardSlug: String!) {\n  card(cardSlug: $cardSlug) {\n    id\n    title\n    slug\n    description\n    introduction\n    chapters {\n      id\n      __typename\n    }\n    __typename\n  }\n}\n"}`
const getChapterTmpl = `{"operationName":"GetChapter","variables":{"chapterId":"%s","cardSlug":"%s"},"query":"query GetChapter($chapterId: String, $cardSlug: String) {\n  chapter(chapterId: $chapterId, cardSlug: $cardSlug) {\n    ...ExtendedChapterDetail\n    description\n    __typename\n  }\n}\n\nfragment ExtendedChapterDetail on ChapterNode {\n  id\n  title\n  slug\n  items {\n    id\n    title\n    type\n    info\n    paidOnly\n    chapterId\n    prerequisites {\n      id\n      chapterId\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n"}`
const questionDataImpl = `{
    "operationName": "questionData",
    "variables": {
        "titleSlug": "%s"
    },
    "query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n  }\n}\n"
}`

type Item struct {
	Id       string
	Title    string
	Question struct {
		QuestionId string
		Title      string
		TitleSlug  string
	}
}

func getItem(id string) *Item {
	body := fmt.Sprintf(getItemTmpl, id)
	decoder := postRequestWith(body)
	wrapper := struct {
		Data struct {
			Item Item
		}
	}{}
	err := decoder.Decode(&wrapper)
	if err != nil {
		log.Fatal(err)
	}
	return &wrapper.Data.Item
}

type Question struct {
	QuestionId        string
	BoundTopicId      int
	Title             string
	TitleSlug         string
	Content           string
	TranslatedTitle   string
	TranslatedContent string
	IsPaidOnly        bool
	Difficulty        string
	TopicTags         []TopicTag
	CodeSnippets      []CodeSnippet
	SampleTestCase    string
}

type TopicTag struct {
	Name           string
	Slug           string
	TranslatedName string
}

type CodeSnippet struct {
	Lang     string
	LangSlug string
	Code     string
}

func getQuestionData(titleSlug string) *Question {
	body := fmt.Sprintf(questionDataImpl, titleSlug)
	decoder := postRequestWith(body)
	wrapper := struct {
		Data struct {
			Question Question
		}
	}{}
	err := decoder.Decode(&wrapper)
	if err != nil {
		log.Fatal(err)
	}
	return &wrapper.Data.Question
}

type Card struct {
	Chapters []struct {
		Id string
	}
}

func getCard(cardSlug string) *Card {
	body := fmt.Sprintf(getCardDetailTmpl, cardSlug)
	decoder := postRequestWith(body)
	wrapper := struct {
		Data struct {
			Card Card
		}
	}{}
	err := decoder.Decode(&wrapper)
	if err != nil {
		log.Fatal(err)
	}
	return &wrapper.Data.Card
}

type Chapter struct {
	Slug  string
	Items []struct {
		Id string
	}
}

func getChapter(chapterId, cardSlug string) *Chapter {
	body := fmt.Sprintf(getChapterTmpl, chapterId, cardSlug)
	decoder := postRequestWith(body)
	wrapper := struct {
		Data struct {
			Chapter Chapter
		}
	}{}
	err := decoder.Decode(&wrapper)
	if err != nil {
		log.Fatal(err)
	}
	return &wrapper.Data.Chapter
}

func postRequestWith(requestBody string) *json.Decoder {
	bodyReader := strings.NewReader(requestBody)
	resp, err := http.Post(ApiUrl, ContentType, bodyReader)
	if err != nil {
		log.Printf("error when query: %T:%v\n", err, err)
		return nil
	}
	return json.NewDecoder(resp.Body)
}

func getAllProblems() *AllProblems {
	endpoint := "https://leetcode-cn.com/api/problems/all/"
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Printf("error: %T:%v\n", err, err)
		return nil
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	allProblems := AllProblems{}
	if err = decoder.Decode(&allProblems); err != nil {
		log.Printf("error: %T:%v\n", err, err)
		return nil
	}
	return &allProblems
}

type AllProblems struct {
	StatStatusPairs StatusPairs `json:"stat_status_pairs"`
}

type StatStatusPair struct {
	Stat struct {
		QuestionID        int    `json:"question_id"`
		QuestionTitleSlug string `json:"question__title_slug"`
	} `json:"stat"`
}

type StatusPairs []StatStatusPair

func (s StatusPairs) Len() int {
	return len(s)
}

func (s StatusPairs) Less(i, j int) bool {
	return s[i].Stat.QuestionID < s[j].Stat.QuestionID
}

func (s StatusPairs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
