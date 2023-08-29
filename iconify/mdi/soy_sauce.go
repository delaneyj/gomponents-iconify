package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoySauce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.9 7.5c0-.7.2-1.2.3-1.5h.6l.9-2.5h.8V2h-9v1.5h.8L9.2 6h.6c.2.3.3.8.3 1.5c0 1.3-4.1 6.2-4.1 10.1v2c0 1.4 2.7 2.3 6 2.3s6-.9 6-2.3v-2c0-3.9-4.1-8.8-4.1-10.1M12 15a2 2 0 0 1-2-2a2 2 0 0 1 2-2a2 2 0 0 1 2 2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}