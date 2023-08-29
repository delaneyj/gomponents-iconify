package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easycontacts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.695 13.774c5.957-.011 8.95 7.188 4.742 11.404c-4.208 4.216-11.413 1.236-11.413-4.72a6.683 6.683 0 0 1 6.67-6.684h.001ZM10.341 40.543v-6.434c0-2.477 5.887-4.493 13.354-4.493s13.354 2.078 13.354 4.555v6.372"/>`),
		g.Group(children),
	)
}