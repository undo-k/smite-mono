package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Affinity string
type Role string
type Mode string

const (
	Physical Affinity = "Physical"
	Magical  Affinity = "Magical"
	Hybrid   Affinity = "Hybrid"
)

const (
	Assassin Role = "Assassin"
	Guardian Role = "Guardian"
	Hunter   Role = "Hunter"
	Mage     Role = "Mage"
	Warrior  Role = "Warrior"
)

const (
	Arena    Mode = "Arena"
	Assault  Mode = "Assault"
	Conquest Mode = "Conquest"
	MOTD     Mode = "MOTD"
	Siege    Mode = "Siege"
	Slash    Mode = "Slash"
)

type God struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Affinity Affinity `json:"affinity"`
	Role     Role     `json:"role"`
}

type Item struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Affinity Affinity `json:"affinity"`
}

type Player struct {
	Id      int    `json:"id"`
	Team    string `json:"team"`
	Won     bool   `json:"won"`
	God     God    `json:"god"`
	Gold    int    `json:"gold"`
	Kills   int    `json:"kills"`
	Deaths  int    `json:"deaths"`
	Assists int    `json:"assists"`
	Items   []Item `json:"items"`
}

type Match struct {
	Id          int      `json:"id"`
	Mode        Mode     `json:"mode"`
	Duration    Duration `json:"duration"`
	Players     []Player `json:"players"`
	WinningTeam string   `json:"winning_team"`
}

func generateId() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var matchId strings.Builder

	for i := 0; i < 10; i++ {
		matchId.WriteString(strconv.Itoa(r.Intn(10)))
	}

	return matchId.String()
}

func (app *Config) generateMatchIds() []string {
	var matchIds []string

	for i := 0; i < 10; i++ {
		matchIds = append(matchIds, generateId())
	}

	return matchIds
}

func generateMatch(matchId string) Match {
	var match Match = Match{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	match.Id, _ = strconv.Atoi(matchId)
	match.Mode = generateMode()
	match.Duration.Duration = (time.Minute * time.Duration((r.Intn(35) + 10))) + (time.Second * time.Duration(r.Intn(60)))
	// match.WinningTeam = generateWinningTeam()
	match.Players = generatePlayers(match.Mode, match.Duration.Duration)

	return match

}

// func generateWinningTeam() string {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	numbie := r.Intn(2)

// 	if numbie == 0 {
// 		return "Chaos"
// 	} else {
// 		return "Order"
// 	}

// }

func generateMode() Mode {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	modes := []Mode{Assault, Arena, Conquest, MOTD, Siege, Slash}
	return modes[r.Intn(6)]
}

func generateItems(role Role) []Item {
	var items = Items()

	itemsForRole := items[getAffinityFromRole(role)]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(itemsForRole), func(i, j int) {
		itemsForRole[i], itemsForRole[j] = itemsForRole[j], itemsForRole[i]
	})

	return itemsForRole[:6]
}

func generateGold(mode Mode, duration time.Duration, kills int, assists int) int {
	var gps int
	switch mode {
	case "Conquest":
		gps = 3
	case "Arena":
		gps = 11
	case "Assault":
		gps = 4
	case "Joust":
		gps = 6
	case "Slash":
		gps = 7
	default:
		gps = 5
	}

	return 1500 + gps*int(duration.Seconds()) + (250 * kills) + (75 * assists)
}

func generatePlayer(mode Mode, duration time.Duration) Player {
	player := Player{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	player.Id, _ = strconv.Atoi(generateId())
	player.Kills = r.Intn(int(duration.Minutes()))
	player.Deaths = r.Intn(10 + int(duration.Minutes()))
	player.Assists = r.Intn(20)*(1+r.Intn(2)) + r.Intn(2)
	player.Gold = generateGold(mode, duration, player.Kills, player.Assists)
	return player
}

func generatePlayers(mode Mode, duration time.Duration) []Player {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	gods := Gods()

	r.Shuffle(len(gods), func(i, j int) {
		gods[i], gods[j] = gods[j], gods[i]
	})

	numPlayers := numPlayersFromMode(mode)
	randomGods := gods[:numPlayers]

	var players []Player

	for i := 0; i < numPlayers; i++ {
		players = append(players, generatePlayer(mode, duration))
		players[i].God = randomGods[i]
		players[i].Items = generateItems(players[i].God.Role)
		if i%2 == 0 {
			players[i].Won = true
		} else {
			players[i].Won = false
		}
	}

	return players
}

func numPlayersFromMode(mode Mode) int {
	switch mode {
	case Assault, Arena, Conquest, Slash, MOTD:
		return 10
	case Siege:
		return 8
	default:
		return 10
	}
}

func getAffinityFromRole(role Role) Affinity {
	switch role {
	case Assassin, Hunter, Warrior:
		return Physical
	default:
		return Magical
	}
}

func Gods() []God {
	var gods = []God{
		{Name: "Achilles", Role: Warrior},
		{Name: "Agni", Role: Mage},
		{Name: "Ah Muzen Cab", Role: Hunter},
		{Name: "Ah Puch", Role: Mage},
		{Name: "Amaterasu", Role: Warrior},
		{Name: "Anhur", Role: Hunter},
		{Name: "Anubis", Role: Mage},
		{Name: "Ao Kuang", Role: Mage},
		{Name: "Aphrodite", Role: Mage},
		{Name: "Apollo", Role: Hunter},
		{Name: "Arachne", Role: Assassin},
		{Name: "Ares", Role: Guardian},
		{Name: "Artemis", Role: Hunter},
		{Name: "Artio", Role: Guardian},
		{Name: "Athena", Role: Guardian},
		{Name: "Atlas", Role: Guardian},
		{Name: "Awilix", Role: Assassin},
		{Name: "Baba Yaga", Role: Mage},
		{Name: "Bacchus", Role: Guardian},
		{Name: "Bakasura", Role: Assassin},
		{Name: "Bake Kujira", Role: Guardian},
		{Name: "Baron Samedi", Role: Mage},
		{Name: "Bastet", Role: Assassin},
		{Name: "Bellona", Role: Warrior},
		{Name: "Cabrakan", Role: Guardian},
		{Name: "Camazotz", Role: Assassin},
		{Name: "Cerberus", Role: Guardian},
		{Name: "Cernunnos", Role: Hunter},
		{Name: "Chaac", Role: Warrior},
		{Name: "Change", Role: Mage},
		{Name: "Charon", Role: Guardian},
		{Name: "Charybdis", Role: Hunter},
		{Name: "Chernobog", Role: Hunter},
		{Name: "Chiron", Role: Hunter},
		{Name: "Chronos", Role: Mage},
		{Name: "Cliodhna", Role: Assassin},
		{Name: "Cthulhu", Role: Guardian},
		{Name: "Cu Chulainn", Role: Warrior},
		{Name: "Cupid", Role: Hunter},
		{Name: "Da Ji", Role: Assassin},
		{Name: "Danzaburou", Role: Hunter},
		{Name: "Discordia", Role: Mage},
		{Name: "Erlang Shen", Role: Warrior},
		{Name: "Eset", Role: Mage},
		{Name: "Fafnir", Role: Guardian},
		{Name: "Fenrir", Role: Assassin},
		{Name: "Freya", Role: Mage},
		{Name: "Ganesha", Role: Guardian},
		{Name: "Geb", Role: Guardian},
		{Name: "Gilgamesh", Role: Warrior},
		{Name: "Guan Yu", Role: Warrior},
		{Name: "Hachiman", Role: Hunter},
		{Name: "Hades", Role: Mage},
		{Name: "He Bo", Role: Mage},
		{Name: "Heimdallr", Role: Hunter},
		{Name: "Hel", Role: Mage},
		{Name: "Hera", Role: Mage},
		{Name: "Hercules", Role: Warrior},
		{Name: "Horus", Role: Warrior},
		{Name: "Hou Yi", Role: Hunter},
		{Name: "Hun Batz", Role: Assassin},
		{Name: "Ishtar", Role: Hunter},
		{Name: "Ix Chel", Role: Mage},
		{Name: "Izanami", Role: Hunter},
		{Name: "Janus", Role: Mage},
		{Name: "Jing Wei", Role: Hunter},
		{Name: "Jormungandr", Role: Guardian},
		{Name: "Kali", Role: Assassin},
		{Name: "Khepri", Role: Guardian},
		{Name: "King Arthur", Role: Warrior},
		{Name: "Kukulkan", Role: Mage},
		{Name: "Kumbhakarna", Role: Guardian},
		{Name: "Kuzenbo", Role: Guardian},
		{Name: "Lancelot", Role: Assassin},
		{Name: "Loki", Role: Assassin},
		{Name: "Maman Brigitte", Role: Mage},
		{Name: "Martichoras", Role: Hunter},
		{Name: "Maui", Role: Guardian},
		{Name: "Medusa", Role: Hunter},
		{Name: "Mercury", Role: Assassin},
		{Name: "Merlin", Role: Mage},
		{Name: "Morgan Le Fay", Role: Mage},
		{Name: "Mulan", Role: Warrior},
		{Name: "Ne Zha", Role: Assassin},
		{Name: "Neith", Role: Hunter},
		{Name: "Nemesis", Role: Assassin},
		{Name: "Nike", Role: Warrior},
		{Name: "Nox", Role: Mage},
		{Name: "Nu Wa", Role: Mage},
		{Name: "Nut", Role: Hunter},
		{Name: "Odin", Role: Warrior},
		{Name: "Olorun", Role: Mage},
		{Name: "Osiris", Role: Warrior},
		{Name: "Pele", Role: Assassin},
		{Name: "Persephone", Role: Mage},
		{Name: "Poseidon", Role: Mage},
		{Name: "Ra", Role: Mage},
		{Name: "Raijin", Role: Mage},
		{Name: "Rama", Role: Hunter},
		{Name: "Ratatoskr", Role: Assassin},
		{Name: "Ravana", Role: Assassin},
		{Name: "Scylla", Role: Mage},
		{Name: "Serqet", Role: Assassin},
		{Name: "Set", Role: Assassin},
		{Name: "Shiva", Role: Warrior},
		{Name: "Skadi", Role: Hunter},
		{Name: "Sobek", Role: Guardian},
		{Name: "Sol", Role: Mage},
		{Name: "Sun Wukong", Role: Warrior},
		{Name: "Surtr", Role: Warrior},
		{Name: "Susano", Role: Assassin},
		{Name: "Sylvanus", Role: Guardian},
		{Name: "Terra", Role: Guardian},
		{Name: "Thanatos", Role: Assassin},
		{Name: "The Morrigan", Role: Mage},
		{Name: "Thor", Role: Assassin},
		{Name: "Thoth", Role: Mage},
		{Name: "Tiamat", Role: Mage},
		{Name: "Tsukuyomi", Role: Assassin},
		{Name: "Tyr", Role: Warrior},
		{Name: "Ullr", Role: Hunter},
		{Name: "Vamana", Role: Warrior},
		{Name: "Vulcan", Role: Mage},
		{Name: "Xbalanque", Role: Hunter},
		{Name: "Xing Tian", Role: Guardian},
		{Name: "Yemoja", Role: Guardian},
		{Name: "Ymir", Role: Guardian},
		{Name: "Yu Huang", Role: Mage},
		{Name: "Zeus", Role: Mage},
		{Name: "Zhong Kui", Role: Mage},
	}

	for i := range gods {
		gods[i].Id = i
		gods[i].Affinity = getAffinityFromRole(gods[i].Role)
	}

	return gods
}

func Items() map[Affinity][]Item {

	var itemMap = make(map[Affinity][]Item)

	var items = []Item{
		{Id: 1, Name: "Absolution", Affinity: Hybrid},
		{Id: 2, Name: "Abyssal Stone", Affinity: Hybrid},
		{Id: 3, Name: "Archdruid's Fury", Affinity: Hybrid},
		{Id: 4, Name: "Archmage's Gem", Affinity: Magical},
		{Id: 5, Name: "Arondight", Affinity: Physical},
		{Id: 6, Name: "Asi", Affinity: Physical},
		{Id: 7, Name: "Atalanta's Bow", Affinity: Physical},
		{Id: 8, Name: "Axe of Animosity", Affinity: Hybrid},
		{Id: 9, Name: "Bancroft's Talon", Affinity: Magical},
		{Id: 10, Name: "Benevolence", Affinity: Hybrid},
		{Id: 11, Name: "Berserker's Shield", Affinity: Hybrid},
		{Id: 12, Name: "Bladed Boomerang", Affinity: Physical},
		{Id: 13, Name: "Bloodforge", Affinity: Physical},
		{Id: 14, Name: "Blood-soaked Shroud", Affinity: Magical},
		{Id: 15, Name: "Bluestone Brooch", Affinity: Physical},
		{Id: 16, Name: "Bluestone Pendant", Affinity: Physical},
		{Id: 17, Name: "Book of Thoth", Affinity: Magical},
		{Id: 18, Name: "Brawler's Beat Stick", Affinity: Physical},
		{Id: 19, Name: "Breastplate of Regrowth", Affinity: Hybrid},
		{Id: 20, Name: "Breastplate of Valor", Affinity: Hybrid},
		{Id: 21, Name: "Bristlebush Acorn", Affinity: Physical},
		{Id: 22, Name: "Bumba's Dagger", Affinity: Hybrid},
		{Id: 23, Name: "Bumba's Hammer", Affinity: Hybrid},
		{Id: 24, Name: "Bumba's Spear", Affinity: Hybrid},
		{Id: 25, Name: "Caduceus Club", Affinity: Physical},
		{Id: 26, Name: "Cannoneer's Cuirass", Affinity: Hybrid},
		{Id: 27, Name: "Charon's Coin", Affinity: Magical},
		{Id: 28, Name: "Chronos' Pendant", Affinity: Magical},
		{Id: 29, Name: "Compassion", Affinity: Hybrid},
		{Id: 30, Name: "Conduit Gem", Affinity: Magical},
		{Id: 31, Name: "Contagion", Affinity: Hybrid},
		{Id: 32, Name: "Corrupted Bluestone", Affinity: Physical},
		{Id: 33, Name: "Cyclopean Ring", Affinity: Magical},
		{Id: 34, Name: "Dawnbringer", Affinity: Physical},
		{Id: 35, Name: "Deathbringer", Affinity: Physical},
		{Id: 36, Name: "Death's Embrace", Affinity: Hybrid},
		{Id: 37, Name: "Death's Temper", Affinity: Hybrid},
		{Id: 38, Name: "Death's Toll", Affinity: Hybrid},
		{Id: 39, Name: "Demon Blade", Affinity: Physical},
		{Id: 40, Name: "Demonic Grip", Affinity: Magical},
		{Id: 41, Name: "Devourer's Gauntlet", Affinity: Physical},
		{Id: 42, Name: "Diamond Arrow", Affinity: Hybrid},
		{Id: 43, Name: "Divine Ruin", Affinity: Magical},
		{Id: 44, Name: "Dominance", Affinity: Physical},
		{Id: 45, Name: "Doom Orb", Affinity: Magical},
		{Id: 46, Name: "Emperor's Armor", Affinity: Hybrid},
		{Id: 47, Name: "Erosion", Affinity: Hybrid},
		{Id: 48, Name: "Ethereal Staff", Affinity: Magical},
		{Id: 49, Name: "Evergreen Acorn", Affinity: Physical},
		{Id: 50, Name: "Eye of the Jungle", Affinity: Hybrid},
		{Id: 51, Name: "Fae-Blessed Hoops", Affinity: Hybrid},
		{Id: 52, Name: "Fail-not", Affinity: Physical},
		{Id: 53, Name: "Fighter's Mask", Affinity: Hybrid},
		{Id: 54, Name: "Frostbound Hammer", Affinity: Physical},
		{Id: 55, Name: "Gauntlet of Thebes", Affinity: Hybrid},
		{Id: 56, Name: "Gem of Focus", Affinity: Magical},
		{Id: 57, Name: "Gem of Isolation", Affinity: Magical},
		{Id: 58, Name: "Genji's Guard", Affinity: Hybrid},
		{Id: 59, Name: "Gilded Arrow", Affinity: Hybrid},
		{Id: 60, Name: "Gladiator's Shield", Affinity: Hybrid},
		{Id: 61, Name: "Golden Blade", Affinity: Physical},
		{Id: 62, Name: "Griffonwing Earrings", Affinity: Hybrid},
		{Id: 63, Name: "Hastened Katana", Affinity: Physical},
		{Id: 64, Name: "Hastened Ring", Affinity: Magical},
		{Id: 65, Name: "Heartseeker", Affinity: Physical},
		{Id: 66, Name: "Heartward Amulet", Affinity: Hybrid},
		{Id: 67, Name: "Heroism", Affinity: Hybrid},
		{Id: 68, Name: "Hunter's Cowl", Affinity: Physical},
		{Id: 69, Name: "Hydra's Lament", Affinity: Physical},
		{Id: 70, Name: "Infused Sigil", Affinity: Physical},
		{Id: 71, Name: "Jotunn's Wrath", Affinity: Physical},
		{Id: 72, Name: "Last Gasp", Affinity: Magical},
		{Id: 73, Name: "Leader's Cowl", Affinity: Physical},
		{Id: 74, Name: "Leather Cowl", Affinity: Physical},
		{Id: 75, Name: "Lono's Mask", Affinity: Hybrid},
		{Id: 76, Name: "Lotus Sickle", Affinity: Hybrid},
		{Id: 77, Name: "Magi's Cloak", Affinity: Hybrid},
		{Id: 78, Name: "Mail of Renewal", Affinity: Hybrid},
		{Id: 79, Name: "Manikin Hidden Blade", Affinity: Hybrid},
		{Id: 80, Name: "Manikin Mace", Affinity: Hybrid},
		{Id: 81, Name: "Manikin Scepter", Affinity: Hybrid},
		{Id: 82, Name: "Manticore's Spikes", Affinity: Hybrid},
		{Id: 83, Name: "Mantle of Discord", Affinity: Hybrid},
		{Id: 84, Name: "Midgardian Mail", Affinity: Hybrid},
		{Id: 85, Name: "Mystical Mail", Affinity: Hybrid},
		{Id: 86, Name: "Obsidian Shard", Affinity: Magical},
		{Id: 87, Name: "Odysseus' Bow", Affinity: Physical},
		{Id: 88, Name: "Oni Hunter's Garb", Affinity: Hybrid},
		{Id: 89, Name: "Ornate Arrow", Affinity: Hybrid},
		{Id: 90, Name: "Pendulum of Ages", Affinity: Magical},
		{Id: 91, Name: "Pestilence", Affinity: Hybrid},
		{Id: 92, Name: "Phalanx", Affinity: Hybrid},
		{Id: 93, Name: "Polynomicon", Affinity: Magical},
		{Id: 94, Name: "Pridwen", Affinity: Hybrid},
		{Id: 95, Name: "Prophetic Cloak", Affinity: Hybrid},
		{Id: 96, Name: "Protector of the Jungle", Affinity: Hybrid},
		{Id: 97, Name: "Protector's Mask", Affinity: Hybrid},
		{Id: 98, Name: "Pythagorem's Piece", Affinity: Magical},
		{Id: 99, Name: "Qin's Sais", Affinity: Physical},
		{Id: 100, Name: "Rage", Affinity: Physical},
		{Id: 101, Name: "Rangda's Mask", Affinity: Hybrid},
		{Id: 102, Name: "Rejuvenating Heart", Affinity: Magical},
		{Id: 103, Name: "Relic Dagger", Affinity: Hybrid},
		{Id: 104, Name: "Rod of Asclepius", Affinity: Magical},
		{Id: 105, Name: "Rod of Tahuti", Affinity: Magical},
		{Id: 106, Name: "Runeforged Hammer", Affinity: Physical},
		{Id: 107, Name: "Sacrificial Shroud", Affinity: Magical},
		{Id: 108, Name: "Sands of Time", Affinity: Magical},
		{Id: 109, Name: "Seer of the Jungle", Affinity: Hybrid},
		{Id: 110, Name: "Sekhmet's Scepter", Affinity: Physical},
		{Id: 111, Name: "Sentinel's Boon", Affinity: Hybrid},
		{Id: 112, Name: "Sentinel's Embrace", Affinity: Hybrid},
		{Id: 113, Name: "Sentinel's Gift", Affinity: Hybrid},
		{Id: 114, Name: "Serrated Edge", Affinity: Physical},
		{Id: 115, Name: "Shadowdrinker", Affinity: Physical},
		{Id: 116, Name: "Shogun's Kusari", Affinity: Hybrid},
		{Id: 117, Name: "Sigil of the Old Guard", Affinity: Physical},
		{Id: 118, Name: "Silverbranch Bow", Affinity: Physical},
		{Id: 119, Name: "Soul Eater", Affinity: Physical},
		{Id: 120, Name: "Soul Gem", Affinity: Magical},
		{Id: 121, Name: "Soul Reaver", Affinity: Magical},
		{Id: 122, Name: "Sovereignty", Affinity: Hybrid},
		{Id: 123, Name: "Spartan Flag", Affinity: Hybrid},
		{Id: 124, Name: "Spear of Desolation", Affinity: Magical},
		{Id: 125, Name: "Spear of the Magus", Affinity: Magical},
		{Id: 126, Name: "Spectral Armor", Affinity: Hybrid},
		{Id: 127, Name: "Sphinx's Baubles", Affinity: Hybrid},
		{Id: 128, Name: "Spirit Robe", Affinity: Hybrid},
		{Id: 129, Name: "Staff of Myrddin", Affinity: Magical},
		{Id: 130, Name: "Stone of Binding", Affinity: Hybrid},
		{Id: 131, Name: "Stone of Gaia", Affinity: Hybrid},
		{Id: 132, Name: "Sundering Axe", Affinity: Hybrid},
		{Id: 133, Name: "Tablet of Destinies", Affinity: Magical},
		{Id: 134, Name: "Tainted Amulet", Affinity: Hybrid},
		{Id: 135, Name: "Tainted Breastplate", Affinity: Hybrid},
		{Id: 136, Name: "Tainted Steel", Affinity: Hybrid},
		{Id: 137, Name: "Talisman of Energy", Affinity: Hybrid},
		{Id: 138, Name: "Telkhines Ring", Affinity: Magical},
		{Id: 139, Name: "The Alternate Timeline", Affinity: Magical},
		{Id: 140, Name: "The Crusher", Affinity: Physical},
		{Id: 141, Name: "The Executioner", Affinity: Physical},
		{Id: 142, Name: "Thickbark Acorn", Affinity: Physical},
		{Id: 143, Name: "Thistlebark Acorn", Affinity: Physical},
		{Id: 144, Name: "Titan's Bane", Affinity: Physical},
		{Id: 145, Name: "Toxic Blade", Affinity: Hybrid},
		{Id: 146, Name: "Transcendence", Affinity: Physical},
		{Id: 147, Name: "Typhon's Fang", Affinity: Magical},
		{Id: 148, Name: "Vampiric Shroud", Affinity: Magical},
		{Id: 149, Name: "Vital Amplifier", Affinity: Physical},
		{Id: 150, Name: "War Banner", Affinity: Hybrid},
		{Id: 151, Name: "Warding Sigil", Affinity: Physical},
		{Id: 152, Name: "War Flag", Affinity: Hybrid},
		{Id: 153, Name: "Warlock's Staff", Affinity: Magical},
		{Id: 154, Name: "Warrior's Axe", Affinity: Hybrid},
		{Id: 155, Name: "Winged Blade", Affinity: Hybrid},
	}

	for _, item := range items {
		itemMap[item.Affinity] = append(itemMap[item.Affinity], item)
	}

	itemMap[Physical] = append(itemMap[Physical], itemMap[Hybrid]...)
	itemMap[Magical] = append(itemMap[Magical], itemMap[Hybrid]...)

	return itemMap
}
