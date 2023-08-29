package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugBreakpointConditionalUnverified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.778 4.674a4 4 0 1 1 4.444 6.652a4 4 0 0 1-4.444-6.652zm.694 5.612a2.75 2.75 0 1 0 3.056-4.572a2.75 2.75 0 0 0-3.056 4.572zM9.5 6.5h-3v1h3v-1zm0 2h-3v1h3v-1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}