package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAutofitHeightInTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.78 9.78l2.5-2.5a.75.75 0 0 0-1.06-1.06L17 7.44V3.75a.75.75 0 0 0-1.5 0v3.69l-1.22-1.22a.75.75 0 1 0-1.06 1.06l2.5 2.5a.75.75 0 0 0 1.06 0ZM4 6.25A2.25 2.25 0 0 1 6.25 4h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h4a.75.75 0 0 1 0 1.5h-4A2.25 2.25 0 0 1 4 17.75V6.25Zm15.28 10.47l-2.5-2.5a.75.75 0 0 0-1.06 0l-2.5 2.5a.75.75 0 1 0 1.06 1.06l1.22-1.22v3.69a.75.75 0 0 0 1.5 0v-3.69l1.22 1.22a.75.75 0 1 0 1.06-1.06Z"/>`),
		g.Group(children),
	)
}