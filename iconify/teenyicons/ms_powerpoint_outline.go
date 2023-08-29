package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MsPowerpointOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.755 3.5a7 7 0 1 1 0 8M2.5 10V8.5m0 0v-3H5a1.5 1.5 0 1 1 0 3H2.5Zm-1-5h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}