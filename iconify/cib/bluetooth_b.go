package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18.28 16.251l5.789-6.458L14.945 0v12.896L9.563 7.514L7.6 9.477l6.754 6.775L7.6 23.027l1.963 1.963l5.382-5.382l.17 12.393l9.284-9.29zm2.554-6.437l-3.124 3.124l-.021-6.268zm-3.124 9.751l3.124 3.124l-3.145 3.145z"/>`),
		g.Group(children),
	)
}