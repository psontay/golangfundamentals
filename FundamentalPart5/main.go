package main

import (
	"FundamentalPart5/note"
	"fmt"
)

func main() {
	//inputFirstName := getUserInput("First Name: ")
	//inputLastName := getUserInput("Last Name: ")
	//inputDOB := getUserInput("Dob: ")
	//appUser, err := user.New(inputFirstName, inputLastName, inputDOB)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//appUser.GetUserDetail()
	//appUser.ClearUsername()
	//appUser.GetUserDetail()

	//appAdmin, err := user.NewAdmin("test@123", "test@123")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//appAdmin.GetUserDetail()
	//appAdmin.ClearUsername()
	//appAdmin.GetUserDetail()
	noteTitle := note.InputDisplay("Input Note Title:")
	noteContent := note.InputDisplay("Input Note Content:")
	appNote, err := note.New(noteTitle, noteContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	appNote.WriteToJson()
	dataFile, err := appNote.ReadFromJson("test1.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dataFile.Title)
	fmt.Println(dataFile.Content)
	fmt.Println(dataFile.CreatedAt)
}
