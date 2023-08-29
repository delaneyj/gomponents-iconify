package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 1.5a.5.5 0 0 0-1 0V3H1.5a.5.5 0 0 0 0 1H3v8H1.5a.5.5 0 0 0 0 1H3v1.5a.5.5 0 0 0 1 0V13h8v1.5a.5.5 0 0 0 1 0V13h1.5a.5.5 0 0 0 0-1H13V4h1.5a.5.5 0 0 0 0-1H13V1.5a.5.5 0 0 0-1 0V3H4V1.5ZM12 12H4V4h8v8Z"/>`),
		g.Group(children),
	)
}