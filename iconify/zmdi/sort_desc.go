package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDesc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M282 76H135l74-73zM135 420h147l-73 73zm131-137h81l-40-111zm24-158h34l93 246h-38l-19-53H254l-20 53h-38zM52 337h132v34H0v-28l128-183H1v-35h179v27z"/>`),
		g.Group(children),
	)
}