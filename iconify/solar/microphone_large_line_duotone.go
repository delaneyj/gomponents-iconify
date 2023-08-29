package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneLargeLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M6 8a6 6 0 1 1 12 0v5a6 6 0 0 1-12 0V8Z"/><path stroke-linecap="round" d="M10 6.5s.473-.5 2-.5c1.527 0 2 .5 2 .5m-4 3s.473-.5 2-.5c1.527 0 2 .5 2 .5m7 1.5v2a9 9 0 1 1-18 0v-2" opacity=".5"/></g>`),
		g.Group(children),
	)
}