package greetings
  
//import "fmt"
import (
        "errors"
        "fmt"
        "math/rand"
        "time"
)


//Hello returns a greeting for the named person.
func Hello(name string) (string, error) {//Any Go function can return multiple values!
        //If no name was given, return an error with a message
        if name == ""{
                return "", errors.New("empty name")//Import the Go standard library errors package so you can use its errors.New function
        }


        //If no name was received, return a value that embeds the name

        //Return a greeting that embeds the name in a message.

	//message := fmt.Sprintf("Hi, %v. Welcome!(v3)", name)
        message := fmt.Sprintf(randomFormat(), name)
        return message, nil
}

//init sets initial values for variables used in the function.
func init(){
        rand.Seed(time.Now().UnixNano())
}

//randomFormat returns one of a set of greeting messages.The returned message is selected at random.
func randomFormat() string {
        //A slice of message formats.
        formats := []string{
                "Hi,%v. Welcome!",
                "Great to see you, %v!",
                "Hail, %v Well met!",
        }

        //Return a randomly selected message format by specifying a random index for the slice of formats
        return formats[rand.Intn(len(formats))]
}

/**
Declare a greetings package to collect related functions.
Implement a Hello function to return the greeting.
**

/**
In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:

var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
**/
