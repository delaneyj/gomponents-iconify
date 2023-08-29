package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 6c-2.745 0-5 2.255-5 5v10c0 2.745 2.255 5 5 5h16c2.745 0 5-2.255 5-5V11c0-2.745-2.255-5-5-5H8zm0 2h16c1.655 0 3 1.345 3 3v10c0 1.655-1.345 3-3 3H8c-1.655 0-3-1.345-3-3V11c0-1.655 1.345-3 3-3z"/>`),
		g.Group(children),
	)
}