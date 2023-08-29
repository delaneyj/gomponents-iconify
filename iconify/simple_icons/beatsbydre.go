package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beatsbydre(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.01 15.6a3.6 3.6 0 1 0-3.6-3.6a3.6 3.6 0 0 0 3.6 3.6zm0-15.598a11.998 11.998 0 0 0-3.6.552V7.2a6 6 0 1 1-2.4 4.8V1.603a11.998 11.998 0 1 0 6-1.601z"/>`),
		g.Group(children),
	)
}