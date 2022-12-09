package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
    "math/rand"
)

func clear() {
    fmt.Print("\033[H\033[2J")
}

func print_map(arr[9]string) {
    colorRed := "\033[31m"
    colorBlue := "\033[34m"
    colorReset := "\033[0m"
    color_to_apply := [9]string{"","","","","","","","",""}

    for i := 0 ; i < len(arr) ; i++ {
        if arr[i] == "X" {
            color_to_apply[i] = colorRed

        } else if arr[i] == "O" {
            color_to_apply[i] = colorBlue

        } else {
            color_to_apply[i] = colorReset
        }

    }
    fmt.Println(string(color_to_apply[0]),arr[0],string(color_to_apply[1]),arr[1],string(color_to_apply[2]),arr[2])
    fmt.Println(string(color_to_apply[3]),arr[3],string(color_to_apply[4]),arr[4],string(color_to_apply[5]),arr[5])
    fmt.Println(string(color_to_apply[6]),arr[6],string(color_to_apply[7]),arr[7],string(color_to_apply[8]),arr[8])
    fmt.Println(string(colorReset))
}

func ia_turn(arr[9]string,left_space[]int) int{
    num := rand.Intn(len(left_space))
    _ = num
    return left_space[num]

}

func check_win_or_loose(arr[9]string,player1 string,player2 string,Sign string) {
    draw := true

    for i := 0 ; i < len(arr) ; i++ {
        if arr[i] == "." {
            draw = false
        }
    }
    if draw == true {
        fmt.Println("It's draw time")
        os.Exit(0)
    }

    if (arr[0] == Sign && arr[1] == Sign && arr[2] == Sign) || (arr[3] == Sign && arr[4] == Sign && arr[5] == Sign) || (arr[6] == Sign && arr[7] == Sign && arr[8] == Sign) {
        fmt.Println("Well done",player1, "you won")
        fmt.Println(player2," lost")
        os.Exit(0)
        
    } else if (arr[0] == Sign && arr[3] == Sign && arr[6] == Sign) || (arr[1] == Sign && arr[4] == Sign && arr[7] == Sign) || (arr[2] == Sign && arr[5] == Sign && arr[8] == Sign) {
        fmt.Println("Well done",player1, "you won")
        fmt.Println(player2,"lost")
        os.Exit(0)
    

    } else if (arr[0] == Sign && arr[4] == Sign && arr[8] == Sign) || (arr[2] == Sign && arr[4] == Sign && arr[6] == Sign) {
        fmt.Println("Well done",player1, "you won")
        fmt.Println(player2,"lost")
        os.Exit(0)
    
    }
}
func vs_ia() {
    scanner := bufio.NewScanner(os.Stdin)
	arr := [9]string{".",".",".",".",".",".",".",".","."}
    left_space :=[]int{}
    empty :=[]int{}

    print_map(arr)
    fmt.Println("Choose an empty spot from 1 to 9: ")
    for scanner.Scan() {
		fmt.Println("Choose an empty spot from 1 to 9: ")
        x , err := strconv.Atoi(scanner.Text())
        if x < 10 && x > 0 && arr[x - 1] == "."{
            for i := 0 ; i < len(arr) ; i++ {
                if arr[i] == "." {
                    left_space = append(left_space,i)
                }
            }
            arr[x - 1] = "X"
            arr[ia_turn(arr,left_space)] = "O"
            
            if err != nil {
                panic(err)
            }    
            print_map(arr)
    
	    } else {
            fmt.Println("Put a number beetween 1 to 9 or an empty space")
            print_map(arr)
        }
        left_space = empty
        check_win_or_loose(arr,"Steeven","Ia","X")
        check_win_or_loose(arr,"Ia","Steeven","O")
    }
}

func vs_human() {
    n_1 := bufio.NewScanner(os.Stdin)
    n_2 := bufio.NewScanner(os.Stdin)
    
    player1 := bufio.NewScanner(os.Stdin)
	player2 := bufio.NewScanner(os.Stdin)
    arr := [9]string{".",".",".",".",".",".",".",".","."}
    name1 , name2 :=  "" , ""
    fmt.Println("Write your name Player 1: ")
    for n_1.Scan() {
        name1 = n_1.Text()
        break
    }
    fmt.Print("Write your name Player 2: ")
    for n_2.Scan() {
        name2 = n_2.Text()
        break
    }
    once := true
    print_map(arr)
    fmt.Println("Player 1 choose an empty spot from 1 to 9: ")
    for true {
        for player1.Scan() {
            if once == false {
                fmt.Println("Player 1 choose an empty spot from 1 to 9: ")
            }
            if once == true {
                once = false
            }
            x , err := strconv.Atoi(player1.Text())
            if x < 10 && x > 0 && arr[x - 1] == "."{
                arr[x - 1] = "X"
                
                if err != nil {
                    panic(err)
                }    
                print_map(arr)
        
            } else {
                fmt.Println("Put a number beetween 1 to 9 or an empty space")
                print_map(arr)
            }
            check_win_or_loose(arr,name1,name2,"X")
            check_win_or_loose(arr,name2,name1,"O")
            break
        }
        fmt.Println("Player 2 choose an empty spot from 1 to 9: ")
        for player2.Scan() {
            x , err := strconv.Atoi(player2.Text())
            if x < 10 && x > 0 && arr[x - 1] == "."{
                arr[x - 1] = "O"
                
                if err != nil {
                    panic(err)
                }    
                print_map(arr)
        
            } else {
                fmt.Println("Put a number beetween 1 to 9 or an empty space")
                print_map(arr)
            }
            check_win_or_loose(arr,name1,name2,"X")
            check_win_or_loose(arr,name2,name1,"O")
            break
        }
        fmt.Println("Player 1 choose an empty spot from 1 to 9: ")
    }
}

func maingame() {
    fmt.Println("TICTACTOE")
    fmt.Println("")
    fmt.Println("1 VS IA\n2 VS HUMAN\n3 EXIT")
    fmt.Print("Choose: ")
    choice := bufio.NewScanner(os.Stdin)
    
    for choice.Scan() {
        if choice.Text() == "1" {
            clear()
            vs_ia()

        } else if choice.Text() == "2" {
            clear()
            vs_human()

        } else if strings.ToLower(choice.Text()) == "EXIT" {
            clear()
            fmt.Println("Bye Bye")
            os.Exit(0)
        }else {
            fmt.Println("Choose beetween 1 or 2 or write EXIT to quit")
        }
    }
}

func main() {
    clear()
    maingame()
}
