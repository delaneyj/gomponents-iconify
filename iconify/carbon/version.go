package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Version(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2v2h10v15h2V4a2.002 2.002 0 0 0-2-2Z"/><path fill="currentColor" d="M11 7v2h10v15h2V9a2.002 2.002 0 0 0-2-2Z"/><path fill="currentColor" d="M6 12h10a2.002 2.002 0 0 1 2 2v14a2.002 2.002 0 0 1-2 2H6a2.002 2.002 0 0 1-2-2V14a2.002 2.002 0 0 1 2-2Zm10 2l-10-.001V28h10Z"/>`),
		g.Group(children),
	)
}