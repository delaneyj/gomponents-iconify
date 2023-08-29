package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m16.818 7.646l-5.94-5.44a.64.64 0 0 0-.849.002l-.005 3.793H2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h8.019l-.006 3.877a.642.642 0 0 0 .849-.003l5.954-5.452a.518.518 0 0 0 .002-.777z"/>`),
		g.Group(children),
	)
}