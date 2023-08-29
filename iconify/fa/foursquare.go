package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m968 434l37-194q5-23-9-40t-35-17H249q-23 0-38.5 17T195 237v1101q0 7 6 1l291-352q23-26 38-33.5t48-7.5h239q22 0 37-14.5t18-29.5q24-130 37-191q4-21-11.5-40T861 652H567q-29 0-48-19t-19-48v-42q0-29 19-47.5t48-18.5h346q18 0 35-13.5t20-29.5zm227-222q-15 73-53.5 266.5t-69.5 350t-35 173.5q-6 22-9 32.5t-14 32.5t-24.5 33t-38.5 21t-58 10H622q-13 0-22 10q-8 9-426 494q-22 25-58.5 28.5T67 1658q-55-22-55-98V150q0-55 38-102.5T170 0h888q95 0 127 53t10 159zm0 0l-158 790q4-17 35-173.5t69.5-350T1195 212z"/>`),
		g.Group(children),
	)
}