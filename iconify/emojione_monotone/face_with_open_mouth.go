package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 57.5C16.836 59.5 4.5 47.164 4.5 32S16.836 4.5 32 4.5S59.5 16.836 59.5 32S47.164 59.5 32 59.5z"/><circle cx="32" cy="45.139" r="7" fill="currentColor"/><circle cx="20.248" cy="25.045" r="4.5" fill="currentColor"/><circle cx="42.75" cy="25.045" r="4.5" fill="currentColor"/>`),
		g.Group(children),
	)
}