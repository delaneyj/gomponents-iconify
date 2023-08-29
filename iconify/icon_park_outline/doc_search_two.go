package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocSearchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 27V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h11m-4-32h14m-14 8h14m-14 8h6"/><path d="M37 37c0 1.38-.56 2.63-1.465 3.535A5 5 0 1 1 37 37Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m39 44l-3.465-3.465m0 0a5 5 0 1 0-7.071-7.07a5 5 0 0 0 7.072 7.07Z"/></g>`),
		g.Group(children),
	)
}