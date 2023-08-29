package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyYen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.271 5H22l-6 11l-6-11H7.729l6.065 11H8v2h7v3H8v2h7v4h2v-4h7v-2h-7v-3h7v-2h-5.794l6.065-11z"/>`),
		g.Group(children),
	)
}