package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWinkFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2ZM8 13h6v2H8Zm8 11a8 8 0 0 1-6.85-3.89l1.71-1a6 6 0 0 0 10.28 0l1.71 1A8 8 0 0 1 16 24Zm4.5-8a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}