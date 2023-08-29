package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RenameSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 2a.5.5 0 0 0 0 1h1v10h-1a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-1V3h1a.5.5 0 0 0 0-1h-3ZM4 4h2.5v8H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Zm8 8H9.5V4H12a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}