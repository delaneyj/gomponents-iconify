package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nclient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.12 32.62v-5.16A3.12 3.12 0 0 0 24 24.34h0a3.12 3.12 0 0 0-3.12 3.12v5.16m0-5.16v-3.12M8.81 15.38c3.72 0 6.89.3 9.57 3.24c0 0-3.3 1.47-3.3 3.64s2.7 3.19 4 3.19v1.84a3.05 3.05 0 0 0-2.4.94s-1.84-2.7-4.31-2.7a10.93 10.93 0 0 0-2.26.26a3.6 3.6 0 0 0 1.7-3.27c0-2.4-3.68-6.79-7.32-5.77a6.67 6.67 0 0 1 4.32-1.37Zm30.38 0c-3.72 0-6.89.3-9.57 3.24c0 0 3.3 1.47 3.3 3.64s-2.7 3.19-4 3.19v1.84a3.05 3.05 0 0 1 2.4.94s1.84-2.7 4.31-2.7a10.93 10.93 0 0 1 2.26.26a3.6 3.6 0 0 1-1.7-3.27c0-2.4 3.68-6.79 7.32-5.77a6.67 6.67 0 0 0-4.32-1.37Z"/>`),
		g.Group(children),
	)
}