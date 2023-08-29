package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluentTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.915 2.082a.75.75 0 0 1 .67 0l6 3a.75.75 0 0 1 0 1.341l-4.658 2.33l4.658 2.329a.75.75 0 0 1 0 1.341L13 15.216v6.034a.75.75 0 0 1-1.149.635l-6-3.77a.75.75 0 0 1-.351-.634V5.752a.75.75 0 0 1 .415-.671l6-3ZM7 6.216v10.85l4.5 2.827v-5.14a.75.75 0 0 1 .415-.671l4.658-2.33l-4.658-2.329a.75.75 0 0 1 0-1.341l4.658-2.33l-4.323-2.16L7 6.215Z"/>`),
		g.Group(children),
	)
}