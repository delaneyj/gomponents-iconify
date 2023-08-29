package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugBreakpointFunctionUnverified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 11h8L8 4l-4 7zm2.154-1.25h3.692L8 6.52L6.154 9.75z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}