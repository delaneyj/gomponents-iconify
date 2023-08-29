package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterMatrix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="9.5" cy="9.5" r="2.5" fill="currentColor"/><circle cx="9.5" cy="22.5" r="2.5" fill="currentColor"/><circle cx="22.5" cy="22.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M28 2H4a2.002 2.002 0 0 0-2 2v24a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2Zm0 13H17V4h11ZM15 4v11H4V4ZM4 17h11v11H4Zm13 11V17h11v11Z"/>`),
		g.Group(children),
	)
}