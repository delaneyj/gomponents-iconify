package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafeSquareLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z"/><path stroke-linecap="round" d="M6 7v10" opacity=".5"/><path d="M11 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/><path stroke-linecap="round" d="M16.5 9.5L18 8m-8 8l1.5-1.5m0-5L10 8m8 8l-1.5-1.5" opacity=".5"/></g>`),
		g.Group(children),
	)
}