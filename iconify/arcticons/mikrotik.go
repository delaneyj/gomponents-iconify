package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mikrotik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.878 6.587a21.415 21.415 0 0 0-.11 2.165a21.173 21.173 0 0 0 21.174 21.173a21.31 21.31 0 0 0 3.558-.297M5.65 9.446a29.347 29.347 0 0 0-.15 2.965a29.002 29.002 0 0 0 29.002 29.002a29.2 29.2 0 0 0 4.874-.408"/>`),
		g.Group(children),
	)
}