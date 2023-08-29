package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarthBoxMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 17v2h-8v-2h8M5 3h14a2 2 0 0 1 2 2v7.8c-.61-.35-1.28-.6-2-.72V5h-3.22c-.11 1-.95 1.79-1.98 1.79h-2v2c0 .56-.45 1-1 1h-2v2h6c.43 0 .8.27.94.65A6.013 6.013 0 0 0 12 17.83c-1.3-.04-2.2-.93-2.2-2.04v-1L5 10.29V19h7.08c.12.72.37 1.39.72 2H5a2 2 0 0 1-2-2V5c0-1.11.89-2 2-2Z"/>`),
		g.Group(children),
	)
}