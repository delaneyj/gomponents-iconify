package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brackets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3.5h5v2H4v13h3v2H2v-17Zm15 0h5v17h-5v-2h3v-13h-3v-2Z"/>`),
		g.Group(children),
	)
}