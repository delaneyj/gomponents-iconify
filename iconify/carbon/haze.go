package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Haze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 28H3a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2zm-4-8H3a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2zm10.793 2.379l2.828 2.828l-1.414 1.414l-2.828-2.828zM28 15h4v2h-4zm-3.621-6.793l2.828-2.828l1.414 1.414l-2.828 2.828zM17 2h2v4h-2zm-6.793 7.621L7.38 6.793l1.414-1.414l2.828 2.828z"/><path fill="currentColor" d="M18 8a8.01 8.01 0 0 0-8 8h2a6 6 0 1 1 6 6H7a1 1 0 0 0 0 2h11a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}