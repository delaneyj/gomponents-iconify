package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2.281L8 6.844V16l8-4.563zm0 9.156L24 16V6.844zM24 16l-8 4.563v9.156l8-4.563zm-8 4.563L8 16v9.156z"/>`),
		g.Group(children),
	)
}