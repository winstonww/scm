package core

import "math/rand"

func Initialize(numControllers, numSwitches int,
	cMin, cTotal, cUtil, nMin, nTotal, nUtil float64) ([]Controller, []SwitchMigration) {

	controllers := InitializeControllers(numControllers, cMin, cTotal, nMin, nTotal)
	migrations := InitializeMigrations(numSwitches, cTotal, cUtil, nTotal, nUtil, controllers)

	return controllers, migrations
}

func InitializeControllers(numControllers int,
	cMin, cTotal, nMin, nTotal float64) []Controller {

	// Create an array of switch migrations and controllers
	controllers := []Controller{}
	cCap := NormalizedRandomArray(cMin, cTotal, numControllers)
	nCap := NormalizedRandomArray(nMin, nTotal, numControllers)
	// Initialize controllers
	for i := 0; i < numControllers; i++ {
		capacity := Load{cCap[i], nCap[i]}
		c := Controller{i, capacity, Load{}}
		controllers = append(controllers, c)
	}
	return controllers
}

func InitializeMigrations(numSwitches int,
	cTotal, cUtil, nTotal, nUtil float64, controllers []Controller) []SwitchMigration {
	// Initalize migrations edges
	migrations := []SwitchMigration{}
	cUsed := NormalizedRandomArray(0, cTotal*cUtil, numSwitches)
	nUsed := NormalizedRandomArray(0, nTotal*nUtil, numSwitches)
	for i := 0; i < numSwitches; i++ {
		// Random pick a controller and assign switch to it
		c := &controllers[rand.Int()%len(controllers)]
		usage := Load{cUsed[i], nUsed[i]}
		sm := SwitchMigration{c, c, usage}
		c.Used = (c.Used).Add(usage)
		migrations = append(migrations, sm)
	}
	return migrations
}
