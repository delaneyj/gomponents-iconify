package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guitartuna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c5.374 0 16.755-18.833 16.755-27.622S33.222 4.5 24 4.5S7.245 7.09 7.245 15.878S18.626 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.863 27.937c1.797-5.036 18.95-15.566 22.16-15.786c1.334-.092 1.334 1.112.785 1.896c-1.92 2.742-7.73 10.208-7.73 15.093c0 5.258 3.374 5.807 5.258 5.807a8.292 8.292 0 0 0 2.048-.354"/>`),
		g.Group(children),
	)
}