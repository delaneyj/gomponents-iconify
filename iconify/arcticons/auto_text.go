package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.138 15.762h7.39m-4.516 8.504h12.563M7.138 32.769h14.78m0-.005L30.67 24l-8.753-8.764V5.828c0-1.181 1.429-1.773 2.264-.936L43.266 24L24.182 43.108c-.835.837-2.264.245-2.264-.936v-9.408Z"/>`),
		g.Group(children),
	)
}