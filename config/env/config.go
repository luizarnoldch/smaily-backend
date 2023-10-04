package env

type (

	CONFIG struct {
		MICRO MICRO
	}

	MICRO struct {
		DB DB
		API API
	}

	DB struct {
		PSQL PSQL
	}

	API struct {
		HOST string
		PORT string
	}

	PSQL struct {
		HOST string
		PORT string
		USER string
		PASS string
		SCHEMA string
	}
)