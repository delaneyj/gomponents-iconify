package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormRefresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M17.333 9.333C16.398 7.36 14.358 6 12 6a6 6 0 1 0 6 6m.5-6v4h-4"/>`),
		g.Group(children),
	)
}