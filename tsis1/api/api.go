package api


type Driver struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Country string `json:"country"`
	Team string `json:"team"`
	Age int `json:"age"`
}

var Drivers = []Driver{
	{"Lewis", "Hamilton", "United Kingdom", "Mercedes", 37},
	{"Max", "Verstappen", "Netherlands", "Red Bull Racing", 24},
	{"Valtteri", "Bottas", "Finland", "Mercedes", 32},
	{"Lando", "Norris", "United Kingdom", "McLaren", 22},
	{"Charles", "Leclerc", "Monaco", "Ferrari", 24},
	{"Carlos", "Sainz", "Spain", "Ferrari", 27},
	{"Daniel", "Ricciardo", "Australia", "McLaren", 32},
	{"Sebastian", "Vettel", "Germany", "Aston Martin", 34},
	{"Fernando", "Alonso", "Spain", "Alpine", 40},
	{"Pierre", "Gasly", "France", "AlphaTauri", 26},
	{"Esteban", "Ocon", "France", "Alpine", 25},
	{"Kimi", "Räikkönen", "Finland", "Alfa Romeo Racing", 42},
	{"Antonio", "Giovinazzi", "Italy", "Alfa Romeo Racing", 28},
	{"Nikita", "Mazepin", "Russia", "Haas", 23},
	{"Mick", "Schumacher", "Germany", "Haas", 23},
	{"Nicholas", "Latifi", "Canada", "Williams", 26},
	{"George", "Russell", "United Kingdom", "Williams", 24},
	{"Yuki", "Tsunoda", "Japan", "AlphaTauri", 21},
	{"Robert", "Kubica", "Poland", "Alfa Romeo Racing", 37},
	{"Nikolas", "Hülkenberg", "Germany", "Aston Martin", 34},
}

func GetDriverByFirstName(name string) Driver {
	for _, driver := range Drivers {
		if driver.Firstname == name {
			return driver
		}
	}
	return Driver{"", "", "", "", 0}
}