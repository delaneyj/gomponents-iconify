package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphScatter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M15 13v1H1.5l-.5-.5V0h1v13h13Z"/><path d="M5 2h2v2H5zm7-1h2v2h-2zM8 5h2v2H8zM5 9h2v2H5zm7-1h2v2h-2z"/></g>`),
		g.Group(children),
	)
}