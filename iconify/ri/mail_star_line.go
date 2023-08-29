package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailStarLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13h-2V7.238l-7.928 7.1L4 7.216V19h10v2H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v9ZM4.511 5l7.55 6.662L19.502 5H4.511ZM19.5 21.75l-2.645 1.39l.505-2.945l-2.14-2.086l2.957-.43L19.5 15l1.323 2.68l2.957.43l-2.14 2.085l.505 2.946L19.5 21.75Z"/>`),
		g.Group(children),
	)
}