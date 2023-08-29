package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m222.29 123.06l-88-112a8 8 0 0 0-12.58 0l-88 112a8 8 0 0 0 0 9.88l88 112a8 8 0 0 0 12.58 0l88-112a8 8 0 0 0 0-9.88ZM136 39.13l67.42 85.8L136 155.58Zm-16 116.45l-67.42-30.65L120 39.13Zm0 17.57v43.72l-53.43-68Zm16 0l53.43-24.29l-53.43 68Z"/>`),
		g.Group(children),
	)
}