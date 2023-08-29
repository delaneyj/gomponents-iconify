package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudArrowUpTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.08 8.02a6.001 6.001 0 0 1 11.84 0a4.5 4.5 0 0 1 4.053 4.973A6.5 6.5 0 0 0 10.018 17H6.5a4.5 4.5 0 0 1-.42-8.982ZM22 16.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Zm-6 3a.5.5 0 0 0 1 0v-4.793l1.646 1.647a.5.5 0 0 0 .708-.708l-2.5-2.5a.5.5 0 0 0-.708 0l-2.5 2.5a.5.5 0 0 0 .708.708L16 14.707V19.5Z"/>`),
		g.Group(children),
	)
}