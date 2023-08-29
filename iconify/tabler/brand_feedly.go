package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandFeedly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7.833 12.278l4.445-4.445M10.055 14.5l2.223-2.222m0 4.444l.555-.555m6.995-1.339a4 4 0 0 0 0-5.656l-5-5a4 4 0 0 0-5.656 0l-5 5a4 4 0 0 0 0 5.656L10.343 21h3.314l6.171-6.172z"/>`),
		g.Group(children),
	)
}