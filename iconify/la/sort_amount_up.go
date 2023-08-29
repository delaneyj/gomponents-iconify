package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAmountUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 5v2h12V5H4zm18 .5l-.72.69L17 10.5l1.41 1.41L21 9.31V28h2V9.31l2.59 2.6L27 10.5l-4.28-4.31L22 5.5zM4 9v2h10V9H4zm0 4v2h8v-2H4zm0 4v2h6v-2H4zm0 4v2h4v-2H4zm0 4v2h2v-2H4z"/>`),
		g.Group(children),
	)
}