package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm2 2v14h6v-4h4v-4h4V9H9zm2 2h10v2h-4v4h-4v4h-2V11z"/>`),
		g.Group(children),
	)
}