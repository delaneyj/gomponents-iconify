package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#fff" d="M512 512H0V0h512z"/><path fill="#388d00" d="M512 512H0V341.3h512z"/><path fill="#d43516" d="M512 170.8H0V.1h512z"/></g>`),
		g.Group(children),
	)
}