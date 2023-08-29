package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersW(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4v14h5V4h2v14h5V4h2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4h2Z"/>`),
		g.Group(children),
	)
}