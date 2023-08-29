package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M14 19h1v-1h-1v1Zm-9 4h14V1H5v22ZM8 5h8h-8Zm0 4h8h-8Zm0 4h8h-8Z"/>`),
		g.Group(children),
	)
}