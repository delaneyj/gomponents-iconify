package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.071 13.142L13.414 18.8a2 2 0 0 1-2.828 0l-5.657-5.657A5 5 0 1 1 12 6.072a5 5 0 0 1 7.071 7.07Z"/>`),
		g.Group(children),
	)
}