package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gotube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.118 44.435C-1.401 46.386-.426 3.5 17.118 3.5c11.696 0 27.29 12.67 27.29 18.518c0 6.823-15.595 22.418-27.29 22.418Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.97 21.044a4.828 4.828 0 1 0-9.655 0v4.919a4.828 4.828 0 0 0 9.656 0h-4.828m13.645 4.827a4.843 4.843 0 0 1-4.843-4.843v-4.934a4.843 4.843 0 1 1 9.686 0v4.934a4.843 4.843 0 0 1-4.843 4.843Z"/>`),
		g.Group(children),
	)
}