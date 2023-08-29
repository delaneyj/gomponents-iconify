package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M226.4 100.56a12 12 0 0 0 9.6-11.75V64a12 12 0 0 0-12-12H32a12 12 0 0 0-12 12v24.81a12 12 0 0 0 9.6 11.75a28 28 0 0 1 0 54.88a12 12 0 0 0-9.6 11.75V192a12 12 0 0 0 12 12h192a12 12 0 0 0 12-12v-24.81a12 12 0 0 0-9.6-11.75a28 28 0 0 1 0-54.88ZM28 192v-24.81a4 4 0 0 1 3.2-3.91a36 36 0 0 0 0-70.56a4 4 0 0 1-3.2-3.91V64a4 4 0 0 1 4-4h60v136H32a4 4 0 0 1-4-4Zm168-64a36.09 36.09 0 0 0 28.8 35.28a4 4 0 0 1 3.2 3.91V192a4 4 0 0 1-4 4H100V60h124a4 4 0 0 1 4 4v24.81a4 4 0 0 1-3.2 3.91A36.09 36.09 0 0 0 196 128Z"/>`),
		g.Group(children),
	)
}