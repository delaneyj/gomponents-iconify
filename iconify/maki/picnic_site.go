package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PicnicSite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 3c-.554 0-1 .446-1 1s.446 1 1 1h1.297l-.645 2H2.5c-.554 0-1 .446-1 1s.446 1 1 1h1.51l-.969 3.01a.75.75 0 1 0 1.426.465l.002-.004L5.586 9h3.828l1.117 3.47a.75.75 0 1 0 1.428-.46L10.99 9h1.51c.554 0 1-.446 1-1s-.446-1-1-1h-2.152l-.645-2H11c.554 0 1-.446 1-1s-.446-1-1-1H4zm2.873 2h1.254l.644 2H6.23l.644-2z"/>`),
		g.Group(children),
	)
}