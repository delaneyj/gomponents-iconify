package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 5.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM2 14a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" clip-rule="evenodd"/><path d="M16.707 15.293a1 1 0 0 1-1.414 1.414l-8.5-8.5a1 1 0 0 1 1.414-1.414l8.5 8.5Z"/><path d="M6.793 11.293a1 1 0 1 0 1.414 1.414l8.5-8.5a1 1 0 0 0-1.414-1.414l-8.5 8.5Z"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}