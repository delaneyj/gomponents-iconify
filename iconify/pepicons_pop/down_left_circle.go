package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.707 19.707a1 1 0 0 1-1.414 0l-4-4a1 1 0 0 1 0-1.414l4-4a1 1 0 0 1 1.414 1.414L7.414 15l3.293 3.293a1 1 0 0 1 0 1.414Z"/><path d="M15.75 14c.595 0 1.166-.238 1.588-.663a2.284 2.284 0 0 0 .662-1.61V6a1 1 0 1 1 2 0v5.727a4.285 4.285 0 0 1-1.242 3.02A4.239 4.239 0 0 1 15.75 16H6a1 1 0 1 1 0-2h9.75Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}