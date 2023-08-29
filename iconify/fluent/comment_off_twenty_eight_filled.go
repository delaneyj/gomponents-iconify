package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentOffTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l.848.849A3.738 3.738 0 0 0 2 6.75v10.5A3.75 3.75 0 0 0 5.75 21H7v3.296c0 1.427 1.616 2.254 2.774 1.419L16.309 21h3.63l4.78 4.78a.75.75 0 0 0 1.061-1.06L3.28 2.22ZM6.182 3l17.65 17.65A3.75 3.75 0 0 0 26 17.25V6.75A3.75 3.75 0 0 0 22.25 3H6.182Z"/>`),
		g.Group(children),
	)
}