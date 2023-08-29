package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M15.25 13a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Zm2 2.5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/><path fill-rule="evenodd" d="M17.5 7h-9a5 5 0 0 0-5 5v2a5 5 0 0 0 5 5h9a5 5 0 0 0 5-5v-2a5 5 0 0 0-5-5Zm-12 5a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3v-2Z" clip-rule="evenodd"/><path d="M7 14a1 1 0 1 1 0-2h4a1 1 0 0 1 0 2H7Z"/><path d="M10 15a1 1 0 1 1-2 0v-4a1 1 0 0 1 2 0v4Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}