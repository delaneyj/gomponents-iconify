package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M20 14.12A7.78 7.78 0 0 1 9.88 4a7.782 7.782 0 0 0 2.9 15A7.782 7.782 0 0 0 20 14.12z"/>`),
		g.Group(children),
	)
}