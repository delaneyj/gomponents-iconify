package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.925 15.075L7.45 4.625q1.125-.3 2.25-.463T12 4q3.125 0 5.988 1.087T23.2 8.3q.325.275.338.713t-.288.737l-5.325 5.325Zm1.85 7.5L15.1 17.9l-2.4 2.4q-.3.3-.7.3t-.7-.3L.75 9.75q-.3-.3-.3-.75t.325-.725Q1.425 7.7 2.1 7.2t1.375-.9l-2.1-2.1q-.3-.3-.288-.7t.313-.7q.3-.3.713-.3t.712.3L21.2 21.175q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3Z"/>`),
		g.Group(children),
	)
}