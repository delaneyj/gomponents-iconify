package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassesTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.33 13.497a6.987 6.987 0 0 0 2.67-5.5h2a8.987 8.987 0 0 1-3.662 7.246l2.528 4.378a2 2 0 0 1-.732 2.732l-3.527-6.109a8.97 8.97 0 0 1-3.607.753a8.969 8.969 0 0 1-3.608-.753l-3.526 6.109a2 2 0 0 1-.732-2.732l5.063-8.77a4.002 4.002 0 0 1 1.802-6.728V1.997h2v2.126a4.002 4.002 0 0 1 1.803 6.728l1.528 2.646Zm-1.731 1.001l-1.529-2.646a4.003 4.003 0 0 1-2.141 0L9.4 14.498c.803.322 1.68.499 2.598.499c.919 0 1.796-.177 2.6-.499Zm-2.6-5.501a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}