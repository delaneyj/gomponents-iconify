package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squeegee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2v3H2V2h20M2 8h5l2 2h1v10a2 2 0 0 0 2 2a2 2 0 0 0 2-2V10h1l2-2h5V6H2v2Z"/>`),
		g.Group(children),
	)
}