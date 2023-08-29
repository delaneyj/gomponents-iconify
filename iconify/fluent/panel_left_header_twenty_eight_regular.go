package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.004 7.75A3.75 3.75 0 0 1 5.754 4H22.25A3.75 3.75 0 0 1 26 7.75v11.5A3.75 3.75 0 0 1 22.25 23H5.755a3.75 3.75 0 0 1-3.75-3.75V7.75Zm3.75-2.25a2.25 2.25 0 0 0-2.25 2.25v11.5a2.25 2.25 0 0 0 2.25 2.25H9v-16H5.754Zm4.746 6v10h11.75a2.25 2.25 0 0 0 2.25-2.25V11.5h-14Zm14-1.5V7.75a2.25 2.25 0 0 0-2.25-2.25H10.5V10h14Z"/>`),
		g.Group(children),
	)
}