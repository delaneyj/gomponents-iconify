package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Java(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 18V9h16v9c0 4-2 5-8 5s-8-1-8-5Zm16-9v3a3 3 0 1 0 3-3h-3Zm-4-3V2M5 6V4m4 2V0"/>`),
		g.Group(children),
	)
}