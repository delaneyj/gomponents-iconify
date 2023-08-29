package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 23h18V6H3v17Zm6-9h6v-4H9v4ZM1 6h22V1H1v5Z"/>`),
		g.Group(children),
	)
}