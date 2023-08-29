package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laugh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm-6.447 9.105l2.459-1.229l-1.567-1.044l1.109-1.664l3 2a1 1 0 0 1-.108 1.727l-4 2l-.893-1.79zM12 18c-4 0-5-4-5-4h10s-1 4-5 4zm5.553-5.105l-4-2a1 1 0 0 1-.108-1.727l3-2l1.109 1.664l-1.566 1.044l2.459 1.229l-.894 1.79z"/>`),
		g.Group(children),
	)
}