package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abdominal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M7.9 5c2.837 3.02 4.256 6.94 4.256 11.765c0 7.235-6.157 12.773-7.187 16.735c-.687 2.641-1.01 5.808-.97 9.5M39.256 5c-2.838 3.02-4.257 6.94-4.257 11.765C35 24 41.156 29.538 42.186 33.5c.687 2.641 1.01 5.808.97 9.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M6.244 30.837c6.525 4.13 12.45 6.195 17.777 6.195c5.326 0 10.959-2.065 16.898-6.195"/><path fill="currentColor" d="M24 31a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/></g>`),
		g.Group(children),
	)
}