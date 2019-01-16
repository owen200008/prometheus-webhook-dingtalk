package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/timonwong/prometheus-webhook-dingtalk/models"
	"github.com/timonwong/prometheus-webhook-dingtalk/template"
	"net/http"
	"strings"
)

func BuildDingTalkNotification(promMessage *models.WebhookMessage) (*models.DingTalkNotification, error) {
	title, err := template.ExecuteTextString(`{{ template "ding.link.title" . }}`, promMessage)
	if err != nil {
		return nil, err
	}
	content, err := template.ExecuteTextString(`{{ template "ding.link.content" . }}`, promMessage)
	if err != nil {
		return nil, err
	}

	ayMobiles, err := template.ExecuteMultiString(`{{ template "ding.link.at" . }}`, promMessage)
	if err != nil {
		return nil, err
	}
	sizeMobile := len(ayMobiles)
	ayMobilesContent := ""
	for i := 0; i < sizeMobile; i++ {
		ayMobiles[i] = strings.TrimSpace(ayMobiles[i])
		ayMobiles[i] = strings.Trim(ayMobiles[i], "\r\n")
		ayMobiles[i] = strings.Trim(ayMobiles[i], "\n")
		ayMobilesContent += "@" + ayMobiles[i] + " "
	}

	setAtAllUse := false
	if sizeMobile == 0 || (sizeMobile == 1 && len(ayMobiles[0]) == 0) {
		setAtAllUse = true
		ayMobilesContent = ""
	} else {
		ayMobilesContent += "\r\n"
	}

	var buttons []models.DingTalkNotificationButton
	for i, alert := range promMessage.Alerts.Firing() {
		buttons = append(buttons, models.DingTalkNotificationButton{
			Title:     fmt.Sprintf("Graph for alert #%d", i+1),
			ActionURL: alert.GeneratorURL,
		})
	}

	notification := &models.DingTalkNotification{
		MessageType: "markdown",
		Markdown: &models.DingTalkNotificationMarkdown{
			Title: title,
			Text:  ayMobilesContent + content,
		},
		At: &models.DingTalkNotificationAt{
			AtMobiles: ayMobiles,
			IsAtAll:   setAtAllUse,
		},
	}
	return notification, nil
}

func SendDingTalkNotification(httpClient *http.Client, webhookURL string, notification *models.DingTalkNotification) (*models.DingTalkNotificationResponse, error) {
	body, err := json.Marshal(&notification)
	if err != nil {
		return nil, errors.Wrap(err, "error encoding DingTalk request")
	}

	httpReq, err := http.NewRequest("POST", webhookURL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building DingTalk request")
	}
	httpReq.Header.Set("Content-Type", "application/json")

	req, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "error sending notification to DingTalk")
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		return nil, errors.Errorf("unacceptable response code %d", req.StatusCode)
	}

	var robotResp models.DingTalkNotificationResponse
	enc := json.NewDecoder(req.Body)
	if err := enc.Decode(&robotResp); err != nil {
		return nil, errors.Wrap(err, "error decoding response from DingTalk")
	}

	return &robotResp, nil
}
