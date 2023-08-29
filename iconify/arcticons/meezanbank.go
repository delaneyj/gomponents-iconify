package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meezanbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.272 40.536L5.5 36.339l4.925-9.291h5.485l4.925 9.291l-1.772 4.197H7.272zm12.499-33.11l-7.576 15.411h6.723l5.28-10.217l5.283 10.217h6.721L28.629 7.426h-8.858zm9.166 33.148l-1.772-4.197l4.925-9.291h5.485l4.925 9.291l-1.772 4.197H28.937z"/>`),
		g.Group(children),
	)
}