package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentLess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 6h10v2H2zm3 6h7v2H5zm-3 6h10v2H2zm3 6h7v2H5zM16 4h2v24h-2zm12.15 19.5l1.41-1.38L23.27 16l6.29-6.12l-1.41-1.38l-7.71 7.5l7.71 7.5z"/>`),
		g.Group(children),
	)
}