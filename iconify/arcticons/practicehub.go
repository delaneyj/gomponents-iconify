package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Practicehub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.64 4.5a4.28 4.28 0 0 1 2.36.72c5.3 3.67 10.6 7.31 12.06 11.58s-.23 9.15-3.27 14c5.07-12.68-9.29-15.6-9.29-15.6v20.48c0 4.3-3.91 7.82-9 7.82S8.4 40 8.4 35.71s4.06-7.81 9.13-7.81a13.06 13.06 0 0 1 3.41.45V8.4a3.67 3.67 0 0 1 3.69-3.9Z"/>`),
		g.Group(children),
	)
}