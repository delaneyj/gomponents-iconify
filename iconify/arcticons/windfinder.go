package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windfinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.13 36.411c8.188-1.906 15.337-2.327 26.047-2.804c-.953-9.869 1.094-20.691 2.664-24.616C22.318 16 11.551 27.83 8.131 36.41Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.026 25.161c-1.556-1.409-5.09-1.43-6.897 1.598m6.897-1.598c-.506 2.037 1.145 5.16 4.668 5.323m-4.668-5.323c2.034-.517 4.006-3.448 2.485-6.63"/>`),
		g.Group(children),
	)
}