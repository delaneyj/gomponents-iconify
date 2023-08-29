package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 108.4a20 20 0 0 0 16-19.59V64a20 20 0 0 0-20-20H32a20 20 0 0 0-20 20v24.81a20 20 0 0 0 16 19.59a20 20 0 0 1 0 39.2a20 20 0 0 0-16 19.59V192a20 20 0 0 0 20 20h192a20 20 0 0 0 20-20v-24.81a20 20 0 0 0-16-19.59a20 20 0 0 1 0-39.2ZM36 170.34a44 44 0 0 0 0-84.68V68h48v120H36Zm184 0V188H108V68h112v17.66a44 44 0 0 0 0 84.68Z"/>`),
		g.Group(children),
	)
}