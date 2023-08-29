package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.754 4a3.75 3.75 0 0 0-3.75 3.75v11.5A3.75 3.75 0 0 0 5.754 23H22.25A3.75 3.75 0 0 0 26 19.25V7.75A3.75 3.75 0 0 0 22.25 4H5.755ZM24.5 19.25a2.25 2.25 0 0 1-2.25 2.25H10.5v-10h14v7.75Zm0-9.25h-14V5.5h11.75a2.25 2.25 0 0 1 2.25 2.25V10Z"/>`),
		g.Group(children),
	)
}