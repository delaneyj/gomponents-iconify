package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.707 16.707a1 1 0 0 1-1.414 0l-4-4a1 1 0 0 1 0-1.414l4-4a1 1 0 0 1 1.414 1.414L4.414 12l3.293 3.293a1 1 0 0 1 0 1.414Z"/><path d="M12.75 11c.595 0 1.166-.238 1.588-.663A2.284 2.284 0 0 0 15 8.727V3a1 1 0 1 1 2 0v5.727a4.285 4.285 0 0 1-1.242 3.02A4.239 4.239 0 0 1 12.75 13H3a1 1 0 1 1 0-2h9.75Z"/></g>`),
		g.Group(children),
	)
}