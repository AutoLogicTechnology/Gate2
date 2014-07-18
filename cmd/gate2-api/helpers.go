
package main 

func haveuser (userid string) (bool) {
    for _, user := range gates {
        if user.UserID == userid {
            return true
        }
    }

    return false 
}