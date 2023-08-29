package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braille(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm10 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6zM7 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm10 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2zM7 20a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm10-10a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/>`),
		g.Group(children),
	)
}