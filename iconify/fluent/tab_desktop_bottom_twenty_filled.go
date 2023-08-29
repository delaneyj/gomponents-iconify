package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopBottomTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.503 17a2.5 2.5 0 0 0 2.5-2.5v-9a2.5 2.5 0 0 0-2.5-2.5h-9a2.5 2.5 0 0 0-2.5 2.5V13h7.5a1.5 1.5 0 0 1 1.5 1.5V17h2.5Zm-3.5 0v-2.5a.5.5 0 0 0-.5-.5h-7.5v.5a2.5 2.5 0 0 0 2.5 2.5h5.5Z"/>`),
		g.Group(children),
	)
}