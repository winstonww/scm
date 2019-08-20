package core

import (
	"reflect"
)

type SwitchMigration struct {
	Source      *Controller
	Destination *Controller
	Utilization Load
}

func (sm *SwitchMigration) MigrateTo(c *Controller) {
	if sm.Destination != nil {
		sm.Destination.Used = sm.Destination.Used.Subtract(sm.Utilization)
	}
	sm.Destination = c
	sm.Destination.Used = sm.Destination.Used.Add(sm.Utilization)
}
func (sm *SwitchMigration) Cost(c *Controller) float64 {
	// This function set destination of migration to Controller c
	cost := 0.0
	if sm.Destination != nil {
		// subtract previous cost
		originalCost := (sm.Destination).Capacity.Subtract(sm.Destination.Used).Add(sm.Utilization)
		cost -= Cost(originalCost, sm.Utilization)
	}
	cost += Cost(c.Capacity.Subtract(c.Used), sm.Utilization)
	return cost
}

func Cost(remaining Load, usage Load) float64 {
	// cost of assign usage load to the remaining load
	cost, rr, ss := 0.0, reflect.ValueOf(remaining), reflect.ValueOf(usage)
	for i := 0; i < rr.NumField(); i++ {
		if rload, ok := rr.Field(i).Interface().(float64); ok {
			if sload, ok := ss.Field(i).Interface().(float64); ok {
				if rload < sload {
					cost += sload - rload
				}
			}
		}
	}
	return cost
}
