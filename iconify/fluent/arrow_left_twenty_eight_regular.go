package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.735 4.21a.75.75 0 0 1 1.04 1.08L5.5 13.25h18.753a.75.75 0 0 1 0 1.5H5.501l8.274 7.959a.75.75 0 0 1-1.04 1.08L3.307 14.72a1 1 0 0 1 0-1.441l9.428-9.07Z"/>`),
		g.Group(children),
	)
}