package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseSymbolForBeginner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#24bac5" d="M32 20.8V62l20-18.8V2z"/><path fill="#ffce31" d="M12 2v41.2L32 62V20.8z"/><path fill="none" stroke="#3e4347" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="3" d="M32 20.8L12 2v41.2L32 62l20-18.8V2z"/>`),
		g.Group(children),
	)
}