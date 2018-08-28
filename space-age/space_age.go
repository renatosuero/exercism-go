package space

// Planet is a place where we use to calculate age
type Planet string

// base are the seconds for one year in the earth
const base float64 = 31557600.0

// Age calculate how old someone would be on planets
func Age(seconds float64, planet Planet) float64 {
	return (seconds / base) / getYearOfPlanet(planet)
}

// getYearOfPlanet returns seconds by planet
func getYearOfPlanet(planet Planet) float64 {
	planets := map[Planet]float64{}
	planets["Earth"] = 1
	planets["Mercury"] = 0.2408467
	planets["Venus"] = 0.61519726
	planets["Mars"] = 1.8808158
	planets["Jupiter"] = 11.862615
	planets["Saturn"] = 29.447498
	planets["Uranus"] = 84.016846
	planets["Neptune"] = 164.79132
	return planets[planet]
}
