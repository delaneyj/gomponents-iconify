package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6Zm6.5-2v11H15a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H8.5Zm-1 0H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h2.5V4Z"/>`),
		g.Group(children),
	)
}