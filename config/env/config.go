package env

type (

	CONFIG struct {
		MICRO MICRO
	}

	MICRO struct {
		// DB DB
		API API
	}

	// DB struct {
	// 	Psql PsqlEnv
	// }

	API struct {
		HOST string
		PORT string
	}

	// PsqlEnv struct {
	// 	Host string
	// 	Port string
	// 	User string
	// 	Name string
	// 	Pass string
	// }
)