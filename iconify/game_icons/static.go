package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Static(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M142.7 18.72L20 77.56v20.72l120.3-57.72L20 132v112.7l129.4-46.8l-63.4-9.3l188-45.8L20 326.5V370l349.6-189.6l-81.1 18.3l115.2-71.4l91.6-34.58v-20C432 101.4 368 115.2 301 131.6l44.4-30.2l149.9-72.81v-9.87h-22.5C393.2 63.83 310.6 89.86 224.7 119.7L383.5 18.72H302L112.3 114.1l122.2-95.38zM495.3 143.3L230.6 299.6l150.5-39.1L78.72 418.7l102.38-90.8L20 409v85h32.94S223 392.8 317.3 354.7L247 407l112.6-34.4l-216.2 119.5l286-102.1l-82 57.6L495.3 406V298.4l-139.9 52.3l72.7-58.5l-133.7 47.2l177-115.1l-24.9 42.5l48.8-18.5v-105z"/>`),
		g.Group(children),
	)
}