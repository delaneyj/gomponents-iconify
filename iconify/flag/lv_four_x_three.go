package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LvFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#fff" d="M0 0h640v480H0z"/><path fill="#981e32" d="M0 0h640v192H0zm0 288h640v192H0z"/></g>`),
		g.Group(children),
	)
}