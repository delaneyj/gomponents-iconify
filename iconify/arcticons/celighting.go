package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Celighting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.326 8.174a15.826 15.826 0 0 0 0 31.652h19.218A3.957 3.957 0 0 0 43.5 35.87h0a3.957 3.957 0 0 0-3.957-3.957H20.327a7.913 7.913 0 0 1-.106-15.826h19.324A3.957 3.957 0 0 0 43.5 12.13h0a3.957 3.957 0 0 0-3.957-3.957H20.327Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.326 20.044A3.957 3.957 0 0 0 16.37 24h0a3.957 3.957 0 0 0 3.957 3.957h19.216A3.957 3.957 0 0 0 43.5 24h0a3.957 3.957 0 0 0-3.957-3.957H20.327Z"/>`),
		g.Group(children),
	)
}