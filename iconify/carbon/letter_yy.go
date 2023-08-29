package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterYy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 9h-2l-2 7l-2-7H7l3 9v5h2v-5l3-9zm8 4l-2 7.52L19.08 13H17l3.15 9.87l-.62 2.13H17v2h2.26a2 2 0 0 0 1.91-1.42L25 13z"/>`),
		g.Group(children),
	)
}