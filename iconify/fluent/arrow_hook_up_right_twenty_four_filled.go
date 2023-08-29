package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowHookUpRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 18H16a1 1 0 1 1 0 2h-5.5a6.5 6.5 0 1 1 0-13h5.14l-1.933-1.933a1 1 0 0 1 1.414-1.414l3.53 3.529a.997.997 0 0 1 .21.308a.997.997 0 0 1-.18 1.243a.94.94 0 0 1-.028.029l-3.531 3.53a1 1 0 0 1-1.415-1.413L15.586 9H10.5a4.5 4.5 0 1 0 0 9Z"/>`),
		g.Group(children),
	)
}