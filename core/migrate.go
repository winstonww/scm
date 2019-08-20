package core

// import "fmt"

func GreedyMigrate(migrations []SwitchMigration, controllers []Controller) {
	// This function sets the destination of all migrations
	// Input: Array of migrations with destination equal to the source
	//        Array of controllers with pre-specified capacities and usage
	// Return: None

	// Sort migrations by load first
	// Using L2 norm for stability and increase sensitivity to heavy CPU or Bandwidth load
	l2 := func(s1, s2 *SwitchMigration) bool {
		return s1.Utilization.L2Norm() < s2.Utilization.L2Norm()
	}

	OrderedBy(l2).Sort(migrations) // sort migrations based on that

	for i, _ := range migrations {
		for j, _ := range controllers {
			if (&migrations[i]).Cost(&controllers[j]) < 0 {
				Debugln(migrations[i])
				Debugf("%p \n", &controllers[j])
				(&migrations[i]).MigrateTo(&controllers[j])
				Debugln(&migrations[i])
			}
		}
	}

	// after migration
	for _, c := range migrations {
		Debugln(c)
	}
}
