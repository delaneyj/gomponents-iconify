package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapBoxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.005 5.003v14h16v-14h-16Zm-1-2h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1Zm12 4l3.5 3l-3.5 3v-2h-4v-2h4v-2Zm-6 10l-3.5-3l3.5-3v2h4v2h-4v2Z"/>`),
		g.Group(children),
	)
}