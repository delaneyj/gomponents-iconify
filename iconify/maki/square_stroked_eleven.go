package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareStrokedEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.19 1H1.81a.81.81 0 0 0-.81.81v7.38c0 .447.363.81.81.81h7.38a.81.81 0 0 0 .81-.81V1.81A.81.81 0 0 0 9.19 1zM2 2h7v7H2V2z" fill="currentColor"/>`),
		g.Group(children),
	)
}