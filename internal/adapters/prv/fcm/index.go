package fcm

import (
	"context"
	"fmt"

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
		Data:   obj.Data,
		Tokens: obj.Tokens,
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
