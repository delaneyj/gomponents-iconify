package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerShutdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 1v8M6.994 4.52a9.044 9.044 0 0 0-1.358 1.116a9 9 0 1 0 11.37-1.117"/>`),
		g.Group(children),
	)
}