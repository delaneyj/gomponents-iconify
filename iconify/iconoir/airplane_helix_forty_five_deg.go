package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneHelixFortyFiveDeg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M14.12 14.121A3 3 0 1 0 9.879 9.88a3 3 0 0 0 4.243 4.242Z"/><path d="M9.879 9.879s-2.803.009-4.243-1.415c-1.409-1.41-2.864-2.793-1.414-4.242c1.378-1.377 2.81-.015 4.242 1.414C9.87 7.037 9.88 9.879 9.88 9.879Zm4.242 0s-.009-2.803 1.415-4.243c1.41-1.409 2.793-2.864 4.242-1.414c1.377 1.378.015 2.81-1.414 4.242c-1.402 1.406-4.243 1.415-4.243 1.415Zm-4.242 4.242s.009 2.803-1.415 4.243c-1.41 1.409-2.793 2.864-4.242 1.414c-1.377-1.378-.015-2.81 1.414-4.242c1.401-1.406 4.243-1.415 4.243-1.415Zm4.242 0s2.803-.009 4.243 1.415c1.409 1.41 2.864 2.793 1.414 4.242c-1.378 1.377-2.81.015-4.242-1.414c-1.406-1.402-1.415-4.243-1.415-4.243Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}