package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentLinkTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 3.5a2.25 2.25 0 0 0 0 4.5h.5a.75.75 0 0 1 0 1.5h-.5a3.75 3.75 0 1 1 0-7.5h.5a.75.75 0 0 1 0 1.5h-.5ZM17 5.75a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5a.75.75 0 0 1-.75-.75ZM23.25 8a2.25 2.25 0 0 0 0-4.5h-.5a.75.75 0 0 1 0-1.5h.5a3.75 3.75 0 1 1 0 7.5h-.5a.75.75 0 0 1 0-1.5h.5ZM13 5.75c0-1.025.325-1.974.877-2.75H5.75A3.75 3.75 0 0 0 2 6.75v10.5A3.75 3.75 0 0 0 5.75 21H7v3.296c0 1.427 1.616 2.254 2.774 1.419L16.309 21h5.941A3.75 3.75 0 0 0 26 17.25V9.623a4.729 4.729 0 0 1-2.75.877h-5.5A4.75 4.75 0 0 1 13 5.75Z"/>`),
		g.Group(children),
	)
}