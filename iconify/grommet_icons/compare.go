package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11 7H1v10h6V8m4-3v4l2-2l-2-2Zm0 12l2 2v-4l-2 2Zm2 0h10V7h-6v9"/>`),
		g.Group(children),
	)
}