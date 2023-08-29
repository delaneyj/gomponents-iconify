package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.79 2.093A.5.5 0 0 0 0 2.5v10a.5.5 0 0 0 .79.407L7 8.472V12.5a.5.5 0 0 0 .79.407l7-5a.5.5 0 0 0 0-.814l-7-5A.5.5 0 0 0 7 2.5v4.028L.79 2.093Z"/>`),
		g.Group(children),
	)
}