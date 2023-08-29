package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsMessageSquareDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6zm-5 10.5A1.5 1.5 0 0 1 9.5 14c-.086 0-.168-.011-.25-.025c-.083.01-.164.025-.25.025a2 2 0 1 1 2-2c0 .085-.015.167-.025.25c.013.082.025.164.025.25zm4 1.5c-.086 0-.167-.015-.25-.025a1.471 1.471 0 0 1-.25.025a1.5 1.5 0 0 1-1.5-1.5c0-.085.012-.168.025-.25c-.01-.083-.025-.164-.025-.25a2 2 0 1 1 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}