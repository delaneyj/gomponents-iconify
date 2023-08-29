package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneLargeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M18 8v5a6 6 0 0 1-12 0V8a6 6 0 0 1 10.472-4"/><path d="M10 6.5s.473-.5 2-.5c1.527 0 2 .5 2 .5m-4 3s.473-.5 2-.5c1.527 0 2 .5 2 .5m7 1.5v2a9 9 0 0 1-5 8.064M3 11v2a9 9 0 0 0 9 9"/></g>`),
		g.Group(children),
	)
}