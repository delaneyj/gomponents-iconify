package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M5 10.5a.5.5 0 0 1 0-1h10a.5.5 0 0 1 0 1H5Z"/>`),
		g.Group(children),
	)
}