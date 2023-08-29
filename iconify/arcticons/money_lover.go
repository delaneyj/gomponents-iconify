package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyLover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8" ry="8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.952a8 8 0 0 0 8 8h21a8 8 0 0 0 8-8"/><circle cx="14.523" cy="31.774" r="1.804" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.477" cy="31.774" r="1.804" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="23.836" cy="35.957" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.686" ry="4.52"/><circle cx="21.184" cy="35.797" r=".967" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.841" cy="35.797" r=".967" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.883 21.913H33.48m-3.439 0a7.536 7.536 0 1 0-11.72 0h11.72Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.623 11.39H24.41c-.765 1.727-2.147 3.082-3.842 3.743v.462h2.018v6.319h3.037V11.388Z"/>`),
		g.Group(children),
	)
}