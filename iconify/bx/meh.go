package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8z"/><circle cx="8.5" cy="10.5" r="1.5" fill="currentColor"/><circle cx="15.493" cy="10.493" r="1.493" fill="currentColor"/><path fill="currentColor" d="M7.974 15H16v2H7.974z"/>`),
		g.Group(children),
	)
}