package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Survivalmanual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.48 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5H8.43a2 2 0 0 0-1.95 1.95ZM21 30.78A6.78 6.78 0 1 1 27.78 24A6.79 6.79 0 0 1 21 30.78Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 20.76l-4 1.91a.75.75 0 0 0-.33.33l-1.9 4a.18.18 0 0 0 .09.24a.17.17 0 0 0 .14 0l4-1.91a.83.83 0 0 0 .37-.37l1.91-4a.19.19 0 0 0-.11-.23a.17.17 0 0 0-.17.03Z"/><path fill="currentColor" d="M21 24.75a.75.75 0 1 1 .75-.75a.76.76 0 0 1-.75.75Zm12.28 0A.75.75 0 1 1 34 24a.75.75 0 0 1-.75.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.68 29h1.14a2.71 2.71 0 0 0 2.7-2.7v-4.62a2.71 2.71 0 0 0-2.7-2.68h-1.14V6.45a2 2 0 0 0-2-1.95H10.81v39h24.92a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.28 27a3 3 0 0 1 0-6h4.4v6Z"/>`),
		g.Group(children),
	)
}