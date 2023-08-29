package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Espn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.044 5.825L8.746 16.283l32.456-.072L42.5 5.825H10.044ZM8.421 19.78L5.5 42.176h32.456l.974-9.736H18.807l.65-4.22h20.122l1.298-8.438H8.421Z"/>`),
		g.Group(children),
	)
}