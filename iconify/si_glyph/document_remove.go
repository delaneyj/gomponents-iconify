package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.969 16V6.012H10V.062L4 0v14l4-.031V16h7.969z"/><path d="M11 0v4.969h4.917L11 0zM2 15v.937h4.99V15H2z"/></g>`),
		g.Group(children),
	)
}