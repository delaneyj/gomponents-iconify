package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pawn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 16h12l-3-4V5l2-2V0h-1.938v1H9.937V0H8.062v1H5.937V0H4v3l2 2v7l-3 4zm7-11h1v7h-1V5z"/>`),
		g.Group(children),
	)
}