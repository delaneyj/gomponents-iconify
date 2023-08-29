package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tentenklooni(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l17.92 9.75v19.5L24 43.5L6.08 33.75v-19.5ZM6.08 14.25L24 24m17.92-9.75L24 24m0 19.5V24"/>`),
		g.Group(children),
	)
}