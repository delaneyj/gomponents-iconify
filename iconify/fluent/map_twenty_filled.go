package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 6.04l-4-2.8v10.72l4 2.8V6.04Zm1 10.72l4.787-3.35A.5.5 0 0 0 18 13V3.5a.5.5 0 0 0-.787-.41L13 6.04v10.72ZM2.213 6.59L7 3.24v10.72l-4.213 2.95A.5.5 0 0 1 2 16.5V7a.5.5 0 0 1 .213-.41Z"/>`),
		g.Group(children),
	)
}