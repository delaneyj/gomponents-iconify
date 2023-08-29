package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirbudsLeftLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-width="1.5" d="M2 18.667v.833a2.5 2.5 0 0 0 5 0v-.833m-5 0V7.559A5.588 5.588 0 0 1 7.56 2h.104A3.353 3.353 0 0 1 11 5.336V8a3 3 0 0 1-3 3a1 1 0 0 0-1 1v6.667m-5 0h5"/><path stroke-linecap="round" stroke-width="1.676" d="M8 5v3"/><circle cx="5.5" cy="5.5" r="5.5" stroke-width="1.5" transform="matrix(-1 0 0 1 21 11)"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 14v5h3"/><path stroke-linecap="round" stroke-width="1.5" d="M14 5.1A5.006 5.006 0 0 1 17.9 9"/></g>`),
		g.Group(children),
	)
}