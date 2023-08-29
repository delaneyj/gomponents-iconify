package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterI(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M15.675 9.156c0-1.55.992-2.418 2.326-2.418c1.333 0 2.325.868 2.325 2.418v17.611c0 1.551-.992 2.418-2.325 2.418c-1.333 0-2.326-.867-2.326-2.418V9.156z"/>`),
		g.Group(children),
	)
}