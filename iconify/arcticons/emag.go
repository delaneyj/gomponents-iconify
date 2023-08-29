package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.409 27.503a2.315 2.315 0 0 1-2.034 1.197a2.4 2.4 0 0 1-2.392-2.392v-1.555a2.392 2.392 0 1 1 4.785 0v.837H7.983m31.915-3.156a3.228 3.228 0 0 0-3.335-3.215a3.354 3.354 0 0 0-2.977 3.334v2.977a3.215 3.215 0 1 0 6.43 0h-3.214m-22.116 3.258v-9.57l4.785 9.57l4.784-9.57v9.57m8.193-.04l-3.098-9.532l-3.217 9.532m1.079-3.218h4.171"/>`),
		g.Group(children),
	)
}