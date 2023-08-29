package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitterX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M1 2h2.5L3.5 2h-2.5z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M1 2h2.5L3.5 2h-2.5z;M1 2h2.5L18.5 22h-2.5z"/></path><path d="M5.5 2h2.5L7.2 2h-2.5z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M5.5 2h2.5L7.2 2h-2.5z;M5.5 2h2.5L23 22h-2.5z"/></path><path d="M3 2h5v0h-5z" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.4s" values="M3 2h5v0h-5z;M3 2h5v2h-5z"/></path><path d="M16 22h5v0h-5z" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/><animate fill="freeze" attributeName="d" begin="0.4s" dur="0.4s" values="M16 22h5v0h-5z;M16 22h5v-2h-5z"/></path><path d="M18.5 2h3.5L22 2h-3.5z" opacity="0"><set attributeName="opacity" begin="0.5s" to="1"/><animate fill="freeze" attributeName="d" begin="0.5s" dur="0.4s" values="M18.5 2h3.5L22 2h-3.5z;M18.5 2h3.5L5 22h-3.5z"/></path></g>`),
		g.Group(children),
	)
}