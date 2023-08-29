package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindyDust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 28a5.006 5.006 0 0 1-5-5h2a3 3 0 1 0 3-3h-1v-2h1a5 5 0 0 1 0 10zm-7-10h4v2h-4zm-6 0h4v2h-4zm-6 0h4v2H4zm17-3h-1v-2h1a3 3 0 1 0-3-3h-2a5 5 0 1 1 5 5zm-7-2h4v2h-4zm-6 0h4v2H8z"/>`),
		g.Group(children),
	)
}