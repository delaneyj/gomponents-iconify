package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentDecreaseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 4H3V3h12v1ZM1.207 7.5l1.647-1.646l-.708-.708l-2 2a.5.5 0 0 0 0 .708l2 2l.708-.708L1.207 7.5ZM15 8H7V7h8v1Zm0 4H3v-1h12v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}