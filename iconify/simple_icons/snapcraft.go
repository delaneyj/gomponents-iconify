package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapcraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.804 13.367V5.69l5.292 2.362l-5.292 5.315zM3.701 23.514l6.49-12.22l2.847 2.843L3.7 23.514zM0 .486l13.355 4.848v8.484L0 .486zm21.803 4.848H14.11L24 9.748z"/>`),
		g.Group(children),
	)
}