package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InColumns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h4v12.998H1zm5 0h4v12.973H6zm5 0h4v13.019h-4z"/>`),
		g.Group(children),
	)
}