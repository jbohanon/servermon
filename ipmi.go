package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func resetServer(s *Settings) error {
	payload := bytes.NewBuffer([]byte(s.IpmiResetPayload))
	req, err := http.NewRequest("POST", s.IpmiResetUrl, payload)
	if err != nil {
		err = errors.Wrap(err, "creating request to reset server over IPMI")
		return err
	}
	req.AddCookie(&http.Cookie{
		Name:  "SID",
		Value: s.SidCookie,
	})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "making http request to reset server over IPMI")
		return err
	}
	defer resp.Body.Close()

	header := resp.Header.Get("content-type")
	if header != "application/xml" {
		return errors.Errorf("reset request returned incorrect content-type: %s", header)
	}

	time.Sleep(time.Minute * 10)

	return nil
}

func loginIpmi(s *Settings) error {
	payload := bytes.NewBuffer([]byte(fmt.Sprintf(s.IpmiResetPayload, s.IpmiUser, s.IpmiPassword)))
	req, err := http.NewRequest("POST", s.IpmiResetUrl, payload)
	if err != nil {
		err = errors.Wrap(err, "creating request to login to IPMI")
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.Wrap(err, "making http request to reset server over IPMI")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return errors.Errorf("login returned error %s", resp.Status)
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name != "SID" || cookie.Expires.Before(time.Now()) {
			continue
		}
		s.SidCookie = cookie.Value
	}
	return nil
}
