package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleExcludeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 2H9.5v7.5H2V5.25A3.25 3.25 0 0 1 5.25 2ZM2 11v4.25a3.25 3.25 0 0 0 3.25 3.25H9.5V11H2Zm16.5-1.5V5.25A3.25 3.25 0 0 0 15.25 2H11v7.5h7.5Zm-6 5.25a2.25 2.25 0 0 1 2.25-2.25h5A2.25 2.25 0 0 1 22 14.75v5A2.25 2.25 0 0 1 19.75 22h-5a2.25 2.25 0 0 1-2.25-2.25v-5Z"/>`),
		g.Group(children),
	)
}