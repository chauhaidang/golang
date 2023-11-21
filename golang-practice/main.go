package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	Generic()
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error = nil
	if denominator == 0 {
		err = errors.New("can not divide to zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}

func funcCall() {
	numerator := 11
	denominator := 2
	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if numerator == 0 {
		fmt.Printf("result is : %v because numerator equal to 0", numerator)
	} else {
		fmt.Printf("result is: %v, remainder is: %v", result, remainder)
	}

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case numerator == 0:
		fmt.Printf("result is : %v because numerator equal to 0", numerator)
	default:
		fmt.Printf("result is: %v, remainder is: %v", result, remainder)
	}

}

func Arrays() {
	// Arrays
	// Fixed length
	// Same type
	// Continuous

	var intArr1 [4]int32
	intArr1[1] = 1
	intArr1[2] = 2
	intArr1[3] = 3
	fmt.Println(intArr1[0])
	fmt.Println(intArr1[1:3])

	intArr2 := [4]int32{1, 2, 3, 4}
	fmt.Println(intArr2)
	intArr3 := [...]int32{1, 2, 3, 54}
	fmt.Println(intArr3)

}

func Slice() {
	// Slice
	intSlice := []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	intSlice2 := []int32{7, 8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	intSlice3 := make([]int32, 3, 8)
	intSlice4 := append(intSlice3, 1)
	fmt.Println(intSlice4)

}

func Map() {
	// Map
	myMap := map[string]int{"David": 28, "Thanh": 24}
	age, ok := myMap["David"]

	if ok {
		fmt.Printf("the age is %v \n", age)
	} else {
		fmt.Println("Invalid Name")
	}

	delete(myMap, "Thanh")
	fmt.Println(myMap)

	myMap["Thanh"] = 25
	fmt.Println(myMap)

	for k, v := range myMap {
		fmt.Print(k)
		fmt.Print("|")
		fmt.Print(v)
		fmt.Print("\n")
	}
}

func StringAdvance() {
	myString := "résumé"
	indexed := myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	myString2 := []rune("résumé")
	indexed2 := myString2[1]
	fmt.Printf("%v, %T\n", indexed2, indexed2)
	for i, v := range myString2 {
		fmt.Println(i, v)
	}
	fmt.Printf("my lengh is: %v \n", len(myString2))

	myRune := 'a'
	fmt.Printf("My Rune = %v\n", myRune)

	myString3 := []string{"d", "a", "v", "i", "d"}
	catStr := "hello "
	for i := range myString3 {
		catStr += myString3[i]
	}
	fmt.Println(catStr)

	// Better built-in library "strings"
	var strBuilder strings.Builder
	catStr = ""
	for i := range myString3 {
		strBuilder.WriteString(myString3[i])
	}
	catStr = strBuilder.String()
	fmt.Println(catStr)
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
	owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

// Func take 1st parameter as interface
func canMakeIt(e engine, miles uint8) {
	fmt.Print("\n")
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

type engine interface {
	milesLeft() uint8
}

// Struct & Interface
func Struct() {
	myEngine := gasEngine{25, 15, owner{name: "David"}}
	// not re-useable
	myEngine2 := struct {
		mpg     uint8
		gallons uint8
	}{25, 25}

	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Printf("Result of extension function is: %v", myEngine.milesLeft())

	myEEngine := electricEngine{kwh: 4, mpkwh: 10}
	canMakeIt(myEEngine, 10)
	canMakeIt(myEngine, 20)
}

// Pointer
func Pointer() {
	myArr := make([]string, 1)
	myArrCopied := myArr

	myArr[0] = "David"
	fmt.Println(myArr)
	fmt.Println(myArrCopied)

	myEngine := gasEngine{25, 15, owner{name: "David"}}
	// Use ampersand to reference to the address of myEngine instead of copying its value
	myEngineCopied := &myEngine
	myEngineCopied.owner.name = "David Chau"
	fmt.Println(myEngine)
	fmt.Println(myEngineCopied)
}

// Goroutines
var dbData = []string{"d1", "d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "d10"}
var results = []string{}
var wg = sync.WaitGroup{}

// var m = sync.Mutex{}
var m = sync.RWMutex{}

func ioCall(i int) {
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func Goroutines() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go ioCall(i)
	}
	wg.Wait()
	fmt.Println("Total execution time is: ", time.Since(t0))
}

func save(result string) {
	// Full lock, no go routine can reach resource
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	// Only read when no full lock process happen.
	// This allows multiple goroutines access and read resource in parallel
	m.RLock()
	fmt.Println("The result is: ", results)
	m.RUnlock()
}

func Channel() {
	c := make(chan int)
	go processC(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

// ChannelBuffer Create a buffer help channel store multiple value in its storage
// rather than single slot
func ChannelBuffer() {
	c := make(chan int, 5)
	go processC(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func processC(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("exit process")
}

func ChannelReal() {
	chickenChan := make(chan string)
	tofuChan := make(chan string)
	website := []string{"wallmart", "tosco", "wholefood"}
	for i := range website {
		go checkChickenPrices(website[i], chickenChan)
		go checkTofuPrices(website[i], tofuChan)
	}
	sendMessage(chickenChan, tofuChan)
}

const MaxChickenPrice float32 = 5
const MaxTofuPrice float32 = 5

func checkTofuPrices(site string, tofuChan chan string) {
	for {
		time.Sleep(time.Second * 1)
		tofuPrice := rand.Float32() * 20
		if tofuPrice <= MaxTofuPrice {
			// Pass the site to channel which only have 1 slot
			tofuChan <- site
			break
		}
	}
}

func checkChickenPrices(site string, chickenChan chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		if chickenPrice <= MaxChickenPrice {
			// Pass the site to channel which only have 1 slot
			chickenChan <- site
			break
		}
	}
}

func sendMessage(chickChan chan string, tofuChan chan string) {
	// This like and if else for channel,
	// whenever one of these channels receive a message, execute the code of its "case"
	select {
	case <-chickChan:
		fmt.Printf("\n found a deal on CHICKEN at store: %s", <-chickChan)
		close(chickChan)
	case <-tofuChan:
		fmt.Printf("\n found a deal on TOFU at store: %s", <-tofuChan)
		close(tofuChan)
	}
}

func Generic() {
	intSlice := []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))
	floatSlice := []float32{2, 2, 4}
	fmt.Println(sumSlice(floatSlice))

	contacts := loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)
	purchases := loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)

	electricCar := &car[electricEngine]{
		carMake:  "Mitsubishi",
		carModel: "Attrage",
		engine: electricEngine{
			mpkwh: 47,
			kwh:   57,
			owner: owner{
				name: "David",
			},
		},
	}

	fmt.Println(electricCar)
	electricCar.carModel = "Attrage 1.2 Premium"
	electricCar.engine.owner.name = "Thanh Truong"
	fmt.Println(electricCar)
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

type car[T electricEngine | gasEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	loaded := []T{}
	byteData, _ := os.ReadFile(filePath)
	json.Unmarshal(byteData, &loaded)
	return loaded
}
