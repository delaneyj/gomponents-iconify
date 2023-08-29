package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bereal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.837 27.771V20.23h2.47c1.395 0 2.527 1.134 2.527 2.533s-1.132 2.533-2.528 2.533h-2.47m2.47-.001l2.469 2.475m-7.018-.95a1.885 1.885 0 0 1-1.639.951h0a1.886 1.886 0 0 1-1.885-1.885V24.66c0-1.042.844-1.886 1.885-1.886h0c1.042 0 1.886.845 1.886 1.886v.613h-3.771M30.82 26.82a1.885 1.885 0 0 1-1.638.951h0a1.886 1.886 0 0 1-1.886-1.885V24.66c0-1.042.844-1.886 1.886-1.886h0c1.041 0 1.885.845 1.885 1.886v.613h-3.771m10.978-5.044v6.6c0 .52.422.943.943.943h.283m-2.951-1.886a1.886 1.886 0 0 1-1.886 1.885h0a1.886 1.886 0 0 1-1.885-1.885V24.66c0-1.042.844-1.886 1.885-1.886h0c1.042 0 1.886.845 1.886 1.886m0 3.111v-4.997M11.611 24a1.886 1.886 0 1 1 0 3.771H8.5v-7.543h3.111a1.886 1.886 0 1 1 0 3.772h0Zm0 0H8.5"/>`),
		g.Group(children),
	)
}