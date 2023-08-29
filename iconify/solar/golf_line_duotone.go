package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><ellipse cx="12" cy="18.5" opacity=".5" rx="10" ry="3.5"/><path stroke-linecap="round" d="M12 18V2m0 1.5l5.422 2.711c1.561.78 2.342 1.171 2.342 1.789c0 .618-.78 1.008-2.342 1.789L12 12.5"/></g>`),
		g.Group(children),
	)
}