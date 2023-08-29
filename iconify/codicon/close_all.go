package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m8.621 8.086l-.707-.707L6.5 8.793L5.086 7.379l-.707.707L5.793 9.5l-1.414 1.414l.707.707L6.5 10.207l1.414 1.414l.707-.707L7.207 9.5l1.414-1.414z"/><path d="m5 3l1-1h7l1 1v7l-1 1h-2v2l-1 1H3l-1-1V6l1-1h2V3zm1 2h4l1 1v4h2V3H6v2zm4 1H3v7h7V6z"/></g>`),
		g.Group(children),
	)
}