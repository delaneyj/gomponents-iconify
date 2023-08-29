package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrenchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.946 9.72l-2.872-.767l-.77-2.874l3.187-3.232a5.73 5.73 0 0 0-5.847 1.39a5.747 5.747 0 0 0-1.28 6.173l-3.475 3.48l-3.478 3.477a5.745 5.745 0 0 0-6.173 1.277a5.731 5.731 0 0 0-1.39 5.85l3.23-3.19l2.875.77l.77 2.873l-3.24 3.197a5.743 5.743 0 0 0 7.146-7.586l3.464-3.464l3.464-3.463c2.07.83 4.523.407 6.202-1.27a5.732 5.732 0 0 0 1.384-5.877L24.946 9.72z"/>`),
		g.Group(children),
	)
}