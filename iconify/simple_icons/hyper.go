package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hyper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.565 17.91H24v1.964H13.565zm-3.201-5.09l-9.187 8.003l2.86-7.004L0 11.179l9.187-8.002l-3.11 7.451z"/>`),
		g.Group(children),
	)
}