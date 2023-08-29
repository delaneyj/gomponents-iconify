package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Legend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22h14v2H16z"/><rect width="6" height="6" x="4" y="20" fill="currentColor" rx="1"/><path fill="currentColor" d="M16 8h14v2H16zm-6.5 4h-5a.5.5 0 0 1-.447-.724l2.5-5.022a.52.52 0 0 1 .894 0l2.5 5.023A.5.5 0 0 1 9.5 12z"/>`),
		g.Group(children),
	)
}