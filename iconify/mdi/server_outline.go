package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4.6v4.8c0 .9.5 1.6 1.2 1.6h17.7c.6 0 1.2-.7 1.2-1.6V4.6C22 3.7 21.5 3 20.8 3H3.2C2.5 3 2 3.7 2 4.6M10 8V6H9v2h1M5 8h2V6H5v2m15 1H4V5h16v4M2 14.6v4.8c0 .9.5 1.6 1.2 1.6h17.7c.6 0 1.2-.7 1.2-1.6v-4.8c0-.9-.5-1.6-1.2-1.6H3.2c-.7 0-1.2.7-1.2 1.6m8 3.4v-2H9v2h1m-5 0h2v-2H5v2m15 1H4v-4h16v4Z"/>`),
		g.Group(children),
	)
}