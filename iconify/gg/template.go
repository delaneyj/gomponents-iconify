package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Template(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 3v6h18V3H3Zm16 2H5v2h14V5ZM3 11v10h8V11H3Zm6 2H5v6h4v-6Z" clip-rule="evenodd"/><path d="M21 11h-8v2h8v-2Zm-8 4h8v2h-8v-2Zm8 4h-8v2h8v-2Z"/></g>`),
		g.Group(children),
	)
}