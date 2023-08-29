package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiBadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12.075ZM16.4 22L15 20.6l2.1-2.1l-2.1-2.1l1.4-1.4l2.1 2.1l2.1-2.1l1.4 1.4l-2.075 2.1L22 20.6L20.6 22l-2.1-2.075L16.4 22Zm-3.7-1.7l-.7.7L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9l-3.7 3.7q-.425-.125-.875-.2t-.925-.075q-.2 0-.413.013t-.412.037L21.1 9.05q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l9.1 9.1l.475-.475q-.025.2-.038.412t-.012.413q0 .475.075.925t.2.875Z"/>`),
		g.Group(children),
	)
}