package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersZ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.571 5.1a1.1 1.1 0 0 0-1.1-1.1H7v2h7.571v1.135l-7.571 9V18.9A1.1 1.1 0 0 0 8.1 20H17v-2H9v-1.135l7.571-9V5.1Z"/>`),
		g.Group(children),
	)
}