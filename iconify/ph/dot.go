package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M140 128a12 12 0 1 1-12-12a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}