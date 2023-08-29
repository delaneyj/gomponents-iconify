package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 21V7h5l2-4h8l2 4h5v14H1Zm11-3a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/>`),
		g.Group(children),
	)
}