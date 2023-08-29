package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiZeroBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9L12 21Zm0-2.85l9.1-9.1q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l9.1 9.1Z"/>`),
		g.Group(children),
	)
}