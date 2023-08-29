package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m21 299l43-86H0V85h128v128l-43 86H21zm171 0l43-86h-64V85h128v128l-43 86h-64z"/>`),
		g.Group(children),
	)
}