package core

/*
const requestRetryCount = 2
const requestRetryInterval = 3 * time.Second

var (
	httpClient = &http.Client{Timeout: 20 * time.Second}
)

func (c *St) Send(reqSt *entities.SendReqSt) error {
	var err error

	// validation
	if len(reqSt.UsrIds) == 0 {
		c.lg.Errorw("Bad user ids", errs.BadUserIds)
		return errs.BadUserIds
	}

	err = c.sendToPlatform(cns.PlatformUndefined, reqSt)
	if err != nil {
		return err
	}

	err = c.sendToPlatform(cns.PlatformIOS, reqSt)
	if err != nil {
		return err
	}

	err = c.sendToPlatform(cns.PlatformAndroid, reqSt)
	if err != nil {
		return err
	}

	err = c.sendToPlatform(cns.PlatformWeb, reqSt)
	if err != nil {
		return err
	}

	return nil
}

func (c *St) sendToPlatform(platformId int, reqSt *entities.SendReqSt) error {
	var err error

	tokens, err := c.repo.GetTokens(platformId, reqSt.UsrIds)
	if err != nil {
		c.lg.Errorw("Cant get tokens from repo", err)
		return errs.ServerNA
	}

	var chunk []string
	chunkSize := 1000
	l := len(tokens)
	i := 0
	for i < l {
		if i+chunkSize > l {
			chunk = tokens[i:]
		} else {
			chunk = tokens[i : i+chunkSize]
		}
		go c.sendChunk(platformId, chunk, reqSt)
		i += chunkSize
	}

	return nil
}

func (c *St) sendChunk(platformId int, tokens []string, reqSt *entities.SendReqSt) {
	var err error

	commonMessage := entities.CommonMessageSt{
		Tokens:     tokens,
		TimeToLive: 86400,
		Priority:   "high",
		Data:       reqSt.Data,
	}

	commonNotification := entities.CommonNotificationSt{
		Title: reqSt.Title,
		Body:  reqSt.Body,
	}

	var msg interface{}

	switch platformId {
	case cns.PlatformUndefined:
		msg = entities.UndefinedMessageSt{
			CommonMessageSt: commonMessage,
			Notification: &entities.UndefinedNotificationSt{
				CommonNotificationSt: commonNotification,
				Sound:                "default",
				Badge:                strconv.Itoa(reqSt.Badge),
				ClickAction:          "FCM_PLUGIN_ACTIVITY",
			},
		}
	case cns.PlatformAndroid:
		msg = entities.AndroidMessageSt{
			CommonMessageSt: commonMessage,
			Notification: &entities.AndroidNotificationSt{
				CommonNotificationSt: commonNotification,
				Tag:                  reqSt.AndroidTag,
				Sound:                "default",
				ClickAction:          "FCM_PLUGIN_ACTIVITY",
			},
		}
	case cns.PlatformIOS:
		msg = entities.IosMessageSt{
			CommonMessageSt: commonMessage,
			Notification: &entities.IosNotificationSt{
				CommonNotificationSt: commonNotification,
				Sound:                "default",
				Badge:                strconv.Itoa(reqSt.Badge),
			},
		}
	case cns.PlatformWeb:
		commonMessage.Data["_n_title"] = reqSt.Title
		commonMessage.Data["_n_body"] = reqSt.Body

		msg = entities.WebMessageSt{
			CommonMessageSt: commonMessage,
		}
	}

	err = c.sendMsg(tokens, msg)
	if err != nil {
		return
	}

	if reqSt.Badge > 0 && platformId == cns.PlatformAndroid { // push without notification, just for badge
		badgeMsg := entities.AndroidMessageSt{
			CommonMessageSt: commonMessage,
		}

		badgeMsg.Data = map[string]string{
			"type":          "android_badge",
			"android_badge": strconv.Itoa(reqSt.Badge),
		}

		err = c.sendMsg(tokens, badgeMsg)
		if err != nil {
			return
		}
	}
}

func (c *St) sendMsg(tokens []string, msg interface{}) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		c.lg.Errorw("Cant marshal json", errs.JSONParseErr)
		return err
	}

	// fmt.Println(string(msgBytes))

	var replyObj *entities.FcmReplySt

	for i := 0; i < requestRetryCount; i++ {
		replyObj, err = c.sendFcmReq(msgBytes)
		if err == nil {
			break
		}

		time.Sleep(requestRetryInterval)
	}

	if err == nil {
		badTokens := make([]string, 0, len(tokens))

		for i, result := range replyObj.Results {
			if result.Error == "InvalidRegistration" || result.Error == "NotRegistered" || result.Error == "MismatchSenderId" {
				badTokens = append(badTokens, tokens[i])
			}
		}
		if len(badTokens) > 0 {
			err = c.repo.DeleteTokens(badTokens)
			if err != nil {
				c.lg.Errorw("Cant delete bad tokens", err)
				return err
			}
		}
	} else {
		c.lg.Errorw("(fcm) fail -", err)
		return err
	}

	return nil
}

func (c *St) sendFcmReq(msgBytes []byte) (*entities.FcmReplySt, error) {
	var err error

	req, err := http.NewRequest("POST", cns.FcmSendUrl, bytes.NewBuffer(msgBytes))
	if err != nil {
		c.lg.Errorw("Fail to create http-request", err)
		return nil, err
	}
	req.Header.Add("Authorization", "key="+c.fcmServerKey)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	rep, err := httpClient.Do(req)
	if err != nil {
		c.lg.Errorw("Fail to send http-request", err)
		return nil, err
	}
	defer rep.Body.Close()

	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		body, err := ioutil.ReadAll(rep.Body)
		if err != nil {
			body = []byte{}
		}
		c.lg.Errorw("Fail to send http-request, bad status code", nil, "status_code", rep.StatusCode, "body", string(body))
		return nil, errors.New("bad_status_code")
	}

	replyObj := &entities.FcmReplySt{}

	if err = json.NewDecoder(rep.Body).Decode(replyObj); err != nil {
		c.lg.Errorw("Fail to parse http-body", err)
		return nil, err
	}

	return replyObj, nil
}
*/
