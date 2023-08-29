package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserStarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14v8H4a8 8 0 0 1 8-8Zm6 7.5l-2.939 1.545l.561-3.273l-2.377-2.317l3.286-.477L18 14l1.47 2.977l3.285.478l-2.377 2.318l.56 3.272L18 21.5ZM12 13c-3.315 0-6-2.685-6-6s2.685-6 6-6s6 2.685 6 6s-2.685 6-6 6Z"/>`),
		g.Group(children),
	)
}