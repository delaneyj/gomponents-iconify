package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42 18.74V7.5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v11.24m0 10.52V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2V29.26M5 24h37"/>`),
		g.Group(children),
	)
}