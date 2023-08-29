package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorPlayThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m162.22 116.67l-48-32A4 4 0 0 0 108 88v64a4 4 0 0 0 2.11 3.53a4 4 0 0 0 4.11-.2l48-32a4 4 0 0 0 0-6.66ZM116 144.53V95.47L152.79 120ZM208 44H48a20 20 0 0 0-20 20v112a20 20 0 0 0 20 20h160a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20Zm12 132a12 12 0 0 1-12 12H48a12 12 0 0 1-12-12V64a12 12 0 0 1 12-12h160a12 12 0 0 1 12 12Zm-56 48a4 4 0 0 1-4 4H96a4 4 0 0 1 0-8h64a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}