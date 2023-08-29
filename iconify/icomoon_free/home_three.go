package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m16 9.5l-3-3V2h-2v2.5l-3-3l-8 8v.5h2v5h5v-3h2v3h5v-5h2z"/>`),
		g.Group(children),
	)
}