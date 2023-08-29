package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Detach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m4 4l16 16m2-8l-5.28 5.28M15 19l-2 2c-6 6-15-3-9-9l2-2m2-2l5-5c4-4 10 2 6 6l-5 5m-2 2l-2 2c-2 2-5-1-3-3l2-2m2-2l5-5"/>`),
		g.Group(children),
	)
}