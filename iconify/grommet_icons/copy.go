package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M9 15h8h-8Zm0-4h10H9Zm0-4h4h-4Zm7-6v6h6M6 5H2v18h16v-4m4 0H6V1h11l5 5v13Z"/>`),
		g.Group(children),
	)
}