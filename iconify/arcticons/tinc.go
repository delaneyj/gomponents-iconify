package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tinc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.06 4.5l7.55 7.55l-7.55 7.55v-5.66H13.94v8.17h-3.78V10.16h23.9Zm0 21.39h3.78v12h-23.9v5.61L6.39 36l7.55-7.55v5.65h20.12Z"/>`),
		g.Group(children),
	)
}