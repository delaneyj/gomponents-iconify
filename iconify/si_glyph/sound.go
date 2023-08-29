package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.987 14.868c0 .626-.679 1.132-1.516 1.132l-6.514-4.527c-.839 0-1.913-.508-1.913-1.133V5.682c0-.624 1.074-1.132 1.913-1.132L12.471.022c.837 0 1.516.508 1.516 1.133v13.713z"/>`),
		g.Group(children),
	)
}