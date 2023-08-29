package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceActivatedFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm-4.5 9A2.5 2.5 0 1 1 9 13.5a2.48 2.48 0 0 1 2.5-2.5ZM16 24a8.11 8.11 0 0 1-7-4h14a8.11 8.11 0 0 1-7 4Zm4.5-8a2.5 2.5 0 1 1 2.5-2.5a2.48 2.48 0 0 1-2.5 2.5Z"/>`),
		g.Group(children),
	)
}