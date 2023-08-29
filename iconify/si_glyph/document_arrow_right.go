package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10 4h4l-4-4v4z"/><path d="M9 5.042v-5H3v11.403l3.583 3.138l-1.971 1.396H14V5.042H9z"/><path d="m3.084 13.004l.003 1H1.043v.975h2.05l.004.979l1.845-1.452l-1.858-1.502z"/></g>`),
		g.Group(children),
	)
}