package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivideLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 128a6 6 0 0 1-6 6H40a6 6 0 0 1 0-12h176a6 6 0 0 1 6 6Zm-94-50a14 14 0 1 0-14-14a14 14 0 0 0 14 14Zm0 100a14 14 0 1 0 14 14a14 14 0 0 0-14-14Z"/>`),
		g.Group(children),
	)
}