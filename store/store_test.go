package store

import (
	"testing"
)

var testService string = "youtube"
var testPassword string = "36@worlds!"

func setup() *Store {
	initialData := make(map[string]string)
	passwordStore := NewPasswordStore(initialData)
	return passwordStore
}

func setupWithValue() *Store {
	initialData := make(map[string]string)
	passwordStore := NewPasswordStore(initialData)
	passwordStore.data[testService] = testPassword
	return passwordStore
}

func TestNewPasswordStore(t *testing.T) {
	var passwordStore interface{}
	passwordStore = setup()

	if _, ok := passwordStore.(*Store); !ok {
		t.Errorf("TestNewPasswordStore not returning pointer to Store")
	}
}

func TestWriteOk(t *testing.T) {
	passwordStore := setup()
	passwordStore.Write(testService, testPassword)
	_, ok := passwordStore.data[testService]

	if !ok {
		t.Errorf("TestWrite expected key for value to be ok, was not ok")
	}
}

func TestWriteValue(t *testing.T) {
	passwordStore := setup()
	passwordStore.Write(testService, testPassword)
	pwd, _ := passwordStore.data[testService]

	if pwd != testPassword {
		t.Errorf("TestWrite expected password to equal '%s', instead equals '%s'", testPassword, pwd)
	}
}

func TestReadNoError(t *testing.T) {
	passwordStore := setupWithValue()
	_, err := passwordStore.Read(testService)

	if err != nil {
		t.Errorf("TestRead returned error, should have been nil")
	}
}

func TestReadError(t *testing.T) {
	passwordStore := setupWithValue()
	_, err := passwordStore.Read("youtub")

	if err == nil {
		t.Errorf("TestRead should have returned error for non existing service")
	}
}

func TestReadValue(t *testing.T) {
	passwordStore := setupWithValue()
	pwd, _ := passwordStore.Read(testService)

	if pwd != testPassword {
		t.Errorf("TestReadValue expected password to be '%s', instead was '%s'", testPassword, pwd)
	}
}

func TestDelete(t *testing.T) {
	passwordStore := setupWithValue()
	passwordStore.Delete(testService)

	if _, ok := passwordStore.data[testService]; ok {
		t.Errorf("TestDelete key for '%s' was present, should have been empty", testService)
	}
}

func TestServiceExistsTrue(t *testing.T) {
	passwordStore := setupWithValue()
	if !passwordStore.ServiceExists(testService) {
		t.Errorf("TestServiceExists ServiceExists should return true for '%s', returned false", testService)
	}
}

func TestServiceExistsFalse(t *testing.T) {
	passwordStore := setup()
	if passwordStore.ServiceExists("twitter") {
		t.Errorf("TestServiceExists ServiceExists should return false for 'twitter', returned true")
	}
}

func TestGetStoreReturnsPasswordData(t *testing.T) {
	var passwordData interface{}
	passwordStore := setup()
	passwordData = passwordStore.GetStore()
	if _, ok := passwordData.(PasswordData); !ok {
		t.Errorf("TestGetStoreReturnsStore expected output to be of type PasswordData, it wasn't")
	}
}

func TestGetStoreReturnsDataWithValuePresent(t *testing.T) {
	passwordStore := setupWithValue()
	passwordData := passwordStore.GetStore()

	if _, ok := passwordData[testService]; !ok {
		t.Errorf("TestGetStoreReturnsDataWithValue expected '%s' to be present in map, it was not", testService)
	}
}

func TestGetStoreReturnsDataWithCorrectValue(t *testing.T) {
	passwordStore := setupWithValue()
	passwordData := passwordStore.GetStore()

	if pwd, _ := passwordData[testService]; pwd != testPassword {
		t.Errorf("TestGetStoreReturnsDataWithCorrectValue expected password to be'%s', instead it was '%s'", testPassword, pwd)
	}
}
