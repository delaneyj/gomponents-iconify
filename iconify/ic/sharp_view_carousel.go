package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpViewCarousel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7h4v10H2V7zm5 12h10V5H7v14zM18 7h4v10h-4V7z"/>`),
		g.Group(children),
	)
}