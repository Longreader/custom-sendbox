package main

import (
	"SandBox/internal/models"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Checkpoint struct {
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Infof("Start import")
	defer logrus.Infof("End import")

	var cp = new(Checkpoint)

	login, err := cp.Login("172.16.255.16", "admin", "checkeradmin")
	if err != nil {
		logrus.Errorf("Error login: %s", err)
		return
	}
	logrus.Infof("Login %s", login.Sid)

	defer func() {
		err = cp.Logout("172.16.255.16", login.Sid)
		if err != nil {
			logrus.Errorf("Error logout: %s", err)
		}
	}()

	srvGroup, err := cp.GetServiceGroup("172.16.255.16", login.Sid)
	if err != nil {
		logrus.Errorf("Error get service group: %s", err)
		return
	}
	logrus.Infof("Service group %v", srvGroup)

	netGroup, err := cp.GetNetworkGroup("172.16.255.16", login.Sid)
	if err != nil {
		logrus.Errorf("Error get network group: %s", err)
		return
	}
	logrus.Infof("Network group %v", netGroup)

	excludeGroup, err := cp.GetGroupWithExclusion("172.16.255.16", login.Sid)
	if err != nil {
		logrus.Errorf("Error get group with exclusion: %s", err)
		return
	}
	logrus.Infof("Group with exclusion %v", excludeGroup)

	hostGroup, err := cp.GetHostGroup("172.16.255.16", login.Sid)
	if err != nil {
		logrus.Errorf("Error get host group: %s", err)
		return
	}
	logrus.Infof("Host group %v", hostGroup)

	rangeGroup, err := cp.GetRangeGroup("172.16.255.16", login.Sid)
	if err != nil {
		logrus.Errorf("Error get range group: %s", err)
		return
	}
	logrus.Infof("Range group %v", rangeGroup)

	interfaceGroup, err := cp.GetInterface("172.16.255.16", "cp-81-cluster", login.Sid)
	if err != nil {
		logrus.Errorf("Error get interface group: %s", err)
		return
	}
	logrus.Infof("Interface group %v", interfaceGroup)

}

func (c *Checkpoint) Login(IP, Username, Password string) (*models.LoginStructure, error) {

	logrus.Debugf("Start Login %s", IP)
	defer logrus.Debugf("End Login %s", IP)

	body := map[string]string{
		"user":     Username,
		"password": Password,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/v1.1/login", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("login failed %s", resp.Status)
	}

	result := &models.LoginStructure{}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Login result %v", result)
	return result, nil
}

func (c *Checkpoint) Logout(IP, Token string) error {
	logrus.Debugf("Start Logout %s", IP)
	defer logrus.Debugf("End Logout %s", IP)

	body := map[string]string{}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/logout", IP), bytes.NewReader(byteBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return fmt.Errorf("logout failed %s", resp.Status)
	}
	return nil
}

func (c *Checkpoint) GetServiceGroup(IP, Token string) ([]*models.ServiceGroupStructure, error) {
	logrus.Debugf("Start GetServiceGroup %s", IP)
	defer logrus.Debugf("End GetServiceGroup %s", IP)

	result := make([]*models.ServiceGroupStructure, 0)

	commonGroup, err := c.GetServiceCommonGroup(IP, Token)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Service common group: %v", commonGroup)
	for _, group := range commonGroup.Objects {
		logrus.Debugf("Service group: %v", group)
		curGroup, err := c.GetServiceCurrantGroup(IP, Token, group.Uid)
		if err != nil {
			return nil, err
		}
		result = append(result, curGroup)
	}
	return result, nil
}

func (c *Checkpoint) GetServiceCommonGroup(IP, Token string) (*models.ServiceGroupPluralStructure, error) {
	logrus.Debugf("Start GetServiceGroups %s", IP)
	defer logrus.Debugf("End GetServiceGroups %s", IP)

	body := map[string]int64{
		"offset": 0,
		"limit":  200,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-service-groups", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GetServiceGroups failed %s", resp.Status)
	}

	result := &models.ServiceGroupPluralStructure{}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("GetServiceGroups result %v", result)

	return result, nil
}

func (c *Checkpoint) GetServiceCurrantGroup(IP, Token, GroupID string) (*models.ServiceGroupStructure, error) {
	logrus.Debugf("Start GetServiceGroup %s", IP)
	defer logrus.Debugf("End GetServiceGroup %s", IP)
	body := map[string]string{
		"uid": GroupID,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-service-group", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error GetServiceGroup failed %s", resp.Status)
	}
	result := &models.ServiceGroupStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("GetServiceGroup result %v", result)
	return result, nil
}

func (c *Checkpoint) GetNetworkGroup(IP, Token string) ([]*models.NetworkGroupStructure, error) {
	logrus.Debugf("Start GetNetworkGroup %s", IP)
	defer logrus.Debugf("End GetNetworkGroup %s", IP)

	result := make([]*models.NetworkGroupStructure, 0)
	commonGroup, err := c.GetNetworkCommonGroup(IP, Token)
	if err != nil {
		return nil, errors.New("GetNetworkGroup failed")
	}
	logrus.Debugf("Network common group: %v", commonGroup)
	for _, group := range commonGroup.Objects {
		logrus.Debugf("Network group: %v", group)
		curGroup, err := c.GetNetworkCurrantGroup(IP, Token, group.Uid)
		if err != nil {
			return nil, err
		}
		result = append(result, curGroup)
	}
	return result, nil
}

func (c *Checkpoint) GetNetworkCommonGroup(IP, Token string) (*models.NetworkGroupPluralStructure, error) {
	logrus.Debugf("Start GetNetworkGroups %s", IP)
	defer logrus.Debugf("End GetNetworkGroups %s", IP)
	body := map[string]int64{
		"offset": 0,
		"limit":  200,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("GetNetworkGroups failed")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-groups", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetNetworkGroups failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetNetworkGroups failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, errors.New("GetNetworkGroups failed")
	}
	result := &models.NetworkGroupPluralStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("GetNetworkGroups failed")
	}
	logrus.Debugf("GetNetworkGroups result %v", result)
	return result, nil
}

func (c *Checkpoint) GetNetworkCurrantGroup(IP, Token, GroupID string) (*models.NetworkGroupStructure, error) {
	logrus.Debugf("Start GetNetworkGroup %s", IP)
	defer logrus.Debugf("End GetNetworkGroup %s", IP)

	body := map[string]string{
		"uid": GroupID,
	}
	byteBody, err := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-group", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetNetworkGroup failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetNetworkGroup failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()

	if resp.StatusCode != 200 {
		return nil, errors.New("error status code != 200; GetNetworkGroup failed")
	}
	result := &models.NetworkGroupStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("error unmarshal result; GetNetworkGroup failed")

	}
	logrus.Debugf("GetNetworkGroup result %v", result)
	return result, nil
}

func (c *Checkpoint) GetGroupWithExclusion(IP, Token string) ([]*models.GroupWithExclusionStructure, error) {
	logrus.Debugf("Start GetGroupsWithExclusion %s", IP)
	defer logrus.Debugf("End GetGroupsWithExclusion %s", IP)

	result := make([]*models.GroupWithExclusionStructure, 0)
	commonGroup, err := c.GetGroupWithExclusionCommon(IP, Token)
	if err != nil {
		return nil, errors.New("GetGroupsWithExclusion failed")
	}
	logrus.Debugf("Group with exclusion: %v", commonGroup)
	for _, group := range commonGroup.Objects {
		logrus.Debugf("Group: %v", group)
		curGroup, err := c.GetGroupWithExclusionCurrant(IP, Token, group.Uid)
		if err != nil {
			return nil, err
		}
		result = append(result, curGroup)
	}
	return result, nil
}

func (c *Checkpoint) GetGroupWithExclusionCommon(IP, Token string) (*models.GroupWithExclusionPluralStructure, error) {
	logrus.Debugf("Start GetGroupWithExclusionCommon %s", IP)
	defer logrus.Debugf("End GetGroupWithExclusionCommon %s", IP)

	body := map[string]int64{
		"offset": 0,
		"limit":  200,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCommon failed")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-groups-with-exclusion", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCommon failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCommon failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, errors.New("GetGroupWithExclusionCommon failed")
	}
	result := &models.GroupWithExclusionPluralStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCommon failed")

	}
	logrus.Debugf("GetGroupWithExclusionCommon result %v", result)
	return result, nil
}

func (c *Checkpoint) GetGroupWithExclusionCurrant(IP, Token, GroupID string) (*models.GroupWithExclusionStructure, error) {
	logrus.Debugf("Start GetGroupWithExclusionCurrant %s", IP)
	defer logrus.Debugf("End GetGroupWithExclusionCurrant %s", IP)
	body := map[string]string{
		"uid": GroupID,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCurrant failed")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-group-with-exclusion", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCurrant failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCurrant failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, errors.New("GetGroupWithExclusionCurrant failed")
	}
	result := &models.GroupWithExclusionStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("GetGroupWithExclusionCurrant failed")
	}
	logrus.Debugf("GetGroupWithExclusionCurrant result %v", result)
	return result, nil
}

func (c *Checkpoint) GetHostGroup(IP, Token string) (*models.HostGroupPluralStructure, error) {
	logrus.Debugf("Start GetHost %s", IP)
	defer logrus.Debugf("End GetHost %s", IP)
	body := map[string]int64{
		"offset": 0,
		"limit":  200,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("GetHost failed")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-hosts", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetHost failed")

	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetHost failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, errors.New("GetHost failed")
	}
	result := &models.HostGroupPluralStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("GetHost failed")
	}
	logrus.Debugf("GetHost result %v", result)
	return result, nil
}

func (c *Checkpoint) GetRangeGroup(IP, Token string) ([]*models.RangeGroupStructure, error) {
	logrus.Debugf("Start GetRangeGroup %s", IP)
	defer logrus.Debugf("End GetRangeGroup %s", IP)

	result := make([]*models.RangeGroupStructure, 0)

	commonRange, err := c.GetRangeGroupCommon(IP, Token)
	if err != nil {
		return nil, errors.New("GetRangeGroup failed")
	}
	for _, r := range commonRange.Objects {
		currentRange, err := c.GetRangeGroupCurrant(IP, Token, r.Uid)
		if err != nil {
			return nil, errors.New("GetRangeGroup failed")
		}
		result = append(result, currentRange)
	}
	return result, nil
}

func (c *Checkpoint) GetRangeGroupCommon(IP, Token string) (*models.RangeGroupPluralStructure, error) {
	logrus.Debugf("Start GetRangeGroupCommon %s", IP)
	defer logrus.Debugf("End GetRangeGroupCommon %s", IP)

	body := map[string]int64{
		"limit":  200,
		"offset": 0,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("GetRangeGroupCommon failed")
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-address-ranges", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetRangeGroupCommon failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetRangeGroupCommon failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, errors.New("GetRangeGroupCommon failed")
	}
	result := &models.RangeGroupPluralStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("GetRangeGroupCommon failed")
	}
	logrus.Debugf("GetRangeGroupCommon result %v", result)
	return result, nil
}

func (c *Checkpoint) GetRangeGroupCurrant(IP, Token, GroupID string) (*models.RangeGroupStructure, error) {
	logrus.Debugf("Start GetRangeGroupCurrant %s", IP)
	defer logrus.Debugf("End GetRangeGroupCurrant %s", IP)

	body := map[string]string{
		"uid": GroupID,
	}
	byteBody, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("GetRangeGroupCurrant failed")

	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-address-range", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetRangeGroupCurrant failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetRangeGroupCurrant failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()
	if resp.StatusCode != 200 {
		return nil, errors.New("GetRangeGroupCurrant failed")
	}
	result := &models.RangeGroupStructure{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("GetRangeGroupCurrant failed")
	}
	logrus.Debugf("GetRangeGroupCurrant result %v", result)
	return result, nil
}

func (c *Checkpoint) GetInterface(IP, Name, Token string) ([]*models.InterfaceStructure, error) {
	logrus.Debugf("Start GetInterface %s", IP)
	defer logrus.Debugf("End GetInterface %s", IP)

	result := make([]*models.InterfaceStructure, 0)

	commonInterface, err := c.GetInterfaceCommon(IP, Token)
	if err != nil {
		return nil, errors.New("GetInterface failed")
	}
	for _, i := range commonInterface.Objects {
		if i.Name == Name {
			for _, currentInterface := range i.Interfaces {
				result = append(result, &models.InterfaceStructure{
					InterfaceName:   currentInterface.InterfaceName,
					Ipv4Address:     currentInterface.Ipv4Address,
					Ipv4MaskLength:  currentInterface.Ipv4MaskLength,
					Ipv4NetworkMask: currentInterface.Ipv4NetworkMask,
					DynamicIp:       currentInterface.DynamicIp,
					Topology:        currentInterface.Topology,
				})
			}
		}

	}
	return result, nil
}

func (c *Checkpoint) GetInterfaceCommon(IP, Token string) (*models.InterfacePluralStructure, error) {
	logrus.Debugf("Start GetInterfaceCommon %s", IP)
	defer logrus.Debugf("End GetInterfaceCommon %s", IP)

	body := map[string]any{
		"limit":         200,
		"offset":        0,
		"details-level": "full",
	}
	byteBody, err := json.Marshal(body)

	if err != nil {
		return nil, errors.New("GetInterfaceCommon failed")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/web_api/show-gateways-and-servers", IP), bytes.NewReader(byteBody))
	if err != nil {
		return nil, errors.New("GetInterfaceCommon failed")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-chkp-sid", Token)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("GetInterfaceCommon failed")
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logrus.Errorf("Error close body: %s", err)
		}
	}()

	if resp.StatusCode != 200 {
		return nil, errors.New("GetInterfaceCommon failed")
	}
	result := &models.InterfacePluralStructure{}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.New("GetInterfaceCommon failed")
	}

	logrus.Debugf("GetInterfaceCommon result %v", result)
	return result, nil
}
