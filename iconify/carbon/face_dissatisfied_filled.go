package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceDissatisfiedFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm-4.5 9A2.5 2.5 0 1 1 9 13.5a2.48 2.48 0 0 1 2.54-2.5Zm9.64 12.92a6 6 0 0 0-10.28 0l-1.71-1a8 8 0 0 1 13.7 0ZM20.5 16a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}