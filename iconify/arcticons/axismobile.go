package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axismobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 39.835h11.39l12.616-22.143L24 8.165Zm29.695-12.479H24l7.11 12.479H42.5Z"/>`),
		g.Group(children),
	)
}