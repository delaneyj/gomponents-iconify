package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeliciousCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.916 9.021c0-4.369-3.555-7.91-7.937-7.91c-4.385 0-7.938 3.541-7.938 7.91c0 4.368 3.554 7.909 7.938 7.909c4.382.001 7.937-3.54 7.937-7.909zm-8.013.063H2.976c0-3.339 2.734-6.044 6.107-6.044l.025.001v5.865h6.081c0 3.329-2.926 6.204-6.287 6.219V9.084h.001z"/>`),
		g.Group(children),
	)
}