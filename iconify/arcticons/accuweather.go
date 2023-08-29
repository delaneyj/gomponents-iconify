package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accuweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 11.045a12.954 12.954 0 0 0-9.16 22.114A12.954 12.954 0 1 0 33.16 14.84A13.11 13.11 0 0 0 24 11.045Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.648 7.45L24.03 2.5l4.35 5.007l6.848-2.632v7.029l6.87.732l-1.634 6.998l5.037 4.33l-5.002 4.19l1.581 7.024l-6.852.884v6.88l-6.884-2.567L24.04 45.5l-4.446-5.128l-6.763 2.579v-6.883l-6.934-.902L7.56 28.16L2.5 23.957l5.088-4.29l-1.616-7.006l6.847-.822V4.946l6.83 2.503"/>`),
		g.Group(children),
	)
}