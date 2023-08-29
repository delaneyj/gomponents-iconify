package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHeaderFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h2v6h4V4h2v14H9v-6H5v6H3V4m15 14v-5h-5v-2l5-7h2v7h1v2h-1v5h-2m0-7V7.42L15.45 11H18Z"/>`),
		g.Group(children),
	)
}