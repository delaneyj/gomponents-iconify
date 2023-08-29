package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAmountDesc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 320v-43h128v43H0zM0 64h384v43H0V64zm0 149v-42h256v42H0z"/>`),
		g.Group(children),
	)
}