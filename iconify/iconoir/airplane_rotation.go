package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1.5" stroke-width="1.5"><path d="M9.879 14.122a3 3 0 1 0 4.242-4.243a3 3 0 0 0-4.242 4.243Z"/><path stroke-width="1.497" d="M4.37 16.773A8.956 8.956 0 0 1 3.002 12c0-4.236 2.934-7.792 6.878-8.747A8.998 8.998 0 0 1 12 3.002m7.715 4.365A8.953 8.953 0 0 1 20.999 12c0 3.806-2.368 7.063-5.709 8.378c-1.02.4-2.13.621-3.29.621"/><path d="M14.121 9.88s-.009-2.803 1.415-4.243c1.41-1.409 2.793-2.865 4.242-1.415c1.377 1.378.015 2.81-1.414 4.243c-1.402 1.406-4.243 1.414-4.243 1.414Zm-4.242 4.24s.009 2.803-1.415 4.243c-1.41 1.409-2.793 2.865-4.242 1.415c-1.377-1.378-.015-2.81 1.414-4.243c1.402-1.406 4.243-1.414 4.243-1.414Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}