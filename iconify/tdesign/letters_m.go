package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20V6h-5v14h-2V6H6v14H4V4h14a2 2 0 0 1 2 2v14h-2Z"/>`),
		g.Group(children),
	)
}