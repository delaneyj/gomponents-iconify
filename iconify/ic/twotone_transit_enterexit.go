package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTransitEnterexit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.98 6L9 12.77V8H6v10h10v-3h-4.85L18 8.03z"/>`),
		g.Group(children),
	)
}