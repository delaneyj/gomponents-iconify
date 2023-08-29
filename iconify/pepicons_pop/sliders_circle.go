package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlidersCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.5 7.75a1 1 0 0 1 1-1H13a1 1 0 1 1 0 2H6.5a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H19.5a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z"/><path d="M15.75 8.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Zm-10.25 7a1 1 0 0 1 1-1H13a1 1 0 1 1 0 2H6.5a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H19.5a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z"/><path d="M15.75 18.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Zm-10.25-8a1 1 0 0 1 1-1h1.625a1 1 0 0 1 0 2H6.5a1 1 0 0 1-1-1Zm6.5 0a1 1 0 0 1 1-1h6.5a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1Z"/><path d="M10.75 13.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}