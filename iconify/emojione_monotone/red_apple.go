package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedApple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M34.76 13.734c.801-2.69 3-6.47 8.947-6.47V4.71c-6.627 0-9.826 3.478-11.176 7.034c.31-2.519.504-5.525-.099-6.843c-1.126-2.466-4.095-3.575-6.631-2.478c-2.532 1.097-3.673 3.985-2.545 6.451c.507 1.11 2.309 2.614 4.137 3.936C18.167 8.878 3 11.569 3 29.317C3 42.747 18.781 62 25.668 62c3.125 0 4.25-1.955 6.063-1.955c1.844 0 2.438 1.951 6.436 1.951C45.059 61.996 61 42.809 61 29.317c0-18.994-17.37-20.743-26.24-15.583"/>`),
		g.Group(children),
	)
}