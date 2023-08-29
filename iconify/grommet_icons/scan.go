package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M20 10V3H4v7m-3 2h22H1Zm3 1v3v-3Zm16 3v-3v3ZM7 21H4v-3m16 0v3h-3m-8 0h6h-6Z"/>`),
		g.Group(children),
	)
}