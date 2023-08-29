package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotaryRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3 3v10m9-9l4-4l-4-4m-6 4h10"/>`),
		g.Group(children),
	)
}