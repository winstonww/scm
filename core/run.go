package core

// http://drops.dagstuhl.de/opus/volltexte/2017/7466/pdf/LIPIcs-ICALP-2017-125.pdf
// use Spanning Tree as a guide

import (
	"fmt"
)

func Run() {

	controllers, migrations := Initialize(
		NUM_CONTROLLERS, NUM_SWITCHES,
		CMIN, CTOTAL, CUTIL, NMIN, NTOTAL, NUTIL)

	// before migration
	for _, c := range controllers {
		fmt.Println(c)
	}

	GreedyMigrate(migrations, controllers)

	// after migration
	for _, c := range controllers {
		fmt.Println(c)
	}
}
