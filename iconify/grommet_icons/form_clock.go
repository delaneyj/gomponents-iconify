package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-10v4l3 1"/>`),
		g.Group(children),
	)
}