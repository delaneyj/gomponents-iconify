package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path d="M28 16c0 1.38-.56 2.63-1.465 3.535A5 5 0 1 1 28 16Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m30 23l-3.465-3.465m0 0a5 5 0 1 0-7.071-7.07a5 5 0 0 0 7.072 7.07ZM17 30h14m-14 6h7"/></g>`),
		g.Group(children),
	)
}