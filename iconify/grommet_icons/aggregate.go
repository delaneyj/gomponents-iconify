package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aggregate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 15h7V8a7 7 0 1 0-7 7Zm8-6H9v7a7 7 0 1 0 7-7Z"/>`),
		g.Group(children),
	)
}