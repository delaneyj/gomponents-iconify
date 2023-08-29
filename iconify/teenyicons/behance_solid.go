package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BehanceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h3a3 3 0 0 1 2.051 5.19A3.001 3.001 0 0 1 4 13H0V2Zm1 6v4h3a2 2 0 1 0 0-4H1Zm0-1h2a2 2 0 1 0 0-4H1v4Zm13-3H9V3h5v1ZM8 9.5a3.5 3.5 0 1 1 7 0v.5H9.05a2.5 2.5 0 0 0 4.477.964l.81.586A3.5 3.5 0 0 1 8 9.5ZM9.05 9h4.9a2.5 2.5 0 0 0-4.9 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}