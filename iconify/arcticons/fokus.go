package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fokus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 23.159l12.346 12.413l19.381-19.488l-3.636-3.656l-15.63 15.715l-8.71-8.756Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.482 31.916l3.636 3.656L43.5 16.084l-3.636-3.656l-3.637 3.656m-15.674 8.359l-5.028-5.056l-3.637 3.656"/>`),
		g.Group(children),
	)
}