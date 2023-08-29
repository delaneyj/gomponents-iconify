package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarminNav(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.176V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H15.291"/><circle cx="9.5" cy="9.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.176 8.115a2.678 2.678 0 0 0-2.82-2.678a2.782 2.782 0 0 0-2.532 2.83v2.485a2.679 2.679 0 0 0 2.676 2.68h0a2.679 2.679 0 0 0 2.676-2.68h0H9.5M20.25 35.5v-7.109c0-2.811 1.117-5.508 3.105-7.496L31.75 12.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.837 12.5h7.913v7.913"/>`),
		g.Group(children),
	)
}