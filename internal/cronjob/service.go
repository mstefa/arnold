package cronjob

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"arnold/internal/external_login"
)

type CronJob struct {
	externalSessionService external_login.ExternalLooginService
}

type GymClassRes struct {
	ClassId      int `json:"claseId"`
	DisciplineId int `json:"disciplinaId"`
	CoachId      int `json:"coachId"`
	CoachName    string `json:"coachName"`
	Date         string `json:"fecha"`
	Duration     string `json:"duracion"`
	FreeSpots    string `json:"disponibilidad"`
}

type GymClassListRes struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Result  []GymClassRes `json:"result"`
}

func NewCronJob(externalSessionService external_login.ExternalLooginService) CronJob {
	return CronJob{externalSessionService}
}

func (cj CronJob) Init() {
	c := cron.New()
	c.AddFunc("15 19 * * *", cj.CronJobReservation)  //every day at 19:15
	//c.AddFunc("@every 10s", cj.CronJobReservation) //Testing
	go c.Start()

	// Start cron with one scheduled job
	fmt.Printf("Start \n")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Printf("End job \n")
	c.Stop()
}

func (cj CronJob) CronJobReservation() {
	fmt.Printf("%v\n", time.Now())
	userID, err := uuid.NewUUID()
	session, err := cj.externalSessionService.Login(context.Background(), userID.String())
	auth := "Bearer " + session.AccessToken().String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(auth)
	url := "https://classes.megatlon.com.ar/api/service/class/club/category/list"
	method := "POST"

	payload := strings.NewReader(`{
    "categoryId": "74",
    "clubId": "108"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("sec-fetch-mode", " cors")
	req.Header.Add("authorization", auth)
	req.Header.Add("origin", "https://megatlon.com")
	req.Header.Add("refer", "https://megatlon.com/")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	var gymClassListRes GymClassListRes
	json.Unmarshal(body, &gymClassListRes)
	fmt.Printf("List of Classes %+v\n", gymClassListRes)
	fmt.Println("time now UTC")
	fmt.Println(time.Now().UTC())
	fmt.Println("2 day 23 hour ahead")
	timeIn := time.Now().UTC().Add(time.Hour * (47) + time.Minute * 50)  // Two days ahead with 10 mins tolerance
	//timeIn := time.Now().UTC().Add(time.Hour * (46) + time.Minute * 20)  // Testing
	fmt.Println(timeIn)
	fmt.Println("3 day 1 hour ahead")
	timeOut := time.Now().UTC().Add(time.Hour * (48) + time.Minute * 10) // Two days ahead with 10 mins tolerance
	//timeOut := time.Now().UTC().Add(time.Hour * (48) + time.Minute * 10) // Testing
	fmt.Println(timeOut )

	for i := 0; i < len(gymClassListRes.Result); i++ {
		//fmt.Println("for")
		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout , gymClassListRes.Result[i].Date)
		t = t.Add(time.Hour * 3)
		//fmt.Println("Reservation time")
		//fmt.Println(t)
		//fmt.Println("Response time")
		//println(gymClassListRes.Result[i].Date)

		if err != nil {
			fmt.Println(err)
			return
		}
		if t.After(timeIn) && t.Before(timeOut){
			fmt.Println("Book! ")
			fmt.Println(gymClassListRes.Result[i])
			fmt.Println("ClassId")
			fmt.Println(gymClassListRes.Result[i].ClassId)
			fmt.Println(gymClassListRes.Result[i].FreeSpots)
			BookClass(strconv.Itoa(gymClassListRes.Result[i].ClassId), auth)
		}
	}

}


func BookClass(classId, token string){
	url := "https://classes.megatlon.com.ar/api/service/class/book"
	method := "POST"
	reqBody := `{"claseId": "` + classId + `"}`
	fmt.Println("Body")
	fmt.Println(reqBody)


	payload := strings.NewReader(reqBody)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("sec-fetch-mode", " cors")
	req.Header.Add("authorization", token)
	req.Header.Add("origin", "https://megatlon.com")
	req.Header.Add("refer", "https://megatlon.com/")
	req.Header.Add("authority", "users.megatlon.com.ar")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)) // succes: {"code":0,"message":"OK","result":null}

}