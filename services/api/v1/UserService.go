package v1Service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	helpers "vipulsirdemo/apiHelpers"
	"vipulsirdemo/models"
	v1repo "vipulsirdemo/repository/api/v1"
	v1req "vipulsirdemo/resources/request/api/v1"
	v1res "vipulsirdemo/resources/response/api/v1"
)

type UserService struct {
	UserRequest v1req.SignUpRequest
	User        models.User
	UserRepo    v1repo.UserRepo
}

func (us *UserService) SignUpUser(ur v1req.SignUpRequest) map[string]interface{} {
	user := us.User
	user.Name = ur.Name
	user.Email = ur.Email
	user.Image = ur.ImageName
	user.Password = ur.Password
	user.Gender = ur.Gender
	user.Mobile = ur.Mobile
	user.Hobby = ur.Hobby

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return helpers.Message(helpers.ResponseError, "Please try again")
	}
	user.Password = string(hashedPassword)

	res, err := us.UserRepo.GetUserByEmail(user.Email)
	if err == nil {
		fmt.Println(res)
		helpers.ImageDelete("public/images/users/" + user.Image)
		return helpers.Message(helpers.ResponseError, "Email is already registered with us.")
	}
	us.UserRepo.UserSignUp(user)
	userData := v1res.UserResponse{
		Name:   user.Name,
		Email:  user.Email,
		Mobile: user.Mobile,
		Hobby:  user.Hobby,
		Image:  "http://localhost:" + helpers.GetPort() + "/public/images/users/" + user.Image,
		//Dob: user.Dob,
		Gender: user.Gender,
	}
	response := helpers.Message(helpers.ResponseSuccess, "You have successfully signed up.")
	response["data"] = userData
	return response
}

func (us *UserService) LoginUser(ur v1req.LoginRequest) map[string]interface{} {
	//fmt.Println("--->", helpers.RandomKeyGenerator(4, "number"))
	user := us.User
	user.Email = ur.Email
	user.Password = ur.Password
	resp, err := us.UserRepo.UserLogin(user)
	if err != nil {
		fmt.Println(resp)
		return helpers.Message(helpers.ResponseError, "Please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return helpers.Message(helpers.ResponseError, "email or password is incorrect.")
	}
	fmt.Println("service id ", resp.ID, resp)
	token := helpers.GenerateAuthToken(resp.Email, resp.ID)

	userData := v1res.LoginResponse{
		Name:   resp.Name,
		Email:  resp.Email,
		Mobile: resp.Mobile,
		Hobby:  resp.Hobby,
		Image:  "http://localhost:" + helpers.GetPort() + "/public/images/users/" + resp.Image,
		//Dob: user.Dob,
		Gender: resp.Gender,
		Token:  token,
	}

	response := helpers.Message(helpers.ResponseSuccess, "Logged in successfully.")
	response["data"] = userData
	return response

	//emailId := os.Getenv("email")
	//emailPassword := os.Getenv("emailPassword")
	//from := emailId
	//password := emailPassword
	//smtpHost := "smtp.gmail.com"
	//smtpPort := "587"
	//
	//to := []string{
	//	user.Email,
	//}
	//
	//message := []byte("1234 is your verification code")
	//auth := smtp.PlainAuth("", from, password, smtpHost)
	//
	//// Sending email.
	//mailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	//if mailErr != nil {
	//	fmt.Println(mailErr)
	//	return helpers.Message(helpers.ResponseError, "Couldn't send email.")
	//}

}

func (us *UserService) EditProfile(ur v1req.EditProfileRequest) map[string]interface{} {
	user := us.User
	user.Name = ur.Name
	user.ID = ur.Id
	user.Image = ur.ImageName
	user.Gender = ur.Gender
	user.Mobile = ur.Mobile
	user.Hobby = ur.Hobby

	res, err := us.UserRepo.GetUserById(user.ID)
	if err != nil {
		fmt.Println(res)
		helpers.ImageDelete("public/images/users/" + user.Image)
		return helpers.Message(helpers.ResponseError, "User not register with "+user.Email)
	}
	us.UserRepo.EditProfile(user)
	helpers.ImageDelete("public/images/users/" + res.Image)
	userData := v1res.UserResponse{
		Name:   user.Name,
		Email:  user.Email,
		Mobile: user.Mobile,
		Hobby:  user.Hobby,
		Image:  "http://localhost:" + helpers.GetPort() + "/public/images/users/" + user.Image,
		//Dob: user.Dob,
		Gender: user.Gender,
	}
	response := helpers.Message(helpers.ResponseSuccess, "Profile update successfully.")
	response["data"] = userData
	return response
}
