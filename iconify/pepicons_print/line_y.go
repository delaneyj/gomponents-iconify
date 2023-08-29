package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11 1.5A1.5 1.5 0 0 0 9.5 3v15a1.5 1.5 0 0 0 3 0V3A1.5 1.5 0 0 0 11 1.5Z" opacity=".2"/><path d="M10 .5a.5.5 0 0 0-.5.5v18a.5.5 0 0 0 1 0V1a.5.5 0 0 0-.5-.5Z"/></g>`),
		g.Group(children),
	)
}