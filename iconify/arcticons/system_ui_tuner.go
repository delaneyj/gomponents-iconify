package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemUiTuner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.75 4.5h-15.5a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h15.5a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2ZM23.91 11l-2.07 2.07a2.93 2.93 0 1 0 4.14 0Zm0 .04v7.08"/>`),
		g.Group(children),
	)
}