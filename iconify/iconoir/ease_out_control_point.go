package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EaseOutControlPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20S13 4 21 4M7 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 0h2m5 0h-2"/>`),
		g.Group(children),
	)
}