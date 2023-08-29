package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisabledComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M23 5.99805H9C7.34315 5.99805 6 7.34119 6 8.99805V30.998C6 32.6549 7.34315 33.998 9 33.998H39C40.6569 33.998 42 32.6549 42 30.998V23.998"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 34V42"/><circle cx="36" cy="12" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M32 8L40 16"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 41.998L34 41.998"/></g>`),
		g.Group(children),
	)
}