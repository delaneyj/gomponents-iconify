package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Issues(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M7.5 1a6.5 6.5 0 1 0 0 13a6.5 6.5 0 0 0 0-13Zm0 12a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11Z"/><circle cx="7.5" cy="7.5" r="1"/></g>`),
		g.Group(children),
	)
}