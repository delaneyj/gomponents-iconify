package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.866 3l9.237 16H23v2H1v-2h.896l9.238-16a1 1 0 0 1 1.732 0ZM12 12.925L8.659 19h6.682L12 12.925Z"/>`),
		g.Group(children),
	)
}