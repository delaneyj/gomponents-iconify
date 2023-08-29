package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#0000b4" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 75.4h640v322.3H0z"/><path fill="#d90000" d="M0 157.7h640v157.7H0z"/></g>`),
		g.Group(children),
	)
}