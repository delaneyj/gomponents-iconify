package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VimSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 1H1v3h1v10h3.74L14 3.675V1H8v3h1.432L6 8.119V4h1V1Z"/>`),
		g.Group(children),
	)
}