package main  
  
import (  
	"bufio"  
	"fmt"  
	"os"  
	"strings"  
	"time"  
)  
  
func isLeapYear(year int) bool {  
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)  
}  
  
func getDaysInMonth(year, month int) int {  
	
	days := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC).Day()  
	return days  
}  
  
func printMonthCalendar(year, month int) {  
	
	t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)  
	weekday := int(t.Weekday()) 
  
	days := getDaysInMonth(year, month)  
  
	fmt.Printf("\n%dyear%02dmonth\n", year, month)  
	fmt.Println("Sun Mon Tue Wed Thu Fri Sat")  
  
	for i := 0; i < weekday; i++ {  
		fmt.Print("   ")  
	}  
  
	for day := 1; day <= days; day++ {  
		fmt.Printf("%3d ", day)  
		if (day+weekday-1)%7 == 0 {  
			fmt.Println()  
		}  
	}  
  
	if (days+weekday-1)%7 != 0 {  
		fmt.Println()  
	}  
}  
  
func printYearCalendar(year int) {  
	for month := 1; month <= 12; month++ {  
		printMonthCalendar(year, month)  
	}  
}  
  
func main() {  
	reader := bufio.NewReader(os.Stdin)  
	fmt.Print("Enter year: ")  
	text, _ := reader.ReadString('\n') 
	year, _ := strconv.Atoi(strings.TrimSpace(text)) 
  
	printYearCalendar(year)  
}  
