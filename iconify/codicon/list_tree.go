package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M4 9h9v1H4zm0 3h7v1H4zm0-6h10v1H4zM1 3h11v1H1z"/><path d="M4 4h1v9H4z"/></g>`),
		g.Group(children),
	)
}