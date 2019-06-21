package firstlib

func Checkguess(guess int, whatnum int) (string, bool) {
        if guess == whatnum {
                return "You guessed it", true
        } else if guess < whatnum {
                return "Fizz", false
        } else {
                return "Buzz", false
        }
}

