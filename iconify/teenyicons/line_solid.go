package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 .647 2.854l10 10a1.5 1.5 0 1 0 .707-.707l-10-10A1.5 1.5 0 0 0 1.5 0Z"/>`),
		g.Group(children),
	)
}