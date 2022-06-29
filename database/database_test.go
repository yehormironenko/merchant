package database

import (
	"merchant/pkg/models"
	"testing"
)

func TestGetUser(t *testing.T) {

	got, _ := GetUser("user1")
	wantUsername := "user1"
	wantLongname := "Mr. Bean"

	if got.Username != wantUsername && got.Longname != wantLongname {
		t.Errorf("got %q, wanted %q", got.Username, wantUsername)
		t.Errorf("got %q, wanted %q", got.Longname, wantLongname)
	}

}
func TestGetUserFailed(t *testing.T) {

	got, _ := GetUser("us123dsa")
	wantUsername := ""

	if got.Username != wantUsername {
		t.Errorf("got %q, wanted %q", got.Username, wantUsername)
	}
}

func TestSaveUser(t *testing.T) {
	user := models.User{
		Username: "user4",
		Password: "123",
		Longname: "User New",
	}

	err := SaveUser(user)
	if err != nil {
		t.Errorf("Test failed %q", err)
	}

	got, _ := GetUser("user4")

	if got.Username != user.Username && got.Longname != user.Longname {
		t.Errorf("got %q, wanted %q", got.Username, user.Username)
		t.Errorf("got %q, wanted %q", got.Longname, user.Longname)
	}

}
