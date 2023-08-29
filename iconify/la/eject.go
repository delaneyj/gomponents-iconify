package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 6.594l-.719.687l-9 9L4.594 18h22.812l-1.687-1.719l-9-9zm0 2.843L22.563 16H9.438zM4 22v2h24v-2z"/>`),
		g.Group(children),
	)
}