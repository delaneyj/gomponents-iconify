package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiBad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.4 22L15 20.6l2.1-2.1l-2.1-2.1l1.4-1.4l2.1 2.1l2.1-2.1l1.4 1.4l-2.075 2.1L22 20.6L20.6 22l-2.1-2.075L16.4 22ZM12 21L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9l-3.7 3.7q-.425-.125-.875-.2t-.925-.075q-2.525 0-4.3 1.763T12.425 18.5q0 .475.075.925t.2.875l-.7.7Z"/>`),
		g.Group(children),
	)
}