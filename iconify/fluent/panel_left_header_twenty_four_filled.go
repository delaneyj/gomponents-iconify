package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftHeaderTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 7.25A3.25 3.25 0 0 1 5.25 4h13.5A3.25 3.25 0 0 1 22 7.25v9.5A3.25 3.25 0 0 1 18.75 20H5.25A3.25 3.25 0 0 1 2 16.75v-9.5Zm18.5 0a1.75 1.75 0 0 0-1.75-1.75H9.5V9h11V7.25Zm0 3.25h-11v8h9.25a1.75 1.75 0 0 0 1.75-1.75V10.5Z"/>`),
		g.Group(children),
	)
}