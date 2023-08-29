package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlueCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M57.55 14.449c9.697 14.11 6.111 33.406-8 43.1c-14.11 9.697-33.407 6.112-43.1-8c-9.687-14.11-6.111-33.408 8-43.1c14.11-9.687 33.406-6.109 43.1 8"/>`),
		g.Group(children),
	)
}