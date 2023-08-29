package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 128a4 4 0 0 1-4 4H40a4 4 0 0 1 0-8h176a4 4 0 0 1 4 4ZM40 68h176a4 4 0 0 0 0-8H40a4 4 0 0 0 0 8Zm176 120H40a4 4 0 0 0 0 8h176a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}