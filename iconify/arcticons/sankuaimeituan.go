package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sankuaimeituan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.895 13.28H39.79M9.618 20.6h27.768M8.622 26.445h33.48M7.426 33.021h34.942M24.561 13.254l-.194 19.832c-4.662 3.073-6.224 9.432-18.867 7.708"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.366 33.086c2.506 2.686 6.654 9.371 18.134 7.077M34.463 6.915c-5.471 5.57-5.65 6.386-9.902 6.338M15.53 6.98c3.61 4.942 6.097 6.24 9.03 6.273"/>`),
		g.Group(children),
	)
}