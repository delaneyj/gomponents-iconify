package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10 0v4h4l-4-4z"/><path d="M9 5V.042H3.033v11.02L4 10.258v1.659L6 12v3.042H4v.937h9.98V5H9z"/><path d="m2.9 12.004l-.003 1h2.044v.975H2.892l-.004.979l-1.846-1.452L2.9 12.004z"/></g>`),
		g.Group(children),
	)
}