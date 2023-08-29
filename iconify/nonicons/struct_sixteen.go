package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StructSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2Zm2-.286A.286.286 0 0 0 1.714 2v4c0 .158.128.286.286.286h12A.286.286 0 0 0 14.286 6V2A.286.286 0 0 0 14 1.714H2Zm-2 9.429a2 2 0 0 1 2-2h2.857a2 2 0 0 1 2 2V14a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.857Zm2-.286a.286.286 0 0 0-.286.286V14c0 .158.128.286.286.286h2.857A.286.286 0 0 0 5.143 14v-2.857a.286.286 0 0 0-.286-.286H2Zm7.143.286a2 2 0 0 1 2-2H14a2 2 0 0 1 2 2V14a2 2 0 0 1-2 2h-2.857a2 2 0 0 1-2-2v-2.857Zm2-.286a.286.286 0 0 0-.286.286V14c0 .158.128.286.286.286H14a.286.286 0 0 0 .286-.286v-2.857a.286.286 0 0 0-.286-.286h-2.857Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}