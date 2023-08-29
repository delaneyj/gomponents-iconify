package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 2.75a.75.75 0 0 0-1.5 0v18.5a.75.75 0 0 0 1.5 0V2.75ZM4.25 4A2.25 2.25 0 0 0 2 6.25v11.5A2.25 2.25 0 0 0 4.25 20H10v-1.5H4.25a.75.75 0 0 1-.75-.75V6.25a.75.75 0 0 1 .75-.75H10V4H4.25Zm15 14.5H13.5V20h5.75a2.25 2.25 0 0 0 2.25-2.25V6.25A2.25 2.25 0 0 0 19.25 4H13.5v1.5h5.75a.75.75 0 0 1 .75.75v11.5a.75.75 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}