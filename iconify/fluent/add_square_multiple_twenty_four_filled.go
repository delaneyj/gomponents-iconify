package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSquareMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 5.25A2.25 2.25 0 0 1 5.25 3h11a2.25 2.25 0 0 1 2.25 2.25v11a2.25 2.25 0 0 1-2.25 2.25h-11A2.25 2.25 0 0 1 3 16.25v-11Zm8.5 2a.75.75 0 0 0-1.5 0V10H7.25a.75.75 0 0 0 0 1.5H10v2.75a.75.75 0 0 0 1.5 0V11.5h2.75a.75.75 0 0 0 0-1.5H11.5V7.25ZM7.75 21a2.25 2.25 0 0 1-2.122-1.5H16.25a3.25 3.25 0 0 0 3.25-3.25V5.628A2.25 2.25 0 0 1 21 7.75v8.5A4.75 4.75 0 0 1 16.25 21h-8.5Z"/>`),
		g.Group(children),
	)
}