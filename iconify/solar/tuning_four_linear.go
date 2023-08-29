package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningFourLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="2" transform="rotate(-90 12 12)"/><circle cx="10" cy="20" r="2" transform="rotate(-90 10 20)"/><circle cx="2" cy="2" r="2" transform="matrix(0 -1 -1 0 16 6)"/><path stroke-linecap="round" d="M16 12h3m-5 8h5M10 4H5m0 8h3m-3 8h1M19 4h-1"/></g>`),
		g.Group(children),
	)
}