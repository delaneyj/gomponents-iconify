package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cricket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zm2.88 5.68l-2-2a.5.5 0 0 0-.4-.18H3.75a.74.74 0 0 0-.75.64l-1 7.7a.483.483 0 0 0 0 .14a.5.5 0 0 0 .5.5a.49.49 0 0 0 .5-.34l1.2-3.89l.26-.83l.4.44L6 10.6v2.9a.5.5 0 0 0 1 0v-3a.48.48 0 0 0-.08-.22L5.48 8.5l1-2.5h1.71l2 1.84a.49.49 0 0 0 .37.17a.5.5 0 0 0 .44-.51a.49.49 0 0 0-.12-.32zM14 11.27a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5zm-3.24-1.53V9a.25.25 0 1 0-.5 0v.74a.49.49 0 0 0-.25.42v3.34a.5.5 0 0 0 1 0v-3.34a.49.49 0 0 0-.25-.42z"/>`),
		g.Group(children),
	)
}