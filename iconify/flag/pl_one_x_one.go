package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#fff" d="M512 512H0V0h512z"/><path fill="#dc143c" d="M512 512H0V256h512z"/></g>`),
		g.Group(children),
	)
}