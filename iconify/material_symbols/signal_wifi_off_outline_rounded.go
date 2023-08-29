package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalWifiOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.925 15.075L16.5 13.65l4.6-4.6q-1.975-1.5-4.3-2.275T12 6q-.725 0-1.45.075t-1.45.2l-1.65-1.65q1.125-.3 2.25-.463T12 4q3.125 0 5.988 1.087T23.2 8.3q.325.275.338.713t-.288.737l-5.325 5.325ZM12 18.15l1.675-1.65l-8.75-8.75q-.525.275-1.025.613t-1 .687l9.1 9.1Zm7.775 4.425L15.1 17.9l-2.4 2.4q-.3.3-.7.3t-.7-.3L.75 9.75q-.3-.3-.3-.75t.325-.725Q1.425 7.7 2.1 7.2t1.375-.9l-2.1-2.1q-.3-.3-.288-.7t.313-.7q.3-.3.713-.3t.712.3L21.2 21.175q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM12.8 9.95Zm-3.5 2.175Z"/>`),
		g.Group(children),
	)
}