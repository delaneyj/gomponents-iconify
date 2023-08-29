package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M40.139 12.871a1.25 1.25 0 0 1-.01 1.768l-22.25 22a1.25 1.25 0 0 1-1.75.007l-9.25-9a1.25 1.25 0 1 1 1.743-1.792L16.993 34l21.378-21.138a1.25 1.25 0 0 1 1.768.01Z"/>`),
		g.Group(children),
	)
}