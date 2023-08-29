package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArrowCircleLeftOutline0"><g id="evaArrowCircleLeftOutline1"><g id="evaArrowCircleLeftOutline2" fill="currentColor"><path d="M16 11h-5.66l1.25-1.31a1 1 0 0 0-1.45-1.38l-2.86 3a1 1 0 0 0-.09.13a.72.72 0 0 0-.11.19a.88.88 0 0 0-.06.28S7 12 7 12a1 1 0 0 0 .08.38a1 1 0 0 0 .21.32l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L10.41 13H16a1 1 0 0 0 0-2Z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/></g></g></g>`),
		g.Group(children),
	)
}