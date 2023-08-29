package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaArrowCircleDownOutline0"><g id="evaArrowCircleDownOutline1"><g id="evaArrowCircleDownOutline2" fill="currentColor"><path d="M14.31 12.41L13 13.66V8a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a.54.54 0 0 0 .16-.1a.49.49 0 0 0 .15-.1l3-2.86a1 1 0 0 0-1.38-1.45Z"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/></g></g></g>`),
		g.Group(children),
	)
}