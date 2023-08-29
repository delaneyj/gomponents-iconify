package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreaterThanOrEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.435 1.704l17.3 6.796l-17.3 6.796l-.731-1.861L16.265 8.5L3.704 3.565l.73-1.861ZM3 19h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}