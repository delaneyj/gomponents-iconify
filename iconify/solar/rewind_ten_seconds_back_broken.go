package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindTenSecondsBackBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7.5 10.5l2.5-2v7"/><path d="M12.5 13.75v-3.5a1.75 1.75 0 1 1 3.5 0v3.5a1.75 1.75 0 1 1-3.5 0Z"/><path stroke-linejoin="round" d="M14 4.5L12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10a9.985 9.985 0 0 1-8-3.999M8 2.832A10.017 10.017 0 0 0 5 4.86A9.97 9.97 0 0 0 2 12c0 .685.069 1.354.2 2"/></g>`),
		g.Group(children),
	)
}