package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laugh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8z"/><path fill="currentColor" d="M12 18c4 0 5-4 5-4H7s1 4 5 4zm5.555-9.168l-1.109-1.664l-3 2a1.001 1.001 0 0 0 .108 1.727l4 2l.895-1.789l-2.459-1.229l1.565-1.045zm-6.557 1.23a1 1 0 0 0-.443-.894l-3-2l-1.11 1.664l1.566 1.044l-2.459 1.229l.895 1.789l4-2a.998.998 0 0 0 .551-.832z"/>`),
		g.Group(children),
	)
}