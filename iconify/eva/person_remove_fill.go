package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPersonRemoveFill0"><g id="evaPersonRemoveFill1"><path id="evaPersonRemoveFill2" fill="currentColor" d="M21 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Zm-11 5a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1"/></g></g>`),
		g.Group(children),
	)
}