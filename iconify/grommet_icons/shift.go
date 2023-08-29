package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 0v24M2 12h10m10 0H12M6 8l-4 4l4 4m12-8l4 4l-4 4"/>`),
		g.Group(children),
	)
}