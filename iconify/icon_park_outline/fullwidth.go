package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullwidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 6v36M17 19l-5 5m0 0l5 5m-5-5h24m-5-5l5 5m0 0l-5 5M6 6v36"/>`),
		g.Group(children),
	)
}