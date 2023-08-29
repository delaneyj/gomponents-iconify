package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.556 3.444a4 4 0 0 1 0 5.656l-8.2 8.2a4 4 0 0 1-2.386 1.148l-3.38.374l-2.297 2.3a1 1 0 0 1-1.414-1.415l2.298-2.299l.375-3.378A4 4 0 0 1 6.7 11.645l8.2-8.2a4 4 0 0 1 5.658 0Zm-9.192 9.192L9.95 14.05l2.121 2.122l1.414-1.415l-2.12-2.121Zm2.829-2.828l-1.415 1.414l2.122 2.121l1.414-1.414l-2.121-2.121Zm2.828-2.829l-1.414 1.414l2.121 2.122L19.142 9.1l-2.121-2.12Z"/>`),
		g.Group(children),
	)
}