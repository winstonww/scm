package core

const (
	NUM_CONTROLLERS int     = 5      // Number of controlelrs in this experiment
	NUM_SWITCHES    int     = 50     // Number of switches in this experiment
	CMIN            float64 = 0.6    // The minimal initial cpu sources allocated to controller, [0,1]
	NMIN            float64 = 0.4    // The minimal initial bandwidth sources allocated to controller, [0,1]
	CTOTAL          float64 = 1000.0 //The total CPU resources combined in all controllers
	NTOTAL          float64 = 800.0  //The total bandwidth in all controllers
	CUTIL           float64 = 0.4    // The total Utilization of CPU resources
	NUTIL           float64 = 0.5    //The total Utilization of bandwidth resources
	DEBUG           bool    = false  // for debugging
)
