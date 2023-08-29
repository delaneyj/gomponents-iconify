package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDownThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 3.086L12.914 8.5L11.5 9.914l-3-3V15.5h-2V6.914l-3 3L2.086 8.5L7.5 3.086Zm10 5.414v8.586l3-3l1.414 1.414l-5.414 5.414l-5.414-5.414l1.414-1.414l3 3V8.5h2Z"/>`),
		g.Group(children),
	)
}