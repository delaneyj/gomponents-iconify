package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.414 1.586A2 2 0 0 0 10 1H6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V3a2 2 0 0 0-.586-1.414ZM9.853 12.854a.5.5 0 0 1-.354.146h-3a.5.5 0 1 1 0-1h3a.5.5 0 0 1 .354.854Zm0-2a.5.5 0 0 1-.354.146h-3a.5.5 0 1 1 0-1h3a.5.5 0 0 1 .354.854Zm0-6A.5.5 0 0 1 9.499 5h-3a.498.498 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .354.854Z"/>`),
		g.Group(children),
	)
}