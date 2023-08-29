package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M15 0H0v1h15V0ZM7 3H3v12h4V3Zm5 0H8v8h4V3Z"/>`),
		g.Group(children),
	)
}