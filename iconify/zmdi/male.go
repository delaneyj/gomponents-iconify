package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Male(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 3q18 0 30.5 12.5t12.5 30t-12.5 30T192 88t-30.5-12.5t-12.5-30t12.5-30T192 3zm192 149H256v277h-43V301h-42v128h-43V152H0v-43h384v43z"/>`),
		g.Group(children),
	)
}