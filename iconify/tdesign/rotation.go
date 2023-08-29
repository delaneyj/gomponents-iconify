package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v9a9 9 0 0 1 9 9h9v2H2V2h2Zm0 18h7a7 7 0 0 0-7-7v7Z"/>`),
		g.Group(children),
	)
}