package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M11.354 9.354a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Zm-6 6a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M5 15.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1H5Z"/><path d="M5.5 15a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 1 0v4Zm10-6a.5.5 0 0 1-1 0V5a.5.5 0 0 1 1 0v4Z"/><path d="M11 5.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1h-4Z"/></g>`),
		g.Group(children),
	)
}