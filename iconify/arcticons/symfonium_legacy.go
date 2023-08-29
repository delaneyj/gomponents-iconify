package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymfoniumLegacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.616 19.637A9.695 9.695 0 0 1 12 14.261h0c0-5.391 4.37-9.76 9.76-9.76h5.743c4.172 0 7.17 1.159 9.562 4.272M12.771 39.227c2.392 3.113 5.39 4.273 9.563 4.273h5.774c5.373 0 9.729-4.356 9.729-9.729v-.042a9.684 9.684 0 0 0-1.697-5.492"/><path d="M6.73 22.82h3.354l2.016 3.49l1.898-4.847l2.635 6.309l1.757-7.188l4.073 8.066l2.954-15.732l3.754 22.44l3.035-16.451l2.395 8.066l1.917-6.708l1.515 2.606h3.237"/></g>`),
		g.Group(children),
	)
}