package v1

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"

	"github.com/ToshihitoKon/codename-generator/utils"
)

// 前置詞
var Prepositions []string

var adjectives = []string{
	"Awesome", "Brilliant", "Fantastic", "Gorgeous", "Happy", "Incredible", "Joyful",
	"Kind", "Lovely", "Magnificent", "Nice", "Optimistic", "Perfect", "Radiant",
	"Spectacular", "Terrific", "Vibrant", "Wonderful", "Adventurous", "Benevolent",
	"Courageous", "Dazzling", "Elegant", "Friendly", "Gentle", "Harmonious",
	"Inspiring", "Jubilant", "Lively", "Miraculous", "Nurturing", "Passionate",
	"Resilient", "Sincere", "Thrilling", "Vivacious", "Youthful", "Blissful",
	"Charming", "Delightful", "Exquisite", "Grateful", "Heavenly", "Intelligent",
	"Joyous", "Luminous", "Marvelous", "Peaceful", "Refreshing", "Sunny",
	"Tranquil", "Victorious",
}

var colors = []string{
	"Red", "Orange", "Yellow", "Green", "Blue", "Purple", "Pink", "Magenta", "Cyan",
	"Turquoise", "Lime", "Teal", "Indigo", "Violet", "Maroon", "Olive", "Navy",
	"Aquamarine", "Coral", "Gold", "Silver", "Gray", "Black", "White", "Beige",
	"Brown", "Tan", "Cream", "Lavender", "Mint", "Peach", "Salmon", "Slate", "Ruby",
	"Emerald", "Amber", "Sapphire", "Rose", "Ivory", "Pearl", "Crimson", "Azure",
	"Plum", "Jade", "Onyx", "Sunflower", "Orchid", "Lilac", "Mauve", "Cobalt",
}

var seasons = []string{
	"Spring", "Summer", "Autumn", "Winter", "Rainy", "Sunny", "Breezy", "Frosty",
	"Mild", "Chilly", "Hot", "Humid", "Dry", "Windy", "Snowy", "Stormy", "Gloomy",
	"Balmy", "Crystalline", "Foggy", "Misty", "Cloudy", "Icy", "Sizzling",
	"Blustery", "Leafy", "Blossomy", "Harvest", "Golden", "Freezing", "Tropical",
	"Pleasant", "Shivering", "Sweaty", "Dewy", "Thundery", "Brisk", "Cozy", "Serene",
	"Whirling", "Dusty", "Blazing", "Zesty", "Rainbow", "Damp", "Lush",
	"Frostbitten", "Zephyr", "Arctic", "Heatwave", "Hazy", "Gusty",
}

var spaceTerms = []string{
	"Galaxy", "Star", "Planet", "Moon", "Asteroid", "Comet", "Nebula", "Black Hole",
	"Supernova", "Cosmic Rays", "Interstellar", "Celestial", "Milky Way", "Gravity",
	"Orbit", "Solar System", "Satellite", "Spacecraft", "Astrobiology", "Telescope",
	"Exoplanet", "Dark Matter", "Gamma Rays", "Astronomy", "Cosmonaut", "Astronaut",
	"Lunar", "Rocket", "Space Station", "Constellation", "Interplanetary", "Stellar",
	"Astrophysics", "Astrochemistry", "Cosmos", "Intergalactic", "Lagrangian Point",
	"Planetary System", "Astrogeology", "Space Probe", "Spacewalk",
	"Gravitational Waves", "Space-Time", "Extragalactic", "Interstellar Medium",
	"Parallax", "Space Shuttle", "Cosmic Microwave Background", "Ecliptic",
	"Celestial Sphere", "Zodiac", "Orbital Mechanics", "Stellar Evolution",
}

func init() {
	Prepositions = append(Prepositions, adjectives...)
	Prepositions = append(Prepositions, colors...)
	Prepositions = append(Prepositions, seasons...)
	Prepositions = append(Prepositions, spaceTerms...)
}

// 名前
var Names []string

var sweets = []string{
	"Cookie", "Chocolate", "Candy", "Marshmallow", "Gummy bear", "Jelly bean",
	"Cupcake", "Donut", "Brownie", "Ice cream", "Popsicle", "Lollipop", "Caramel",
	"Toffee", "Macaron", "Peanut brittle", "Fudge", "Cotton candy", "Rock candy",
	"Jawbreaker", "Muffin", "Pretzel", "Pie", "Churro", "Taffy", "Nougat", "Popcorn",
	"Tiramisu", "Cheesecake", "Pancake", "Waffle", "Banana bread", "Gingerbread",
	"Scone", "Shortbread", "Fruitcake", "Custard", "Milkshake", "Pudding",
	"Gelato", "Sorbet", "Sherbet", "Praline", "Truffle", "Caramel apple",
	"Caramel corn", "Sugar cookie", "Peanut butter cup", "Candy cane", "Jelly roll",
	"Baklava", "Halva", "Oreo", "Twix", "Kit Kat", "Snickers", "Skittles",
	"Reeses", "Hersheys", "Milky Way", "Starburst", "Nerds", "Smarties",
	"Gobstopper", "Sour Patch Kids", "Pixy Stix", "Airheads", "Twizzlers",
	"Sour gummy worms", "Jolly Rancher", "Ring Pop", "Peppermint", "Butterscotch",
	"Licorice", "Mints", "Circus peanuts", "Jawbreakers", "Caramel chew",
	"Caramel apple", "Tootsie Roll", "Cracker Jack", "Nutter Butter", "Smores",
	"Candy corn", "Peeps", "Laffy Taffy", "Butterfinger", "Almond Joy", "Rolo",
	"Swedish Fish", "Pocky", "Bubble gum", "Chewing gum", "Fruit roll-up",
	"Fruit by the Foot",
}

var elements = []string{
	"Air", "Water", "Fire", "Earth", "Wind", "Rain", "Snow", "Sun", "Moon", "Star",
	"Cloud", "Mountain", "Valley", "Ocean", "Wave", "River", "Lake", "Island",
	"Forest", "Tree", "Flower", "Grass", "Rock", "Cave", "Desert", "Sand",
	"Volcano", "Thunder", "Lightning", "Rainbow", "Sunrise", "Sunset", "Sky",
	"Hill", "Garden", "Field", "Leaf", "Breeze", "Fog", "Mist", "Pebble", "Pond",
	"Glacier", "Dew", "Haze", "Cliff", "Tide", "Meadow", "Canyon", "Woods",
	"Spring", "Autumn", "Winter", "Summer",
}

var gemstones = []string{
	"Diamond", "Ruby", "Emerald", "Sapphire", "Amethyst", "Topaz", "Opal", "Pearl",
	"Garnet", "Aquamarine", "Peridot", "Tourmaline", "Citrine", "Tanzanite",
	"Onyx", "Jade", "Turquoise", "Lapis Lazuli", "Moonstone", "Agate", "Quartz",
	"Malachite", "Coral", "Alexandrite", "Spinel", "Zircon", "Obsidian",
	"Labradorite", "Carnelian", "Iolite", "Rhodolite", "Apatite", "Amber",
	"Fluorite", "Hematite", "Jasper", "Kunzite", "Morganite", "Sunstone",
	"Tigers Eye", "Chrysoprase", "Sardonyx", "Serpentine", "Sodalite",
	"Uvarovite", "Variscite", "Wulfenite", "Zoisite", "Demantoid", "Hiddenite",
	"Kyanite", "Pietersite",
}

var celestialBodies = []string{
	"Sun", "Moon", "Mercury", "Venus", "Mars", "Jupiter", "Saturn", "Uranus",
	"Neptune", "Pluto", "Earth", "Ceres", "Eris", "Haumea", "Makemake", "Sedna",
	"Io", "Europa", "Ganymede", "Callisto", "Titan", "Enceladus", "Mimas", "Phobos",
	"Deimos", "Triton", "Charon", "Eris", "Oberon", "Titania", "Miranda", "Ariel",
	"Umbriel", "Triton", "Proteus", "Hyperion", "Dione", "Rhea", "Iapetus", "Tethys",
	"Enceladus", "Mimas", "Phobos", "Deimos", "Triton", "Charon", "Eris", "Oberon",
	"Titania", "Miranda", "Ariel", "Umbriel",
}

var instruments = []string{
	"Piano", "Guitar", "Violin", "Drums", "Flute", "Trumpet", "Saxophone", "Cello",
	"Bass", "Harp", "Clarinet", "Trombone", "Accordion", "Oboe", "Banjo",
	"Xylophone", "Sitar", "Ukulele", "Harmonica", "Bagpipes", "Marimba",
	"Didgeridoo", "Theremin", "Mandolin", "Cajon", "Djembe", "Kalimba", "Balalaika",
	"Bassoon", "Tabla", "Erhu", "Guzheng", "Taiko", "Koto", "Pipa", "Ney",
	"Shamisen", "Bongo", "Dhol", "Gamelan", "Steelpan", "Santoor", "Zither",
	"Nyckelharpa", "Gaida", "Duduk", "Ngoni", "Ghatam", "Sarod", "Cymbals",
}

func init() {
	Names = append(Names, sweets...)
	Names = append(Names, elements...)
	Names = append(Names, gemstones...)
	Names = append(Names, celestialBodies...)
	Names = append(Names, instruments...)
}

type Generator struct{}

func New() utils.CodenameGenerator {
	return &Generator{}
}

func (_ *Generator) GenerateCodename(text string) (string, error) {
	textHashBytes := sha256.Sum256([]byte(text))
	randSourceInt64, readBytes := binary.Varint(textHashBytes[:])
	if readBytes < 0 {
		return "", errors.New("error: overflow")
	}

	r := rand.New(rand.NewSource(randSourceInt64))
	return fmt.Sprintf(
		"%s %s",
		Prepositions[r.Intn(len(Prepositions))],
		Names[r.Intn(len(Names))],
	), nil
}
