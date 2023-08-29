package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerspectiveBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 116h-12V48a20 20 0 0 0-23.58-19.67l-160 29.09A20 20 0 0 0 28 77.09V116H16a12 12 0 0 0 0 24h12v38.91a20 20 0 0 0 16.42 19.67l160 29.09A20 20 0 0 0 228 208v-68h12a12 12 0 0 0 0-24ZM52 80.43L204 52.8V116H52ZM204 203.2L52 175.57V140h152Z"/>`),
		g.Group(children),
	)
}