package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14 0v15h1V0h-1Zm-2 3H4v4h8V3Zm0 5H0v4h12V8Z"/>`),
		g.Group(children),
	)
}