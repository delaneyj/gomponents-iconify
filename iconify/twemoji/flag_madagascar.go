package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagMadagascar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FC3D32" d="M32 5H13v13h23V9a4 4 0 0 0-4-4z"/><path fill="#007E3A" d="M13 31h19a4 4 0 0 0 4-4v-9H13v13z"/><path fill="#EEE" d="M13 5H4a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h9V5z"/>`),
		g.Group(children),
	)
}