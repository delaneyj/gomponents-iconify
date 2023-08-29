package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6.416 3.788C8.289 2.44 10.506 2 12 2c3.526 0 5.826 1.492 7.212 3.416C20.56 7.289 21 9.506 21 11v9a1 1 0 0 1-1.707.707L18 19.414L16.414 21a2 2 0 0 1-2.828 0L12 19.414L10.414 21a2 2 0 0 1-2.828 0L6 19.414l-1.293 1.293A1 1 0 0 1 3 20v-9c0-3.526 1.492-5.826 3.416-7.212zm1.168 1.624C6.175 6.426 5 8.126 5 11v6.682A2 2 0 0 1 7.414 18L9 19.586L10.586 18a2 2 0 0 1 2.828 0L15 19.586L16.586 18A2 2 0 0 1 19 17.682V11c0-1.173-.36-2.956-1.412-4.416C16.575 5.175 14.874 4 12 4c-1.173 0-2.956.36-4.416 1.412zM7 10a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm8-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4z"/></g>`),
		g.Group(children),
	)
}