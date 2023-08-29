package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteBalanceIridescent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.96 19.95l1.8-1.8l-1.42-1.41l-1.79 1.79m0-14.07l1.79 1.8l1.42-1.42l-1.8-1.79m15.49 15.48l-1.79-1.79l-1.42 1.41l1.8 1.8M13 22.45V19.5h-2v2.95h2m6.04-19.4l-1.8 1.79l1.42 1.42l1.79-1.8M11 3.5h2V.55h-2M5 14.5h14v-6H5v6Z"/>`),
		g.Group(children),
	)
}