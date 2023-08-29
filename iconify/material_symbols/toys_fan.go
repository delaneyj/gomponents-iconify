package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToysFan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12q0-2.275 1.613-3.888T17.5 6.5q2.275 0 3.888 1.613T23 12H12Zm-5.5 5.5q-2.275 0-3.888-1.613T1 12h11q0 2.275-1.613 3.888T6.5 17.5ZM12 12q-2.275 0-3.888-1.613T6.5 6.5q0-2.275 1.613-3.888T12 1v11Zm0 11V12q2.275 0 3.888 1.613T17.5 17.5q0 2.275-1.613 3.888T12 23Z"/>`),
		g.Group(children),
	)
}