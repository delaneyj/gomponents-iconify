package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Betteruntis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.75 30.85V7.44c0-1.6-.84-2.88-1.89-2.88H7.14c-1 0-1.89 1.28-1.89 2.88v33.12c0 1.59.84 2.88 1.89 2.88h24.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 28.6a8.46 8.46 0 1 0 8.46 8.46h0A8.45 8.45 0 0 0 37 28.6Zm-.05 8.45v-6.03m3.4 9.36l-3.4-3.33M24 4.56v38.88M5.25 17.5h37.5m-37.5 13h26.41"/>`),
		g.Group(children),
	)
}