package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Localbitcoins(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.2 25.988L24 43.504l11.733-17.7m-10.858-7.288a2.692 2.692 0 0 1 0 5.383h-4.442V13.133h4.442a2.692 2.692 0 1 1 0 5.383Zm0 0h-4.44m2.855 5.438v1.355m0-13.65v1.416"/><path fill="none" stroke="currentColor" stroke-dasharray="3.969 3.969" stroke-linecap="round" stroke-linejoin="round" d="M37.864 18.405a13.9 13.9 0 0 1-13.9 13.9c-18.44-.731-18.435-27.072 0-27.801a13.9 13.9 0 0 1 13.9 13.9Z"/>`),
		g.Group(children),
	)
}