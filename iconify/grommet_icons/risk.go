package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Risk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M14 10h1V9h-1v1Zm4 0h1V9h-1v1Zm0-4h1V5h-1v1Zm-4 0h1V5h-1v1ZM9 19h1v-1H9v1Zm-4-4h1v-1H5v1Zm5-5H1v13h13v-9m-4 0h13V1H10v13Z"/>`),
		g.Group(children),
	)
}