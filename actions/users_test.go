package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wedneyyuri/buffalosample/models"
)

func (as *ActionSuite) Test_UsersResource_List() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UsersResource_Show() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UsersResource_Create() {
	// as.Fail("Not Implemented!")
}

func (as *ActionSuite) Test_UsersResource_Update() {
	as.LoadFixture("lots of users")

	firstUser := &models.User{}
	err := as.DB.First(firstUser)
	as.NoError(err)

	userID := firstUser.ID
	as.NotZero(userID)

	// update record
	res := as.JSON("/users/%s", userID).Put(models.User{
		Name: "This is user #2 - Modified",
	})
	as.Equal(res.Response.Code, http.StatusOK)

	// check if it was modified
	res = as.JSON("/users/%s", userID).Get()
	rawcontent, err := ioutil.ReadAll(res.Body)
	as.NoError(err)

	updatedUser := new(models.User)
	err = json.Unmarshal(rawcontent, updatedUser)
	as.NoError(err)

	as.Equal("This is user #2 - Modified", updatedUser.Name)
}

func (as *ActionSuite) Test_UsersResource_Destroy() {
	// as.Fail("Not Implemented!")
}
