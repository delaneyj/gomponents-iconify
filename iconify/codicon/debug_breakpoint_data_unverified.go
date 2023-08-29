package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugBreakpointDataUnverified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9.931 4h-4.62l-2.31 4l2.31 4h4.62l2.31-4l-2.31-4zm-.75 6.7h-3.12L4.501 8l1.56-2.7h3.12l1.56 2.7l-1.56 2.7z"/>`),
		g.Group(children),
	)
}