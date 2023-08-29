package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M235 299h64l-86 85l-85-85h64V149h43v150zM85 0l86 85h-64v150H64V85H0z"/>`),
		g.Group(children),
	)
}