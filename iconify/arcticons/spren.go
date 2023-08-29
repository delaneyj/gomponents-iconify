package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spren(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.487 14.54c-.012-2.463 1.828-4.742 4.82-5.97a12.983 12.983 0 0 1 9.643.036c2.97 1.248 4.773 3.54 4.72 6.003"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.294 7.49a3.08 3.08 0 0 1 1.345-2.6a2.395 2.395 0 0 1 2.667.03a3.092 3.092 0 0 1 1.293 2.63M9.797 36.647l28.403-.033m-14.354.016v4.302m9.823-26.769c-2.202 18.492 5.89 14.164 4.531 22.45M14.487 14.54c2.202 18.492-6.048 13.82-4.69 22.107"/><circle cx="23.846" cy="42.216" r="1.284" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}