package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.197 2.698A3.5 3.5 0 0 1 5.198 1h4.604a3.5 3.5 0 0 1 3 1.698L15 6.358V12.5a1.5 1.5 0 0 1-1.5 1.5h-1a1.5 1.5 0 0 1-1.5-1.5V12H4v.5A1.5 1.5 0 0 1 2.5 14h-1A1.5 1.5 0 0 1 0 12.5V6.358l2.197-3.66ZM12 7H3V6h9v1ZM2 10h3V9H2v1Zm11-1h-3v1h3V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}