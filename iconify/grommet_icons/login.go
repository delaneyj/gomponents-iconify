package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Login(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M9 15v7h13V2H9v7m9 3H0m13-5l5 5l-5 5"/>`),
		g.Group(children),
	)
}