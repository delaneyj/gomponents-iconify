package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v5.5a2 2 0 0 1-2 2H9V20H7V4Zm2 7.5h6V6H9v5.5Z"/>`),
		g.Group(children),
	)
}