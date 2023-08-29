package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPowerFill0"><g id="evaPowerFill1"><g id="evaPowerFill2" fill="currentColor"><path d="M12 13a1 1 0 0 0 1-1V2a1 1 0 0 0-2 0v10a1 1 0 0 0 1 1Z"/><path d="M16.59 3.11a1 1 0 0 0-.92 1.78a8 8 0 1 1-7.34 0a1 1 0 1 0-.92-1.78a10 10 0 1 0 9.18 0Z"/></g></g></g>`),
		g.Group(children),
	)
}