package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.953 12.938H3.046V3.022h9.907v9.916zm-8.99-.876h8.135V3.833H3.963v8.229zM12 0h.959v1.943H12zM3 0h.959v1.984H3zm9 14h.938v1.922H12zm-9 0h.959v1.984H3zm11-2h1.875v.994H14zM0 12h1.855v.918H0zm14-9h1.98v.938H14zM0 3h1.938v.957H0z"/>`),
		g.Group(children),
	)
}