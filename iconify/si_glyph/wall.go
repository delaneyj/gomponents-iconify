package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3h4v2h-4zm-2 3h3v1.957h-3zm4 0h2v1.969h-2zM0 9h2v2H0zm3 0h4v2H3zm5 0h3v2H8zm4 0h4v2h-4zM5 6h4v1.957H5zM0 6h4v1.957H0zm10 6h3v1.957h-3zm4 0h2v1.969h-2zm-9 0h4v1.957H5zm-5 0h4v1.957H0zm8-9h3v2H8zM3 3h4v2H3zM0 3h2v2H0z"/>`),
		g.Group(children),
	)
}