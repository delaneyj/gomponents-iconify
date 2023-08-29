package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Efteling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.45 12.181c-6.87 2.754-7.707 6.59-7.707 10.407s2.86 11.139 12.028 11.139a10.589 10.589 0 0 0 9.884-5.91"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.386 18.456h22.647A17.345 17.345 0 0 0 21.437 6.115C12.181 6.115 3.5 13.384 3.5 23.32s9.1 18.565 20.552 18.565S44.5 33.745 44.5 24.244a27.546 27.546 0 0 0-.191-3.134"/>`),
		g.Group(children),
	)
}