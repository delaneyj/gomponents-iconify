package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowAltLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10.813 9.281L4.093 16l6.72 6.719l1.406-1.438L7.938 17H28v-2H7.937l4.282-4.281z"/>`),
		g.Group(children),
	)
}