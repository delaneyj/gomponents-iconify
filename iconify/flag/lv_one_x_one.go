package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LvOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#fff" d="M0 0h512v512H0z"/><path fill="#981e32" d="M0 0h512v204.8H0zm0 307.2h512V512H0z"/></g>`),
		g.Group(children),
	)
}