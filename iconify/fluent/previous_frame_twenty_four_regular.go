package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviousFrameTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.5 3.75a.75.75 0 0 1 1.5 0v16.5a.75.75 0 0 1-1.5 0V3.75Zm-6.239-.44c1.162-.796 2.74.035 2.74 1.443v14.495c0 1.413-1.59 2.244-2.75 1.437L1.753 13.383a1.75 1.75 0 0 1 .01-2.88L12.26 3.309Zm1.24 1.443a.25.25 0 0 0-.392-.206L2.611 11.74a.25.25 0 0 0-.002.412l10.499 7.301a.25.25 0 0 0 .392-.205V4.753Z"/>`),
		g.Group(children),
	)
}