package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 7.75a1.75 1.75 0 1 1 0-3.5a1.75 1.75 0 0 1 0 3.5Zm0 6a1.75 1.75 0 1 1 0-3.5a1.75 1.75 0 0 1 0 3.5ZM10.25 18a1.75 1.75 0 1 0 3.5 0a1.75 1.75 0 0 0-3.5 0Z"/>`),
		g.Group(children),
	)
}