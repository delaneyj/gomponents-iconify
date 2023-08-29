package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerForceShutdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm0-13v8"/>`),
		g.Group(children),
	)
}