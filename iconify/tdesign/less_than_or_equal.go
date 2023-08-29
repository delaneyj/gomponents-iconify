package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LessThanOrEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.297 3.565L7.735 8.5l12.562 4.935l-.731 1.861L2.266 8.5l17.3-6.796l.73 1.861ZM3 19h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}