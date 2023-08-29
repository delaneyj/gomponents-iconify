package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutMenubar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m1 2l1-1h12l1 1v12l-1 1H2l-1-1V2Zm1 0v12h12V2H2Zm1 1h2v1H3V3Zm3 0h2v1H6V3Zm5 0H9v1h2V3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}