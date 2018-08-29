package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
	"flag"
)


const SchoolID = "5c360334-4cbc-4fd9-934e-61f7eec026dc"
const RequestCoachListURL = "http://mp.zhjp.sz-its.cn/ZHJP.WebApp.ServiceProxy.svc/TransformData"
const AppointmentCoachURL = "http://mp.zhjp.sz-its.cn/ZHJP.WebApp.ServiceProxy.svc/SetData"


var qCoachName string
var qTrainingLocation string
var UserID string
var wxopenid string
var qScheduleDate string
var qTrainingPart string
var qTrainingCategory string

type CoachScheduleList struct {
	Rownumber int `json:"rownumber"`
	Scheduleguid string `json:"ScheduleGuid"`
	Coachname string `json:"CoachName"`
	Coachcode string `json:"CoachCode"`
	Scheduledate string `json:"ScheduleDate"`
	Scheduleperiod string `json:"SchedulePeriod"`
	Trainingpart string `json:"TrainingPart"`
	Trainingcategory string `json:"TrainingCategory"`
	Traininglocation string `json:"TrainingLocation"`

	Coachsex string `json:"CoachSex"`
	Trainedtotal int `json:"TrainedTotal"`
	Behavioraverage int `json:"BehaviorAverage"`
	Qualityaverage int `json:"QualityAverage"`

	Trainingdatedisplay string `json:"TrainingDateDisplay"`
	Trainingcategorydisplay string `json:"TrainingCategoryDisplay"`
	Trainedtotaldisplay string `json:"TrainedTotalDisplay"`
	Behavioraveragedisplay string `json:"BehaviorAverageDisplay"`
	Qualityaveragedisplay string `json:"QualityAverageDisplay"`
}

type ResponseStruct struct {
	Islastpage bool `json:"IsLastPage"`
	Coachschedulelist []CoachScheduleList `json:"CoachScheduleList"`
}


type CourseInfoStruct struct {
	Trainingpart string `json:"TrainingPart"`
	Trainingcategory string `json:"TrainingCategory"`
	Coachguid string `json:"CoachGuid"`
	Coachname string `json:"CoachName"`
	Coachcode string `json:"CoachCode"`
	Trainingdate string `json:"TrainingDate"`
	Trainingperiod string `json:"TrainingPeriod"`
	Traininglocation string `json:"TrainingLocation"`
	Traininghours int `json:"TrainingHours"`
	Receivableamount float64 `json:"ReceivableAmount"`
}

func RequestCoachList() string{
	reqdata := `{"UserID":"` + UserID + `","SchoolGuid":"` + SchoolID + `","DataType":"CoachSchedule_QueryForReservation","JsonTransform":"{\"StudentID\":\"` +
		UserID + `\",\"SchoolID\":\"` + SchoolID + `\",\"CoachName\":\"` +
		qCoachName + `\",\"ScheduleDate\":\"` + qScheduleDate + `\",\"StartTime\":\"00 : 00\",\"TrainingLocation\":\"` +
		qTrainingLocation + `\",\"TrainingPart\":\"` +
		qTrainingPart + `\",\"TrainingCategory\":\"` +
		qTrainingCategory + `\",\"CoachGender\":\"\",\"PageIndex\":0,\"PageSize\":20}"}`
	client := &http.Client{}
	req, err := http.NewRequest("POST", RequestCoachListURL , strings.NewReader(reqdata))

	if err != nil {
		// handle error
		return "nil"
	}
	req.Header.Set("Host", "mp.zhjp.sz-its.cn")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept-Language", "zh-cn")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "http://mp.zhjp.sz-its.cn")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E302 MicroMessenger/6.6.7 NetType/WIFI Language/zh_CN")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "key=" + UserID + "; wxopenid=" + wxopenid)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "nil"
	}
	//fmt.Println(string(body))
	return string(body)
}

func SelectCoach(Scheduleid string) string{
	client := &http.Client{}
	//
	req, err := http.NewRequest("POST", RequestCoachListURL, strings.NewReader(`{"UserID":"` + UserID + `","SchoolGuid":"` + SchoolID +
		`","DataType":"CoachSchedule_GetCoachScheduleForReservation","JsonTransform":"{\"StudentID\":\"` + UserID + `\",\"SchoolID\":\"` +
		SchoolID + `\",\"ScheduleID\":\"` + Scheduleid + `\",\"TrainingPart\":\"` + qTrainingPart + `\",\"TrainingCategory\":\"` + qTrainingCategory + `\"}"}`))
	if err != nil {
		// handle error
		return "nil"
	}
	req.Header.Set("Host", "mp.zhjp.sz-its.cn")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept-Language", "zh-cn")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "http://mp.zhjp.sz-its.cn")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E302 MicroMessenger/6.6.7 NetType/WIFI Language/zh_CN")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "key=" + UserID + "; wxopenid=" + wxopenid)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "nil"
	}
	//fmt.Println(string(body))
	return string(body)
}

func ConfirmCoach(Scheduleid string, Coachguid string, Receivableamount int) string{
	client := &http.Client{}
	//
	req, err := http.NewRequest("POST", AppointmentCoachURL, strings.NewReader(`{"UserID":"` + UserID +
		`","DataGuid":"","DataType":"LrnAppoint_MakeAReservation","JsonData":"{\"StudentID\":\"` + UserID +
		`\",\"SchoolID\":\"` + SchoolID + `\",\"ScheduleID\":\"` + Scheduleid + `\",\"TrainingPart\":\"` +
		qTrainingPart + `\",\"TrainingCategory\":\"` + qTrainingCategory + `\",\"CoachGuid\":\"` +
		Coachguid + `\",\"ReceivableAmount\":\"` + strconv.Itoa(Receivableamount) + `\",\"DiscountPolicyGuid\":\"\",\"DiscountPolicyName\":\"\"}"}`))
	if err != nil {
		// handle error
		return "nil"
	}
	req.Header.Set("Host", "mp.zhjp.sz-its.cn")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Accept-Language", "zh-cn")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "http://mp.zhjp.sz-its.cn")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E302 MicroMessenger/6.6.7 NetType/WIFI Language/zh_CN")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "key=" + UserID + "; wxopenid=" + wxopenid)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "nil"
	}
	fmt.Println(string(body))
	return string(body)
}



func main() {
	
	flag.StringVar(&UserID, "uid", "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", "UserID")
	flag.StringVar(&wxopenid, "wid", "xx-x_xxxxxxxxxxxxxxxxxxrxxxx", "wxopenid")
	flag.StringVar(&qScheduleDate, "sd", "xxxx-xx-xx", "ScheduleDate")
	flag.StringVar(&qTrainingPart, "tp", "科目三", "TrainingPart")
	flag.StringVar(&qTrainingCategory, "tc", "道路操作", "TrainingCategory")
	flag.StringVar(&qCoachName, "cn", "xx", "TrainingPart")
	flag.StringVar(&qTrainingLocation, "tl", "xxx", "TrainingCategory")
	flag.Parse()
	
	fmt.Println("UserID = " + UserID + "wxopenid = " + wxopenid + "qScheduleDate = " + qScheduleDate + "qTrainingPart = " + qTrainingPart + "qTrainingCategory = " + qTrainingCategory)
	ResponseJson := RequestCoachList();
	ResponseJson = strings.Replace(ResponseJson, "\\", "", -1)
	ResponseJson = string(ResponseJson[1:len(ResponseJson) - 1])

	var res1 ResponseStruct
	if err := json.Unmarshal([]byte(ResponseJson), &res1) ; err != nil {
		panic( err)
	} else {
		//fmt.Println(res1.Coachschedulelist[0].Scheduleguid)
	}

	var count int
	count = 0
	for {
		for i := 0; i < len(res1.Coachschedulelist); i++ {

			if len(res1.Coachschedulelist) > 1 && "06:00~08:00" == res1.Coachschedulelist[i].Scheduleperiod {
				fmt.Println("skip 06:00~08:00")
				continue
			}
			if count == 2 {
				fmt.Println("ordered two success")
				return
			}
			ResponseJson = SelectCoach(res1.Coachschedulelist[i].Scheduleguid)
			ResponseJson = strings.Replace(ResponseJson, "\\", "", -1)
			ResponseJson = string(ResponseJson[1:len(ResponseJson) - 1])
			var res2 CourseInfoStruct
			if err := json.Unmarshal([]byte(ResponseJson), &res2) ; err != nil {
				panic( err)
			} else {
				//fmt.Println(res2)
			}

			ret := ConfirmCoach(res1.Coachschedulelist[i].Scheduleguid, res2.Coachguid, int(res2.Receivableamount))
			if "12305" == ret {
				//fmt.Println("12305")
			}

			if "12304" == ret {
				//fmt.Println("12304")
				fmt.Println(res2)
			}

			if "0" == ret {
				count++;
				fmt.Println(res2)
			}
		}
	}

}
