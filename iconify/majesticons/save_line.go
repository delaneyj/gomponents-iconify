package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H4v16h16V8l-4-4h-2M8 4v4h6V4M8 4h6m-2 8c-.667 0-2 .4-2 2s1.333 2 2 2s2-.4 2-2s-1.333-2-2-2z"/>`),
		g.Group(children),
	)
}