package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAmountAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 64h128v43H0V64zm0 256v-43h384v43H0zm0-149h256v42H0v-42z"/>`),
		g.Group(children),
	)
}