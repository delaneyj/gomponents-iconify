package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileXSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm8.146 8.354L7.5 8.207L5.854 9.854l-.708-.708L6.793 7.5L5.146 5.854l.708-.708L7.5 6.793l1.646-1.647l.708.708L8.207 7.5l1.647 1.646l-.708.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}