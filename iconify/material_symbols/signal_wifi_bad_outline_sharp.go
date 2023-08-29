package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiBadOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9l-3 3h-9v9Zm3.4 0L14 19.6l2.1-2.1l-2.1-2.1l1.4-1.4l2.1 2.1l2.1-2.1l1.4 1.4l-2.1 2.1l2.1 2.1l-1.4 1.4l-2.1-2.1l-2.1 2.1Z"/>`),
		g.Group(children),
	)
}