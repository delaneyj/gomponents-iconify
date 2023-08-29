package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13 0H2v14.5a.5.5 0 0 0 .812.39L7.5 11.14l4.688 3.75A.5.5 0 0 0 13 14.5V0Z"/>`),
		g.Group(children),
	)
}