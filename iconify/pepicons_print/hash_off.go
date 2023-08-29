package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M4 8.75a1.5 1.5 0 0 1 1.5-1.5h11a1.5 1.5 0 0 1 0 3h-11A1.5 1.5 0 0 1 4 8.75Zm0 5.5a1.5 1.5 0 0 1 1.5-1.5h11a1.5 1.5 0 0 1 0 3h-11a1.5 1.5 0 0 1-1.5-1.5Z"/><path d="M8.875 3.505a1.5 1.5 0 0 1 1.37 1.62l-1 12a1.5 1.5 0 1 1-2.99-.25l1-12a1.5 1.5 0 0 1 1.62-1.37Zm5.499 0a1.5 1.5 0 0 1 1.37 1.62l-1 12a1.5 1.5 0 1 1-2.989-.25l1-12a1.5 1.5 0 0 1 1.62-1.37Z"/></g><path fill-rule="evenodd" d="M3 7.75a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1-.5-.5Zm0 5.5a.5.5 0 0 1 .5-.5h13a.5.5 0 1 1 0 1h-13a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.791 3.502a.5.5 0 0 1 .457.54l-1 12a.5.5 0 0 1-.996-.084l1-12a.5.5 0 0 1 .54-.456Zm5.5 0a.5.5 0 0 1 .457.54l-1 12a.5.5 0 0 1-.996-.084l1-12a.5.5 0 0 1 .54-.456Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}