package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M.5 12.5v-10l7 5l-7 5Zm7 0v-10l7 5l-7 5Z"/>`),
		g.Group(children),
	)
}