package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm40.49-100.49a12 12 0 0 1-17 17L140 117v51a12 12 0 0 1-24 0v-51l-11.51 11.52a12 12 0 0 1-17-17l32-32a12 12 0 0 1 17 0Z"/>`),
		g.Group(children),
	)
}