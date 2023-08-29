package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15v4H5v-4h14m1-2H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1M7 18.5A1.5 1.5 0 0 1 5.5 17A1.5 1.5 0 0 1 7 15.5A1.5 1.5 0 0 1 8.5 17A1.5 1.5 0 0 1 7 18.5M19 5v4H5V5h14m1-2H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1M7 8.5A1.5 1.5 0 0 1 5.5 7A1.5 1.5 0 0 1 7 5.5A1.5 1.5 0 0 1 8.5 7A1.5 1.5 0 0 1 7 8.5Z"/>`),
		g.Group(children),
	)
}