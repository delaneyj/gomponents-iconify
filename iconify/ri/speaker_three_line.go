package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5v14h14V5H5ZM4 3h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm3 5a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm10 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0 10a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM7 18a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm5-3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 2a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm0-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}