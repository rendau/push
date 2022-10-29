package fcm

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/rendau/dop/adapters/logger"
	"github.com/rendau/push/internal/adapters/prv"
	"google.golang.org/api/option"
)

type St struct {
	lg logger.Lite

	app    *firebase.App
	client *messaging.Client
}

func New(lg logger.Lite, credsPath string) (*St, error) {
	var err error

	ctx := context.Background()

	res := &St{
		lg: lg,
	}

	res.app, err = firebase.NewApp(ctx, nil, option.WithCredentialsFile(credsPath))
	if err != nil {
		lg.Errorw("error initializing fcm-app", err)
		return nil, err
	}

	res.client, err = res.app.Messaging(ctx)
	if err != nil {
		lg.Errorw("error getting Messaging client", err)
		return nil, err
	}

	return res, nil
}

func (o *St) Send(obj *prv.SendReqSt) error {
	/*
		type MulticastMessage struct {
			Tokens       []string
			Data         map[string]string
			Notification *Notification
			Android      *AndroidConfig
			Webpush      *WebpushConfig
			APNS         *APNSConfig
		}

		type Message struct {
			Token        string            `json:"token,omitempty"`
			Data         map[string]string `json:"data,omitempty"`
			Notification *Notification     `json:"notification,omitempty"`
			Android      *AndroidConfig    `json:"android,omitempty"`
			Webpush      *WebpushConfig    `json:"webpush,omitempty"`
			APNS         *APNSConfig       `json:"apns,omitempty"`

			FCMOptions   *FCMOptions       `json:"fcm_options,omitempty"`
			Topic        string            `json:"-"`
			Condition    string            `json:"condition,omitempty"`
		}
	*/

	message := &messaging.MulticastMessage{
		Tokens: obj.Tokens,
		Data:   obj.Data,
		Notification: &messaging.Notification{
			Title: obj.Title,
			Body:  obj.Body,
		},
		Android: &messaging.AndroidConfig{
			CollapseKey:           "",
			Priority:              "",
			TTL:                   nil,
			RestrictedPackageName: "",
			Data:                  nil,
			Notification: &messaging.AndroidNotification{
				Title:                 "",
				Body:                  "",
				Icon:                  "",
				Color:                 "",
				Sound:                 "",
				Tag:                   "",
				ClickAction:           "",
				BodyLocKey:            "",
				BodyLocArgs:           nil,
				TitleLocKey:           "",
				TitleLocArgs:          nil,
				ChannelID:             "",
				ImageURL:              "",
				Ticker:                "",
				Sticky:                false,
				EventTimestamp:        &time.Time{},
				LocalOnly:             false,
				Priority:              0,
				VibrateTimingMillis:   nil,
				DefaultVibrateTimings: false,
				DefaultSound:          false,
				LightSettings: &messaging.LightSettings{
					Color:                  "",
					LightOnDurationMillis:  0,
					LightOffDurationMillis: 0,
				},
				DefaultLightSettings: false,
				Visibility:           0,
				NotificationCount:    nil,
			},
			FCMOptions: &messaging.AndroidFCMOptions{
				AnalyticsLabel: "",
			},
		},
		APNS: &messaging.APNSConfig{
			Headers: nil,
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					AlertString: "",
					Alert: &messaging.ApsAlert{
						Title:           "",
						SubTitle:        "",
						Body:            "",
						LocKey:          "",
						LocArgs:         nil,
						TitleLocKey:     "",
						TitleLocArgs:    nil,
						SubTitleLocKey:  "",
						SubTitleLocArgs: nil,
						ActionLocKey:    "",
						LaunchImage:     "",
					},
					Badge: nil,
					Sound: "",
					CriticalSound: &messaging.CriticalSound{
						Critical: false,
						Name:     "",
						Volume:   0,
					},
					ContentAvailable: false,
					MutableContent:   false,
					Category:         "",
					ThreadID:         "",
					CustomData:       nil,
				},
				CustomData: nil,
			},
			FCMOptions: &messaging.APNSFCMOptions{
				AnalyticsLabel: "",
				ImageURL:       "",
			},
		},
	}

	respObj, err := o.client.SendMulticast(context.Background(), message)
	if err != nil {
		return err
	}

	for _, rsp := range respObj.Responses {
		fmt.Printf("%#v\n", rsp.Error)
	}

	return nil
}
