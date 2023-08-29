package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#0000cd" d="M0 320.3h640V480H0z"/><path fill="#fff" d="M0 160.7h640v159.6H0z"/><path fill="#00cd00" d="M0 0h640v160.7H0z"/></g>`),
		g.Group(children),
	)
}