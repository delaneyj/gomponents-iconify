package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseSymbolForBeginner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="m17 12.007l-1 36l19 16v-40l-18-12z"/><path fill="#5c9e31" d="m55 12.007l1 36l-19 16v-40l18-12z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="2" d="M17.677 12.097c-.922-.6-1.677-.19-1.677.91v34a4.735 4.735 0 0 0 1.56 3.252l16.819 13.496a2.624 2.624 0 0 0 3.123.005l16.934-13.506A4.715 4.715 0 0 0 56 47.007v-34c0-1.1-.755-1.51-1.677-.91l-16.646 10.82a3.352 3.352 0 0 1-3.354 0ZM36 24.007v40"/>`),
		g.Group(children),
	)
}