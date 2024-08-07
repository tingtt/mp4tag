package mp4tag

// GenreNone
type Genre int8

const (
	GenreNone Genre = iota
	GenreBlues
	GenreClassicRock
	GenreCountry
	GenreDance
	GenreDisco
	GenreFunk
	GenreGrunge
	GenreHipHop
	GenreJazz
	GenreMetal
	GenreNewAge
	GenreOldies
	GenreOther
	GenrePop
	GenreRhythmAndBlues
	GenreRap
	GenreReggae
	GenreRock
	GenreTechno
	GenreIndustrial
	GenreAlternative
	GenreSka
	GenreDeathMetal
	GenrePranks
	GenreSoundtrack
	GenreEurotechno
	GenreAmbient
	GenreTripHop
	GenreVocal
	GenreJassAndFunk
	GenreFusion
	GenreTrance
	GenreClassical
	GenreInstrumental
	GenreAcid
	GenreHouse
	GenreGame
	GenreSoundClip
	GenreGospel
	GenreNoise
	GenreAlternativeRock
	GenreBass
	GenreSoul
	GenrePunk
	GenreSpace
	GenreMeditative
	GenreInstrumentalPop
	GenreInstrumentalRock
	GenreEthnic
	GenreGothic
	GenreDarkwave
	GenreTechnoindustrial
	GenreElectronic
	GenrePopFolk
	GenreEurodance
	GenreSouthernRock
	GenreComedy
	GenreCull
	GenreGangsta
	GenreTop40
	GenreChristianRap
	GenrePopSlashFunk
	GenreJungleMusic
	GenreNativeUS
	GenreCabaret
	GenreNewWave
	GenrePsychedelic
	GenreRave
	GenreShowtunes
	GenreTrailer
	GenreLofi
	GenreTribal
	GenreAcidPunk
	GenreAcidJazz
	GenrePolka
	GenreRetro
	GenreMusical
	GenreRockNRoll
	GenreHardRock
)

var resolveGenre = map[uint8]Genre{
	1:  GenreBlues,
	2:  GenreClassicRock,
	3:  GenreCountry,
	4:  GenreDance,
	5:  GenreDisco,
	6:  GenreFunk,
	7:  GenreGrunge,
	8:  GenreHipHop,
	9:  GenreJazz,
	10: GenreMetal,
	11: GenreNewAge,
	12: GenreOldies,
	13: GenreOther,
	14: GenrePop,
	15: GenreRhythmAndBlues,
	16: GenreRap,
	17: GenreReggae,
	18: GenreRock,
	19: GenreTechno,
	20: GenreIndustrial,
	21: GenreAlternative,
	22: GenreSka,
	23: GenreDeathMetal,
	24: GenrePranks,
	25: GenreSoundtrack,
	26: GenreEurotechno,
	27: GenreAmbient,
	28: GenreTripHop,
	29: GenreVocal,
	30: GenreJassAndFunk,
	31: GenreFusion,
	32: GenreTrance,
	33: GenreClassical,
	34: GenreInstrumental,
	35: GenreAcid,
	36: GenreHouse,
	37: GenreGame,
	38: GenreSoundClip,
	39: GenreGospel,
	40: GenreNoise,
	41: GenreAlternativeRock,
	42: GenreBass,
	43: GenreSoul,
	44: GenrePunk,
	45: GenreSpace,
	46: GenreMeditative,
	47: GenreInstrumentalPop,
	48: GenreInstrumentalRock,
	49: GenreEthnic,
	50: GenreGothic,
	51: GenreDarkwave,
	52: GenreTechnoindustrial,
	53: GenreElectronic,
	54: GenrePopFolk,
	55: GenreEurodance,
	56: GenreSouthernRock,
	57: GenreComedy,
	58: GenreCull,
	59: GenreGangsta,
	60: GenreTop40,
	61: GenreChristianRap,
	62: GenrePopSlashFunk,
	63: GenreJungleMusic,
	64: GenreNativeUS,
	65: GenreCabaret,
	66: GenreNewWave,
	67: GenrePsychedelic,
	68: GenreRave,
	69: GenreShowtunes,
	70: GenreTrailer,
	71: GenreLofi,
	72: GenreTribal,
	73: GenreAcidPunk,
	74: GenreAcidJazz,
	75: GenrePolka,
	76: GenreRetro,
	77: GenreMusical,
	78: GenreRockNRoll,
	79: GenreHardRock,
}

var resolveGenreName = map[uint8]string{
	1:  "Blues",
	2:  "ClassicRock",
	3:  "Country",
	4:  "Dance",
	5:  "Disco",
	6:  "Funk",
	7:  "Grunge",
	8:  "HipHop",
	9:  "Jazz",
	10: "Metal",
	11: "NewAge",
	12: "Oldies",
	13: "Other",
	14: "Pop",
	15: "RhythmAndBlues",
	16: "Rap",
	17: "Reggae",
	18: "Rock",
	19: "Techno",
	20: "Industrial",
	21: "Alternative",
	22: "Ska",
	23: "DeathMetal",
	24: "Pranks",
	25: "Soundtrack",
	26: "Eurotechno",
	27: "Ambient",
	28: "TripHop",
	29: "Vocal",
	30: "JassAndFunk",
	31: "Fusion",
	32: "Trance",
	33: "Classical",
	34: "Instrumental",
	35: "Acid",
	36: "House",
	37: "Game",
	38: "SoundClip",
	39: "Gospel",
	40: "Noise",
	41: "AlternativeRock",
	42: "Bass",
	43: "Soul",
	44: "Punk",
	45: "Space",
	46: "Meditative",
	47: "InstrumentalPop",
	48: "InstrumentalRock",
	49: "Ethnic",
	50: "Gothic",
	51: "Darkwave",
	52: "Technoindustrial",
	53: "Electronic",
	54: "PopFolk",
	55: "Eurodance",
	56: "SouthernRock",
	57: "Comedy",
	58: "Cull",
	59: "Gangsta",
	60: "Top40",
	61: "ChristianRap",
	62: "PopSlashFunk",
	63: "JungleMusic",
	64: "NativeUS",
	65: "Cabaret",
	66: "NewWave",
	67: "Psychedelic",
	68: "Rave",
	69: "Showtunes",
	70: "Trailer",
	71: "Lofi",
	72: "Tribal",
	73: "AcidPunk",
	74: "AcidJazz",
	75: "Polka",
	76: "Retro",
	77: "Musical",
	78: "RockNRoll",
	79: "HardRock",
}

func ResolveGenreName(genre Genre) string {
	return resolveGenreName[uint8(genre)]
}
