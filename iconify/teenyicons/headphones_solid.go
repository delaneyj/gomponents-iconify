package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2 6.5a5.5 5.5 0 1 1 11 0V9h-.5a1.5 1.5 0 0 0-1.5 1.5v3a1.5 1.5 0 0 0 1.5 1.5h1a1.5 1.5 0 0 0 1.5-1.5v-3a1.5 1.5 0 0 0-1-1.415V6.5a6.5 6.5 0 1 0-13 0v2.585A1.5 1.5 0 0 0 0 10.5v3A1.5 1.5 0 0 0 1.5 15h1A1.5 1.5 0 0 0 4 13.5v-3A1.5 1.5 0 0 0 2.5 9H2V6.5Z"/>`),
		g.Group(children),
	)
}