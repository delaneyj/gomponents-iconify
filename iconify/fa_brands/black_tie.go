package fa_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackTie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M0 32v448h448V32H0zm316.5 325.2L224 445.9l-92.5-88.7l64.5-184l-64.5-86.6h184.9L252 173.2l64.5 184z"/>`),
		g.Group(children),
	)
}