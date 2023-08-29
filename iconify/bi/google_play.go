package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.222 9.374c1.037-.61 1.037-2.137 0-2.748L11.528 5.04L8.32 8l3.207 2.96l2.694-1.586Zm-3.595 2.116L7.583 8.68L1.03 14.73c.201 1.029 1.36 1.61 2.303 1.055l7.294-4.295ZM1 13.396V2.603L6.846 8L1 13.396ZM1.03 1.27l6.553 6.05l3.044-2.81L3.333.215C2.39-.341 1.231.24 1.03 1.27Z"/>`),
		g.Group(children),
	)
}