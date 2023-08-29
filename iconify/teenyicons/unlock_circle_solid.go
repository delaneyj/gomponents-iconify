package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM7.621 4C6.726 4 6 4.726 6 5.621V6h3.5A1.5 1.5 0 0 1 11 7.5v3A1.5 1.5 0 0 1 9.5 12h-4A1.5 1.5 0 0 1 4 10.5v-3a1.5 1.5 0 0 1 1-1.415v-.464a2.621 2.621 0 0 1 4.475-1.853l.379.378l-.708.708l-.378-.38A1.621 1.621 0 0 0 7.62 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}