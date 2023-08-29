package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendarcolor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.38a13 13 0 0 0-13 13v0a12.93 12.93 0 0 1 13 0a12.93 12.93 0 0 1 13 0a13 13 0 0 0-13-13ZM11 18.39a13 13 0 1 0 13 22.48a13 13 0 0 1-6.5-11.23v0A13 13 0 0 1 11 18.39Zm26 0a13 13 0 0 1-6.5 11.24h0A13 13 0 0 1 24 40.86A13 13 0 1 0 37 18.4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.39a13 13 0 0 0-6.49 11.24a12.93 12.93 0 0 0 13 0A13 13 0 0 0 24 18.39Z"/>`),
		g.Group(children),
	)
}