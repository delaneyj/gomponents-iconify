package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoxCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 21l-2.75-3l1.16-1.16L17 18.43l3.59-3.59l1.16 1.41M12.8 21H5c-1.11 0-2-.89-2-2V5c0-1.11.89-2 2-2h14c1.11 0 2 .89 2 2v7.8c-.61-.35-1.28-.6-2-.72V5H5v14h7.08c.12.72.37 1.39.72 2m-.8-4H7v-2h5m2.68-2H7v-2h10v1.08c-.85.14-1.63.46-2.32.92M17 9H7V7h10"/>`),
		g.Group(children),
	)
}