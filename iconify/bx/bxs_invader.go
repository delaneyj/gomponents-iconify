package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsInvader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M6 3h2v2H6z" fill="currentColor"/><path d="M8 19h3v2H8z" fill="currentColor"/><path d="M16 3h2v2h-2z" fill="currentColor"/><path d="M13 19h3v2h-3z" fill="currentColor"/><path d="M20 11V9h-2V7h-2V5h-2v2h-4V5H8v2H6v2H4v2H2v8h2v-4h2v4h2v-3h8v3h2v-4h2v4h2v-8zm-10 1H8V9h2zm6 0h-2V9h2z" fill="currentColor"/>`),
		g.Group(children),
	)
}