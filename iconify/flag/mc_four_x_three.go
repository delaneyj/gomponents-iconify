package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func McFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#f31830" d="M0 0h640v240H0z"/><path fill="#fff" d="M0 240h640v240H0z"/></g>`),
		g.Group(children),
	)
}