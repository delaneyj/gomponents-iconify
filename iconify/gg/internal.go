package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Internal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m20.708 4.412l-10.25 10.287h3.59v2h-7v-7h2v3.58L19.293 3l1.416 1.412Z"/><path d="M11 4.706v2H5v12h12v-6h2v8H3v-16h8Z"/></g>`),
		g.Group(children),
	)
}