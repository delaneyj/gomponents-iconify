package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineDoneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.77 5.03l1.4 1.4L8.43 19.17l-5.6-5.6l1.4-1.4l4.2 4.2L19.77 5.03m0-2.83L8.43 13.54l-4.2-4.2L0 13.57L8.43 22L24 6.43L19.77 2.2z"/>`),
		g.Group(children),
	)
}