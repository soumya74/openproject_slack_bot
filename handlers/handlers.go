package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	constants "bitbucket.org/slackBot/common/constants"
	models "bitbucket.org/slackBot/models"
)

// health handler
func healthTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}

//webhook handler
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var workPackage models.WorkPackage
	err := json.NewDecoder(r.Body).Decode(&workPackage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	if workPackage.Action != constants.StatusCreated {
		return
	}
	sendPostCallToSlack(workPackage)
}

func getSlackMessageText(workPackage models.WorkPackage) string {
	issueID := strconv.Itoa(workPackage.WorkPackage.ID)
	issueSub, err := url.QueryUnescape(workPackage.WorkPackage.Subject)
	if err != nil {
		issueSub = workPackage.WorkPackage.Description.Raw
		fmt.Println(err)
	}

	messageStr := "*Link :* " + constants.OpenProjectCompanyURL + issueID +
		"\n *Summary :* " + issueSub +
		"\n *Assignee :* " + workPackage.WorkPackage.Embedded.Assignee.Name +
		"\t *Author :* " + workPackage.WorkPackage.Embedded.Author.Name +
		"\n *Priority :* " + workPackage.WorkPackage.Embedded.Priority.Name +
		"\t *Version :* " + workPackage.WorkPackage.Embedded.Version.Name
	fmt.Println(messageStr)
	return messageStr
}

func getChannelName(workPackage models.WorkPackage) string {
	versionName := workPackage.WorkPackage.Embedded.Version.Name
	channelName := "#" + versionName
	if versionName == "" {
		channelName = constants.ChannelPsIssues
	}
	return channelName
}

func sendPostCallToSlack(workPackage models.WorkPackage) {
	hc := http.Client{}
	form := url.Values{}
	form.Add("token", constants.OpBotToken)
	form.Add("channel", getChannelName(workPackage))
	form.Add("text", getSlackMessageText(workPackage))

	req, _ := http.NewRequest("POST", constants.SlackChatURL, strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}
