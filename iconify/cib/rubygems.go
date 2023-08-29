package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rubygems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m10.411 10.531l-3.958 3.938l9.589 9.573l3.943-3.938l5.63-5.635l-3.948-3.948v-.016H10.401zM16 0L2.042 8v16L16 32l13.958-8V8zm11.292 22.469L16 28.974L4.708 22.469V9.495L16 2.985l11.292 6.51z"/>`),
		g.Group(children),
	)
}