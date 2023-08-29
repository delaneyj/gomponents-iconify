package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downstair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.969.047h-2.896v2h-2.011v2h-2.01v2H8.043v2H6.031v2h-2.01v2H1v2.918h2.896L17 2.965L16.969.047zM2.029 5.971l4.924.001l-2.154-1.414L7.98 1.377L6.654.051l-3.18 3.181l-1.448-2.205l.003 4.944z"/>`),
		g.Group(children),
	)
}