package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M9.5 14a3 3 0 1 1 0 6a3 3 0 0 1 0-6Zm5-10a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/><path stroke-linecap="round" d="M11 7H6M3 7H2m11 10h5m3 0h1M2 17h4M22 7h-4"/></g>`),
		g.Group(children),
	)
}