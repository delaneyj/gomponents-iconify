package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a5 5 0 0 0-5 5c0 5 5 11 5 11s5-6 5-11a5 5 0 0 0-5-5zm0 8.063a3.063 3.063 0 1 1 0-6.126a3.063 3.063 0 0 1 0 6.126zM6.063 5a1.938 1.938 0 1 1 3.876 0a1.938 1.938 0 0 1-3.876 0z"/>`),
		g.Group(children),
	)
}