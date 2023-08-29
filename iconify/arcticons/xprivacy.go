package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xprivacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.48 15.03L24 23.26l-5.52 8.24m11.04-16.47L24 23.26l5.52 8.24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.2 4.86L6.69 11.25V27C6.69 35.44 24 43.5 24 43.5S41.31 35.44 41.31 27V11.25L25.8 4.86a4.68 4.68 0 0 0-3.6 0Z"/>`),
		g.Group(children),
	)
}