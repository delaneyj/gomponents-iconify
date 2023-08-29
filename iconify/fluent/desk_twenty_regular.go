package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V8h7v7.5a.5.5 0 0 0 1 0V6a2 2 0 0 0-2-2H4ZM3 8h6v6a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V8Zm0-1V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v1H3Zm2 2.5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1H5Z"/>`),
		g.Group(children),
	)
}