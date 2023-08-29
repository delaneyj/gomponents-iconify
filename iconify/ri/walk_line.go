package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalkLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.617 8.712l3.205-2.328A1.995 1.995 0 0 1 12.065 6a2.616 2.616 0 0 1 2.427 1.82c.186.583.356.977.51 1.181A4.992 4.992 0 0 0 19 11v2a6.986 6.986 0 0 1-5.401-2.547l-.698 3.956l2.061 1.729l2.223 6.108l-1.88.684l-2.039-5.604l-3.39-2.845a2 2 0 0 1-.714-1.904l.509-2.885l-.677.492l-2.127 2.928l-1.618-1.176L7.6 8.7l.017.012ZM13.5 5.5a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-2.97 13.181l-3.214 3.83l-1.532-1.285l2.975-3.546l.746-2.18l1.791 1.5l-.766 1.681Z"/>`),
		g.Group(children),
	)
}