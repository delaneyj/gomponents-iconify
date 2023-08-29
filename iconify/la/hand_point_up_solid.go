package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 2c-1.645 0-3 1.355-3 3v10.75l-.75-.188c-.156-.203-.223-.332-.625-.624c-.64-.47-1.633-.938-2.969-.938C4.23 14 3 15.29 3 16.906v.407l.281.312L10 24.406V30h16V13.156c0-1.41-.996-2.64-2.375-2.937L16 8.563V5c0-1.645-1.355-3-3-3zm0 2c.566 0 1 .434 1 1v5.188l.781.187l8.438 1.781c.468.102.781.524.781 1V23H11.406l-6.312-6.406c.082-.422.254-.594.562-.594c.903 0 1.461.273 1.813.531c.351.258.437.438.437.438l.188.343l.406.125l2.25.594l1.25.313V5c0-.566.434-1 1-1zm-1 21h12v3H12z"/>`),
		g.Group(children),
	)
}