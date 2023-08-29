package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4a2 2 0 0 0-2 2v1h15v8.5a.5.5 0 0 0 1 0V6a2 2 0 0 0-2-2H4Zm6 4H2v6a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V8Zm-5.5 2.5A.5.5 0 0 1 5 10h2a.5.5 0 0 1 0 1H5a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}