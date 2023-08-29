package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBookFill0"><g id="evaBookFill1"><path id="evaBookFill2" fill="currentColor" d="M19 3H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1ZM7 19a1 1 0 0 1 0-2h11v2Z"/></g></g>`),
		g.Group(children),
	)
}