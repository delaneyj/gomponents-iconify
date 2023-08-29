package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortAsc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M282 76H135l74-73zM135 420h147l-73 73zM70 283h81l-41-111zm23-158h35l93 246h-38l-20-53H57l-19 53H0zm192 212h132v34H233v-28l128-183H234v-35h179v27z"/>`),
		g.Group(children),
	)
}