package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitBranchLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.105 15.21A3.001 3.001 0 1 1 5 15.17V8.83a3.001 3.001 0 1 1 2 0V12c.836-.628 1.874-1 3-1h4a3.001 3.001 0 0 0 2.895-2.21a3.001 3.001 0 1 1 2.032.064A5.002 5.002 0 0 1 14 13h-4a3.001 3.001 0 0 0-2.895 2.21ZM6 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM6 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm12 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}