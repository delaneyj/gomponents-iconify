package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RocketEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9 1C7.488 1 5.408 2.146 4.049 4H3c-.801 0-1.184.367-1.5 1L1 6h2l1 1l1 1v2l1-.5c.633-.316 1-.699 1-1.5V6.951C8.854 5.592 10 3.512 10 2V1H9zM7.5 3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zM2.75 7.25l-.25.25C2 8 2 9 2 9s.945.055 1.5-.5l.25-.25l-1-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}