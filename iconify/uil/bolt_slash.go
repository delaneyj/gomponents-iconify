package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoltSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.67 4.23A1 1 0 0 0 9.6 4h5.1l-1.27 4.74a1 1 0 0 0 .17.87a1 1 0 0 0 .79.39H18l-1.13 1.24a1 1 0 0 0 .07 1.41a1 1 0 0 0 .67.26a1 1 0 0 0 .74-.33L21 9.67A1 1 0 0 0 20.23 8h-4.54L17 3.26a1 1 0 0 0-.18-.87A1 1 0 0 0 16 2H9a1 1 0 0 0-1 .74V3a1 1 0 0 0 .67 1.23Zm13 16.06l-18-18a1 1 0 0 0-1.38 1.42L6.61 8l-1.26 4.74a1 1 0 0 0 .18.87a1 1 0 0 0 .79.39h3.84l-1.81 6.74a1 1 0 0 0 .49 1.14a1 1 0 0 0 .48.12a1 1 0 0 0 .74-.33l4.85-5.34l5.38 5.38a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM7.62 12l.63-2.34L10.59 12Zm3.73 5.28l1-3.56l1.2 1.19Z"/>`),
		g.Group(children),
	)
}