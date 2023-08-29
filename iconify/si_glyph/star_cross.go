package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.601 6.4L8 0L6.398 6.4L0 8l6.398 1.601L8 16l1.601-6.399L16 8L9.601 6.4zM8 9.334a1.333 1.333 0 1 1 .003-2.667A1.333 1.333 0 0 1 8 9.334z"/>`),
		g.Group(children),
	)
}