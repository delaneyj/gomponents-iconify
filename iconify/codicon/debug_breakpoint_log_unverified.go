package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugBreakpointLogUnverified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.02 7.98L8 3l4.98 4.98L8 12.96L3.02 7.98zM8 10.77l2.79-2.79L8 5.19L5.21 7.98L8 10.77z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}