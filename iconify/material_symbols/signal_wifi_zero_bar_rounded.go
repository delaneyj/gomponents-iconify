package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiZeroBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.575q-.2 0-.375-.062T11.3 20.3L.7 9.7q-.3-.3-.288-.712T.725 8.3Q3.05 6.175 5.95 5.087T12 4q3.15 0 6.05 1.088T23.275 8.3q.3.275.313.688T23.3 9.7L12.7 20.3q-.15.15-.325.213t-.375.062Zm0-2.425l9.1-9.1q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l9.1 9.1Z"/>`),
		g.Group(children),
	)
}