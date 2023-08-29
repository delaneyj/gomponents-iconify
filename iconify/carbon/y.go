package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Y(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M20 9h-2l-2 7l-2-7h-2l3 9v5h2v-5l3-9z" fill="currentColor"/>`),
		g.Group(children),
	)
}