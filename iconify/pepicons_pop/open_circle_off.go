package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M13.707 13.707a1 1 0 0 1-1.414-1.414l6-6a1 1 0 1 1 1.414 1.414l-6 6Z"/><path d="M18 18v-3.5a1 1 0 1 1 2 0V19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h4.5a1 1 0 0 1 0 2H8v10h10Zm2-7a1 1 0 1 1-2 0V7a1 1 0 1 1 2 0v4Z"/><path d="M15 8a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2h-4ZM4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}