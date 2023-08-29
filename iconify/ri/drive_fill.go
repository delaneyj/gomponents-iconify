package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.94 4.146l3.482 6.03l-5.94 10.293L2 14.44L7.94 4.146Zm2.176 10.294H22l-3.482 6.029H6.635l3.481-6.029Zm4.343-1L8.518 3.145h6.964l5.94 10.295H14.46Z"/>`),
		g.Group(children),
	)
}