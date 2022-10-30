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
	var err error

	const chunkSize = 500

	tokens := obj.Tokens
	obj.Tokens = nil

	l := len(tokens)
	i := 0
	for i < l {
		if i+chunkSize >= l {
			err = o.sendChunk(tokens[i:], obj)
		} else {
			err = o.sendChunk(tokens[i:], obj)
		}
		if err != nil {
			return err
		}

		i += chunkSize
	}

	return nil
}

func (o *St) sendChunk(tokens []string, obj *prv.SendReqSt) error {
	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Data:   obj.Data,
		Notification: &messaging.Notification{
			Title: obj.Title,
			Body:  obj.Body,
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
