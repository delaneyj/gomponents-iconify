package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 16V7m0 4c0-3 0-6-7-6c0 3 0 6 7 6Zm-8 5h16m-2 0l-2 7H8l-2-7m6-9c0-3 0-6 7-6c0 3 0 6-7 6Z"/>`),
		g.Group(children),
	)
}