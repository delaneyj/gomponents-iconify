package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m0 192l111 141q15 15 30 2q14-16 2-30l-77-92h297q21 0 21-21t-21-21H66l77-94q6-7 5-16t-7-14q-17-12-30 4z"/>`),
		g.Group(children),
	)
}