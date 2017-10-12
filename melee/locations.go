package melee

var Locations = []string{
	"004530E0",
	"00453F70",
	"0045310E",
	"00453F9E",
	"004530C0",
	"00453F50",
	"00453090",
	"00453F20",
	"00453094",
	"00453F24",
	"00453130 70",
	"00453130 20CC",
	"00453130 8F4",
	"00453130 8FC",
	"00453130 19EC",
	"00453130 19BC",
	"00453130 23a0",
	"00453130 2174",
	"00453130 19C8",
	"00453130 140",
	"00453130 E0",
	"00453130 E4",
	"00453130 EC",
	"00453130 F0",
	"00453130 14C",
	"00453FC0 70",
	"00453FC0 20CC",
	"00453FC0 8F4",
	"00453FC0 8FC",
	"00453FC0 19EC",
	"00453FC0 19BC",
	"00453FC0 23a0",
	"00453FC0 2174",
	"00453FC0 19C8",
	"00453FC0 140",
	"00453FC0 E0",
	"00453FC0 E4",
	"00453FC0 EC",
	"00453FC0 F0",
	"00453FC0 14C",
	"00454E50 70",
	"00454E50 20CC",
	"00454E50 8F4",
	"00454E50 8FC",
	"00454E50 19EC",
	"00454E50 19BC",
	"00454E50 23a0",
	"00454E50 2174",
	"00454E50 19C8",
	"00454E50 140",
	"00454E50 E0",
	"00454E50 E4",
	"00454E50 EC",
	"00454E50 F0",
	"00454E50 14C",
	"00455CE0 70",
	"00455CE0 20CC",
	"00455CE0 8F4",
	"00455CE0  8FC",
	"00455CE0 19EC",
	"00455CE0 19BC",
	"00455CE0 23a0",
	"00455CE0 2174",
	"00455CE0 19C8",
	"00455CE0 140",
	"00455CE0 E0",
	"00455CE0 E4",
	"00455CE0 EC",
	"00455CE0 F0",
	"00455CE0 14C",
	"003F0E0A",
	"003F0E2E",
	"00479d30",
	"004D6CAD",
	"0111826C",
	"01118270",
	"00479D60",
	"01118DEC",
	"01118DF0",
	"011176EC",
	"011176F0",
	"01116B6C",
	"01116B70",
	"003F0E52",
	"003F0E76",
	"00454E00",
	"00455C90",
	"00454E2E",
	"00455CBE",
	"00454DE0",
	"00455C70",
	"00454DB0",
	"00455C40",
	"00454DB4",
	"00455C44",

	// Character select screen cursor token status for each player
	"004a0bc0 2",
	"004a0bc4 2",
	"004a0bc8 2",
	"004a0bcc 2",
	//"00bda810 28 38",
	//"00bda810 28 3c",
	"004d6cf2",

	//Character controller type
	"003F0E08",
	"003F0E2C",
	"003F0E50",
	"003F0E74",

	// Hitbox data
	"00453130 990",
	"00453130 AC8",
	"00453130 C00",
	"00453130 D38",
	"00453130 974",
	"00453130 AAC",
	"00453130 BE4",
	"00453130 D1C",
	"00453130 9C4",
	"00453130 9C0",
	"00453130 AF8",
	"00453130 AFC",
	"00453130 C30",
	"00453130 C34",
	"00453130 D68",
	"00453130 D6C",
	"00453130 2278",
	"00453FC0 2278",
	"00454E50 2278",
	"00455CE0 2278",

	// Transformed
	"0045308C",
	"00453F1C",
	"00454DAC",
	"00455C3C",

	// Sub players
	"00453134 70",
	"00453134 20CC",
	"00453134 8F4",
	"00453134 19EC",
	"00453134 19BC",
	"00453134 23a0",
	"00453134 2174",
	"00453134 19C8",
	"00453134 140",
	"00453134 E0",
	"00453134 E4",
	"00453134 EC",
	"00453134 F0",
	"00453134 14C",
	"00453134 1890",
	"00453134 8C",
	"00453134 110",
	"00453134 114",
	"00453FC4 70",
	"00453FC4 20CC",
	"00453FC4 8F4",
	"00453FC4 19EC",
	"00453FC4 19BC",
	"00453FC4 23a0",
	"00453FC4 2174",
	"00453FC4 19C8",
	"00453FC4 140",
	"00453FC4 E0",
	"00453FC4 E4",
	"00453FC4 EC",
	"00453FC4 F0",
	"00453FC4 14C",
	"00453FC4 1890",
	"00453FC4 8C",
	"00453FC4 110",
	"00453FC4 114",
	"00454E54 70",
	"00454E54 20CC",
	"00454E54 8F4",
	"00454E54 19EC",
	"00454E54 19BC",
	"00454E54 23a0",
	"00454E54 2174",
	"00454E54 19C8",
	"00454E54 140",
	"00454E54 E0",
	"00454E54 E4",
	"00454E54 EC",
	"00454E54 F0",
	"00454E54 14C",
	"00454E54 1890",
	"00454E54 8C",
	"00454E54 110",
	"00454E54 114",
	"00455CE4 70",
	"00455CE4 20CC",
	"00455CE4 8F4",
	"00455CE4 19EC",
	"00455CE4 19BC",
	"00455CE4 23a0",
	"00455CE4 2174",
	"00455CE4 19C8",
	"00455CE4 140",
	"00455CE4 E0",
	"00455CE4 E4",
	"00455CE4 EC",
	"00455CE4 F0",
	"00455CE4 14C",
	"00455CE4 1890",
	"00455CE4 8C",
	"00455CE4 110",
	"00455CE4 114",
	"00455CE4 1890",
	"00453084",
	"00453F14",
	"00454DA4",
	"00455C34",

	// # controller inputs
	"004C1FAC",
	"004C1FF0",
	"004C2034",
	"004C2078",
}
