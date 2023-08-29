package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiFourBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.575q-.2 0-.375-.062T11.3 20.3L.7 9.7q-.3-.3-.288-.712T.725 8.3Q3.05 6.2 5.95 5.1T12 4q3.15 0 6.05 1.1t5.225 3.2q.3.275.313.688T23.3 9.7L12.7 20.3q-.15.15-.325.213t-.375.062Z"/>`),
		g.Group(children),
	)
}