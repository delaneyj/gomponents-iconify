package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m280 260l90-95l-87 21L145 58V2L-1 150l66-3l118 145l-14 78l87-88l204 180z"/>`),
		g.Group(children),
	)
}