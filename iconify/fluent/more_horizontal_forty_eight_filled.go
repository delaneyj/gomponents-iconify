package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHorizontalFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 24a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm11.5 0a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm8 3.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Z"/>`),
		g.Group(children),
	)
}