package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiplierOneXTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 7.75a.75.75 0 0 0-1.39-.39l-.006.009a8.265 8.265 0 0 1-.598.8c-.402.473-.966 1.033-1.642 1.438a.75.75 0 1 0 .772 1.286A7.655 7.655 0 0 0 8.5 9.833v6.417a.75.75 0 0 0 1.5 0v-8.5Zm3.28 4.47a.75.75 0 0 0-1.06 1.06l1.22 1.22l-1.22 1.22a.75.75 0 1 0 1.06 1.06l1.22-1.22l1.22 1.22a.75.75 0 0 0 1.06-1.06l-1.22-1.22l1.22-1.22a.75.75 0 1 0-1.06-1.06l-1.22 1.22l-1.22-1.22Z"/>`),
		g.Group(children),
	)
}