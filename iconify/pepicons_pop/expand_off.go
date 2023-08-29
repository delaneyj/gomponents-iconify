package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M5.707 15.707a1 1 0 0 1-1.414-1.414l4-4a1 1 0 1 1 1.414 1.414l-4 4Z"/><path d="M5 16a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2H5Z"/><path d="M6 15a1 1 0 1 1-2 0v-4a1 1 0 1 1 2 0v4Zm5.707-5.293a1 1 0 0 1-1.414-1.414l4-4a1 1 0 1 1 1.414 1.414l-4 4Z"/><path d="M16 9a1 1 0 1 1-2 0V5a1 1 0 1 1 2 0v4Z"/><path d="M11 6a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2h-4ZM1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}