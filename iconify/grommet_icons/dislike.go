package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dislike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M23 1H4C2 1 1 2 1 4v10h7v6c0 2 1 3 3 3h2s.016-6 .016-7.326C13.016 14.348 14 13 16 13h7V1Zm-5 0v12"/>`),
		g.Group(children),
	)
}