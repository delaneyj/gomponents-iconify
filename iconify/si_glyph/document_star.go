package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 .058L2.008 0v15.954H14V5H7.977C7.964 5 8 .058 8 .058zm1.052 9.165h3.056l-2.471 1.806l.943 2.924l-2.471-1.808l-2.473 1.808l.943-2.924l-2.471-1.806h3.056l.943-3.27l.945 3.27z"/><path d="M9.024.016v3.979h4.947L9.024.016z"/></g>`),
		g.Group(children),
	)
}