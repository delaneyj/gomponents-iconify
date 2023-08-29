package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpticalAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M4.5 9a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0ZM8 6.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/><path d="M2 14.5a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-3.05a2.5 2.5 0 0 0 0-4.9V4.5a.5.5 0 0 0-.146-.354l-2-2A.5.5 0 0 0 11.5 2h-7a.5.5 0 0 0-.354.146l-2 2A.5.5 0 0 0 2 4.5v2.05a2.5 2.5 0 0 0 0 4.9v3.05Zm1-.5v-3a.5.5 0 0 0-.5-.5a1.5 1.5 0 1 1 0-3A.5.5 0 0 0 3 7V4.707L4.707 3h6.586L13 4.707V7a.5.5 0 0 0 .5.5a1.5 1.5 0 0 1 0 3a.5.5 0 0 0-.5.5v3H3Z"/></g>`),
		g.Group(children),
	)
}