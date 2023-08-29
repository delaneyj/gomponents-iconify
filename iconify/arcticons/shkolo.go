package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shkolo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 35H38a5.5 5.5 0 0 0 5.5-5.5h0A5.5 5.5 0 0 0 38 24H10a5.5 5.5 0 0 1-5.5-5.5h0A5.5 5.5 0 0 1 10 13h32.5"/>`),
		g.Group(children),
	)
}