package partyrobot

import (
	"fmt"
	"strings"
)

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", strings.Title(strings.ToLower(name)))
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", strings.Title(strings.ToLower(name)), age)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	tableFormatted := fmt.Sprintf("%03d", table)
	distanceFormatted := fmt.Sprintf("%.1f", distance)
	str := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %v. Your table is %s, exactly %s meters from here.\nYou will be sitting next to %s.", strings.Title(strings.ToLower(name)), tableFormatted, strings.ToLower(direction), distanceFormatted, strings.Title(strings.ToLower(neighbor)))
	return str
}