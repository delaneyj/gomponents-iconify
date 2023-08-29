package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amongus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.53 21.13a5.48 5.48 0 0 1-4.47 4.48c-3.95.85-9.1.68-12.7-1.29c-5-2.71-3.13-10 2.9-10.69c4.59-.48 7.77-.76 11.51.46c1.7.55 3.93 1.22 2.76 7.04Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.26 13.63a21 21 0 0 1 .86-3.55c1-2 2.4-4.42 8.73-5.25a50.43 50.43 0 0 1 6.28.17c3.89.3 5.79 2.64 7.2 7.2A52.18 52.18 0 0 1 35 19.75l.09 23.48H6.88c-1.21-5.1-.54-13.38.48-18.91M35 19.75a53.53 53.53 0 0 1 6.43 1.44c.56.28 1.14.59 1.17 1.08c.58 7.34 1.25 15.92.61 21h-8.09"/>`),
		g.Group(children),
	)
}