package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 0h.918v16H2zm12.905 0H4v16h10.905c.604 0 1.095-.453 1.095-1.01V1.01C16 .452 15.51 0 14.905 0zm-4.899 2.875c1.095 0 1.985.886 1.985 1.979s-.891 3.184-1.985 3.184c-1.096 0-1.984-2.09-1.984-3.184a1.98 1.98 0 0 1 1.984-1.979zM14.976 12H5.002s-.171-2.987 2.387-2.987c.537.577 1.272.931 2.625.931c1.354 0 1.999-.36 2.529-.944c2.878 0 2.433 3 2.433 3z"/>`),
		g.Group(children),
	)
}