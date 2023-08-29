package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FallingBlocks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.658 35.634l13.523 6.774m9.831-17.936l-18.493-8.23a2 2 0 0 1-1.014-2.64L26.111 5.5M5.628 38.403L21.1 35.116a2 2 0 0 0 1.54-2.372l-5.778-27.2"/>`),
		g.Group(children),
	)
}