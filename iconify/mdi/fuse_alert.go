package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FuseAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7v10h7V7H6m3.16 9v-3.13H7.41L9.91 8v3.14h1.68L9.16 16M14 2v4H5V2c0-.55.45-1 1-1h7c.55 0 1 .45 1 1m0 16v4c0 .55-.45 1-1 1H6c-.55 0-1-.45-1-1v-4h9m5-5h-2V7h2v6m0 4h-2v-2h2v2Z"/>`),
		g.Group(children),
	)
}