package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBenin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FCD116" d="M32 5H14v13h22V9a4 4 0 0 0-4-4z"/><path fill="#E8112D" d="M14 31h18a4 4 0 0 0 4-4v-9H14v13z"/><path fill="#008751" d="M14 5H4a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h10V5z"/>`),
		g.Group(children),
	)
}