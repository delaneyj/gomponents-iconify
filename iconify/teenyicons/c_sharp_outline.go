package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CSharpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 5v5m2-5v5M8 6.5h5m-5 2h5M8 10l-.402.201A2.831 2.831 0 0 1 3.5 7.668v-.336a2.832 2.832 0 0 1 4.098-2.533L8 5m-6.5 5.5v-6l6-3.5l6 3.5v6l-6 3.5l-6-3.5Z"/>`),
		g.Group(children),
	)
}