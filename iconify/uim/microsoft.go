package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microsoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9.503v9.503H2zm10.493 0h9.503v9.503h-9.503zM2 12.497h9.503V22H2zm10.493 0h9.503V22h-9.503z"/>`),
		g.Group(children),
	)
}