package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calctastic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.409 36.409h27.182V43.5H10.409zm27.182-21.864H27.545V4.5h-7.09v10.045H10.409v7.091h10.046v10.046h7.09V21.636h10.046v-7.091z"/>`),
		g.Group(children),
	)
}