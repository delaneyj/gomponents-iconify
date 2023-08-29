package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridFourUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1h1V0H0zm2 0v1h1V0H2zm2 0v1h1V0H4zm2 0v1h1V0H6zM0 2v1h1V2H0zm2 0v1h1V2H2zm2 0v1h1V2H4zm2 0v1h1V2H6zM0 4v1h1V4H0zm2 0v1h1V4H2zm2 0v1h1V4H4zm2 0v1h1V4H6zM0 6v1h1V6H0zm2 0v1h1V6H2zm2 0v1h1V6H4zm2 0v1h1V6H6z"/>`),
		g.Group(children),
	)
}