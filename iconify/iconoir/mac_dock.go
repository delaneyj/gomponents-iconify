package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacDock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M8 17a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/><path d="M21 21H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM2 17.5l2-1m18 1l-2-1"/></g>`),
		g.Group(children),
	)
}