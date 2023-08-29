package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autodark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.827 42.077H26.7V43.5h-5.124v-1.423h-4.413V29.978a14.075 14.075 0 0 1-6.12-15.657C12.608 7.916 18.017 4.5 23.995 4.5a13.413 13.413 0 0 1 12.953 9.821c1.566 5.836-.897 11.956-6.12 15.657Z"/>`),
		g.Group(children),
	)
}