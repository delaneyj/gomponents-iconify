package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 18h4v2H2zm24 0h4v2h-4zm-2 4v6H8v-6H6v6a2.006 2.006 0 0 0 2 2h16a2.006 2.006 0 0 0 2-2v-6zM8 16V4h8v6a2.006 2.006 0 0 0 2 2h6v4h2v-6a.91.91 0 0 0-.3-.7l-7-7A.909.909 0 0 0 18 2H8a2.006 2.006 0 0 0-2 2v12zM18 4.4l5.6 5.6H18zM10 18h4v2h-4zm8 0h4v2h-4z"/>`),
		g.Group(children),
	)
}