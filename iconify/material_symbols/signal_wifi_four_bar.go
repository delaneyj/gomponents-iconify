package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiFourBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9L12 21Z"/>`),
		g.Group(children),
	)
}