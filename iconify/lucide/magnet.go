package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 15l-4-4l6.75-6.77a7.79 7.79 0 0 1 11 11L13 22l-4-4l6.39-6.36a2.14 2.14 0 0 0-3-3L6 15M5 8l4 4m3 3l4 4"/>`),
		g.Group(children),
	)
}