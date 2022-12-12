package deathClocker
import (
	"math/rand"
	"time"
	)


func genDeathdate(x int) time.Time {
	rand.Seed(time.Now().UnixNano())
	now:=time.Now()
	yearRand:=rand.Intn(50)
	monthRand:=rand.Intn(12)
	dayRand:=rand.Intn(30)
	death:=now.AddDate(yearRand+x,monthRand,dayRand)
	return death
	
}

func calculate (age int) string {
	deathDate:= genDeathdate(age) 
	z:=deathDate.Format("02-Jan-2006")
	return "Поздравляю, ты умрешь только %v!"+z
}