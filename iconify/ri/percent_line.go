package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.505 21.003a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0-2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-11-9a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0-2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm12.571-4.485l1.414 1.414L4.934 20.488L3.52 19.074L19.076 3.518Z"/>`),
		g.Group(children),
	)
}