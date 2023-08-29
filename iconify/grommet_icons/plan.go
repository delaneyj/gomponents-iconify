package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M18 4V0v4ZM7 18H5h2Zm12 0H9h10ZM7 14H5h2Zm12 0H9h10ZM6 4V0v4ZM1 9h22H1Zm0 14h22V4H1v19Z"/>`),
		g.Group(children),
	)
}