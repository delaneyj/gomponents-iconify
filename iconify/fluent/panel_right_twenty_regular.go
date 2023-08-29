package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6Zm-6.5-2v11H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h6.5Zm1 0H15a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2h-2.5V4Z"/>`),
		g.Group(children),
	)
}