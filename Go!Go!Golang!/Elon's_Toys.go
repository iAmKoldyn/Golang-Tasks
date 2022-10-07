package elon

import "fmt"
type Track struct {
	distance int
}

func CreateCar(speed, batteryDrain int) *Car {
	return &Car{
		speed, batteryDrain, 100, 0}
}

func CreateTrack(distance int) Track {
	return Track{distance}
}

func (c *Car) Drive() {
	if c.batteryDrain > c.battery {
		return
	}
	c.battery -= c.batteryDrain
	c.distance += c.speed
}

func (c *Car) CanFinish(trackDistance int) bool {
	drives := int(c.battery / c.batteryDrain)
	distance := drives * c.speed
	return distance >= trackDistance
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}