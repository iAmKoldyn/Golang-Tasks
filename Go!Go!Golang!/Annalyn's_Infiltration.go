package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	return knightIsAwake == false
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return archerIsAwake == false && prisonerIsAwake == true
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (prisonerIsAwake && !archerIsAwake && !knightIsAwake && !petDogIsPresent) || (petDogIsPresent && !archerIsAwake)
}