package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 76h-20V64a20 20 0 0 0-20-20H40a20 20 0 0 0-20 20v96a20 20 0 0 0 20 20h116v12a20 20 0 0 0 20 20h48a20 20 0 0 0 20-20V96a20 20 0 0 0-20-20ZM40 172a12 12 0 0 1-12-12V64a12 12 0 0 1 12-12h144a12 12 0 0 1 12 12v12h-20a20 20 0 0 0-20 20v76Zm196 20a12 12 0 0 1-12 12h-48a12 12 0 0 1-12-12V96a12 12 0 0 1 12-12h48a12 12 0 0 1 12 12Zm-104 16a4 4 0 0 1-4 4H88a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4Zm80-96a4 4 0 0 1-4 4h-16a4 4 0 0 1 0-8h16a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}