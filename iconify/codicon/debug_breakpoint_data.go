package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugBreakpointData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m12.238 8l-2.31 4H5.31L3 8l2.31-4h4.618l2.31 4z"/>`),
		g.Group(children),
	)
}