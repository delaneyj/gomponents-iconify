package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emerald(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.11 5.5H14.89L5.5 14.88v18.23l9.39 9.39h18.22l9.39-9.39V14.88Z"/>`),
		g.Group(children),
	)
}