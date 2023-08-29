package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttentionSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 8L6.629 1.31A1 1 0 0 0 5.905 1H3.732a1 1 0 0 0-.866 1.5L6.04 8l-3.175 5.5a1 1 0 0 0 .866 1.5h2.173a1 1 0 0 0 .724-.31L13 8Z"/>`),
		g.Group(children),
	)
}