package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShardSword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m496.514 14.453l-97.79 45.47l-71.57 71.567l-26.558-33.377l-31.772 91.705l-74.644 74.645l-26.56-33.375l-29.968 86.504l20.053 20.05L413.415 81.93l13.216 13.215L170.92 350.86l20.756 20.755l75.763-75.765l83.155-28.807l-30.266-24.084l80.082-80.083l83.158-28.81l-31.148-24.786l44.096-94.827zM43.57 266.525l-.002 39.75l23.793-.334l137.732 137.734l-.336 23.79h39.754l-.28-39.538h-22.673L83.11 289.48v-22.675l-39.54-.28zm71.434 113.49l-88.166 88.167l17.738 17.74l88.168-88.168l-17.74-17.738z"/>`),
		g.Group(children),
	)
}