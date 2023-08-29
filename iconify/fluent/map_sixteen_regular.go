package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.235 2.076a.5.5 0 0 1 .489-.023l4.749 2.374l3.762-2.351A.5.5 0 0 1 15 2.5V11a.5.5 0 0 1-.235.424l-4 2.5a.5.5 0 0 1-.489.023l-4.749-2.374l-3.762 2.351A.5.5 0 0 1 1 13.5V5a.5.5 0 0 1 .235-.424l4-2.5ZM6 10.691l4 2V5.309l-4-2v7.382ZM5 3.402L2 5.277v7.32l3-1.874v-7.32Zm6 1.875v7.32l3-1.874v-7.32l-3 1.874Z"/>`),
		g.Group(children),
	)
}