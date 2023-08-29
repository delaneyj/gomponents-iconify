package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M5 11a1 1 0 1 1 0-2h10a1 1 0 1 1 0 2H5Z"/>`),
		g.Group(children),
	)
}