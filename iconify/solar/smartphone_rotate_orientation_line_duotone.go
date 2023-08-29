package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneRotateOrientationLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 8c0-2.828 0-4.243.879-5.121C3.757 2 5.172 2 8 2h1c2.828 0 4.243 0 5.121.879C15 3.757 15 5.172 15 8v8c0 2.828 0 4.243-.879 5.121C13.243 22 11.828 22 9 22H8c-2.828 0-4.243 0-5.121-.879C2 20.243 2 18.828 2 16V8Z"/><path stroke-linecap="round" d="M17.5 10.007c1.86.027 2.914.165 3.621.872C22 11.758 22 13.172 22 16c0 2.829 0 4.243-.879 5.122c-.707.707-1.761.845-3.621.872"/><path stroke-linecap="round" d="M11 5H6m13 13v-4" opacity=".5"/><path d="M10 17.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="M20.5 6.986L22 8c0-3.015-2.162-5.517-5-6" opacity=".5"/></g>`),
		g.Group(children),
	)
}