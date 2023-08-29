package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M25.5 7A2.5 2.5 0 0 0 23 9.5V15H13V9.5a2.5 2.5 0 1 0-5 0v17a2.5 2.5 0 1 0 5 0V20h10v6.5a2.5 2.5 0 1 0 5 0v-17A2.5 2.5 0 0 0 25.5 7z"/>`),
		g.Group(children),
	)
}