package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DualScreenTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h5.5V4H4Zm2 9.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1-.5-.5ZM10.5 4v12H16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-5.5Zm2 9h1a.5.5 0 0 1 0 1h-1a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}