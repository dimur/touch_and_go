// simple "hello-world" code
package touch
import "fmt"

func GetMessage(message string) (res string) {
	return message + " & Go!" 
}

func touch(){
  fmt.Println(GetMessage("Touch"))
}
