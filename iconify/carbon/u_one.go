package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M13 23H9a2 2 0 0 1-2-2V9h2v12h4V9h2v12a2 2 0 0 1-2 2z" fill="currentColor"/><path d="M22 21V9h-5v2h3v10h-3v2h8v-2h-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}