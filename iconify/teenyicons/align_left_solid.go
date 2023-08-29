package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 0v15h1V0H0Zm11 3H3v4h8V3Zm4 5H3v4h12V8Z"/>`),
		g.Group(children),
	)
}