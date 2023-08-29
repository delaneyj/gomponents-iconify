package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 29H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16v2H5v22h16Z"/><path fill="currentColor" d="M15 9h2v14h-2zm12 0h2v14h-2zm-6 0h2v14h-2z"/>`),
		g.Group(children),
	)
}