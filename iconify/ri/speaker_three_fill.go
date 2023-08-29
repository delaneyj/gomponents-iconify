package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm8 13a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0 2a6 6 0 1 0 0-12a6 6 0 0 0 0 12ZM6 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm12 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM6 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6-5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}