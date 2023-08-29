package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderVerticalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1 1.5h1m8 0h1m2 0h1m-1 3h1m-1 3h1m-4 0h1m2 3h1m-1 3h1m-4 0h1m-7 0h1m-4 0h1m-1-3h1m-1-3h1m2 0h1m-4-3h1m2-3h1M7.5 1v13"/>`),
		g.Group(children),
	)
}