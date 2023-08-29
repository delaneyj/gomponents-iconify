package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCreditCardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.83 4H22v15.17L14.83 12H20V8h-9.17l-4-4zm13.66 19.31L17.17 20H2V4.83L.69 3.51L2.1 2.1l19.8 19.8l-1.41 1.41zM9.17 12l-4-4H4v4h5.17z"/>`),
		g.Group(children),
	)
}