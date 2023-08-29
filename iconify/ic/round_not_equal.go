package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundNotEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9.998H6a1 1 0 1 1 0-2h12a1 1 0 0 1 0 2zm0 6H6a1 1 0 0 1 0-2h12a1 1 0 0 1 0 2z"/><path fill="currentColor" d="M14.999 5H15c.507.219.742.806.525 1.314l-5.212 12.162A.999.999 0 0 1 9 19a1.002 1.002 0 0 1-.525-1.314l5.212-12.162A.999.999 0 0 1 15 5z"/>`),
		g.Group(children),
	)
}