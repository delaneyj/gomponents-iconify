package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 8A1.5 1.5 0 0 1 14 9.5v1.002A1.5 1.5 0 0 1 12.5 12h-9A1.5 1.5 0 0 1 2 10.5v-1A1.5 1.5 0 0 1 3.5 8h9Zm1.058-.766l-1.673-3.507V3.72A1.233 1.233 0 0 0 10.75 3h-5.5a1.234 1.234 0 0 0-1.134.72v.007L2.442 7.234A2.49 2.49 0 0 1 3.5 7h9c.378 0 .737.084 1.058.234ZM12 10.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}