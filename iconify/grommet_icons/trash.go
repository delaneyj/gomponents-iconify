package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4 5h16v18H4V5ZM1 5h22M9 1h6v4H9V1Zm0 0h6v4H9V1Zm6 8v10M9 9v10"/>`),
		g.Group(children),
	)
}