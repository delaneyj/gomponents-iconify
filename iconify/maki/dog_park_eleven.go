package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogParkEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7.5 1s-.346.037-.5.5l-.5 2l2 1.5H10c1 0 1-1 1-1L9.5 2.5C9 2 8.5 2 8 2v-.5S8 1 7.5 1zm-5 1s-.353-.007-.723.178S1 2.833 1 3.5v2C1 6.5 1 7 .5 7c0 0-.5 0-.5.5v2s0 .5.5.5s.5-.5.5-.5V8c.354 0 .69-.137 1-.299V9.5s0 .5.5.5s.5-.5.5-.5V7h3l.664 1.992C7 10 7.5 10 7.5 10H8s.5 0 .5-.5S8 9 8 9V6.5c0-.89-.632-1.255-1-1.5L5.498 4H2v-.5c0-.333.092-.362.223-.428C2.353 3.007 2.5 3 2.5 3c.676.01.676-1.01 0-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}