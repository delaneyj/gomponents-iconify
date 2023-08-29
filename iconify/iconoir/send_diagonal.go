package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22.152 3.553L11.178 21.004l-1.67-8.596L2 7.898l20.152-4.345ZM9.456 12.444l12.696-8.89"/>`),
		g.Group(children),
	)
}