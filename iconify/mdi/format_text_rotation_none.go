package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextRotationNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.5 18l-3 3v-2H5v-2h12.5v-2l3 3m-10.37-8h3.75L12 4.97L10.13 10m2.62-7l4.75 11h-2.08l-.92-2.19h-5L8.58 14H6.5l4.75-11h1.5Z"/>`),
		g.Group(children),
	)
}