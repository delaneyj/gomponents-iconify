package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoverFlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4v8h7.894L13 4H5zm9 0v8h.916V4H14zm2-2h1v12h-1zM3 4h1v8H3zM1 2h.979v11.918H1z"/>`),
		g.Group(children),
	)
}