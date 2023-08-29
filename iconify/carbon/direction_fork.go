package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 13V7.414l9 9V28h2V16.414l9-9V13h2V4h-9v2h5.586L16 14.586L7.414 6H13V4H4v9h2z"/>`),
		g.Group(children),
	)
}