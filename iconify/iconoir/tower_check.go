package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13 19l3 3l5-5M9 22H7a2 2 0 0 1-2-2v-8.818a.6.6 0 0 0-.1-.333L3.1 8.15a.6.6 0 0 1-.1-.333V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h2.8a.6.6 0 0 1 .6.6v1.8a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V2.6a.6.6 0 0 1 .6-.6h1.8a.6.6 0 0 1 .6.6v5.218a.6.6 0 0 1-.1.333l-1.8 2.698a.6.6 0 0 0-.1.333V13.5"/>`),
		g.Group(children),
	)
}