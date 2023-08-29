package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m12 2l10 5v10l-10 5l-10-5V7l10-5ZM2 7l10 5l10-5m-10 5v10v-10Z"/>`),
		g.Group(children),
	)
}