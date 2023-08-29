package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 16a4 4 0 1 1 6.646 3H9.354A3.988 3.988 0 0 1 8 16zm.97 5H21a1 1 0 1 0 0-2h-3.803a6 6 0 1 0-10.394 0H3a1 1 0 1 0 0 2h5.97zM19.071 8.929a1 1 0 0 1 0 1.414l-.707.707a1 1 0 0 1-1.414-1.414l.707-.707a1 1 0 0 1 1.414 0zM4 17a1 1 0 1 0 0-2H3a1 1 0 1 0 0 2h1zm18-1a1 1 0 0 1-1 1h-1a1 1 0 1 1 0-2h1a1 1 0 0 1 1 1zM5.636 11.05A1 1 0 0 0 7.05 9.636l-.707-.707a1 1 0 1 0-1.414 1.414l.707.707zM13 5.586V3a1 1 0 1 0-2 0v2.586l-.293-.293a1 1 0 0 0-1.414 1.414l1.999 1.999l.01.01a.997.997 0 0 0 1.406-.01l2-1.999a1 1 0 0 0-1.415-1.414L13 5.586z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}