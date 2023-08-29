package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Measurement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 1h12v22H6V1Zm2 2v3h2.5v2H8v3h4v2H8v3h2.5v2H8v3h8V3H8Z"/>`),
		g.Group(children),
	)
}