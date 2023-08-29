package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2zm0 1a5 5 0 1 0 0 10A5 5 0 0 0 8 3zM5.5 7.5h5.002a.5.5 0 0 1 .09.992l-.09.008H5.5a.5.5 0 0 1-.09-.992L5.5 7.5h5.002H5.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}