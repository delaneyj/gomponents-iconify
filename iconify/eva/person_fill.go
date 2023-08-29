package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPersonFill0"><g id="evaPersonFill1"><path id="evaPersonFill2" fill="currentColor" d="M12 11a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm6 10a1 1 0 0 0 1-1a7 7 0 0 0-14 0a1 1 0 0 0 1 1Z"/></g></g>`),
		g.Group(children),
	)
}