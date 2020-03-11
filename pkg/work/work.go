package work

import (
	"encoding/json"
	"fmt"
	"log"

	"kusnandartoni/starter/pkg/logging"
	"kusnandartoni/starter/redisdb"
	"kusnandartoni/starter/services/svcmail"
)

// DoWork :
func DoWork(data string, id int) {
	// defer wg.Done()
	var (
		emailData svcmail.EmailData
		logger    = logging.Logger{UUID: "MAILER"}
	)
	to := ""
	err := json.Unmarshal([]byte(data), &emailData)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}

	logger.Debug(fmt.Sprintf("worker [%d] is sending [%s] email \n", id, emailData.EmailType))
	// email process here
	if emailData.EmailType == "forgot" {
		var forgot svcmail.Forgot
		err = json.Unmarshal([]byte(emailData.Data), &forgot)
		if err != nil {
			log.Print(err)
		}
		to = forgot.Email
		err = forgot.Send()
		if err != nil {
			log.Print(err)
		}
	}

	// if emailData.EmailType == "register" {
	// 	var register svcmail.Register
	// 	err = json.Unmarshal([]byte(emailData.Data), &register)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = register.Email
	// 	err = register.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	// if emailData.EmailType == "approve" {
	// 	var approve svcmail.Approve
	// 	err = json.Unmarshal([]byte(emailData.Data), &approve)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = approve.Email
	// 	err = approve.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	// if emailData.EmailType == "reject" {
	// 	var reject svcmail.Reject
	// 	err = json.Unmarshal([]byte(emailData.Data), &reject)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = reject.Email
	// 	err = reject.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	// if emailData.EmailType == "adminInvite" {
	// 	var adminInvite svcmail.AdminInvite
	// 	err = json.Unmarshal([]byte(emailData.Data), &adminInvite)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = adminInvite.Email
	// 	err = adminInvite.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	// if emailData.EmailType == "contactUs" {
	// 	var contactUs svcmail.ContactUs
	// 	err = json.Unmarshal([]byte(emailData.Data), &contactUs)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = contactUs.Email
	// 	err = contactUs.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	// if emailData.EmailType == "setPassword" {
	// 	var setPassword svcmail.SetPassword
	// 	err = json.Unmarshal([]byte(emailData.Data), &setPassword)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = setPassword.Email
	// 	err = setPassword.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	// if emailData.EmailType == "approveMemberSupergroup" {
	// 	var approveMemberSupergroup svcmail.ApproveMemberSupergroup
	// 	err = json.Unmarshal([]byte(emailData.Data), &approveMemberSupergroup)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = approveMemberSupergroup.Email
	// 	err = approveMemberSupergroup.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	// if emailData.EmailType == "rejectMemberSupergroup" {
	// 	var rejectMemberSupergroup svcmail.RejectMemberSupergroup
	// 	err = json.Unmarshal([]byte(emailData.Data), &rejectMemberSupergroup)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	to = rejectMemberSupergroup.Email
	// 	err = rejectMemberSupergroup.Send()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }
	//
	logger.Debug(fmt.Sprintf("worker [%d] sending to [%s] done... \n", id, to))
	if err != nil {
		redisdb.AddList("starter_email", data)
	}
}
