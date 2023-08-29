package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BwFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#00cbff" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 160h640v160H0z"/><path d="M0 186h640v108H0z"/></g>`),
		g.Group(children),
	)
}