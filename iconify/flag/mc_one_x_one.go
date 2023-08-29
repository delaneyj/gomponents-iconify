package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func McOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#f31830" d="M0 0h512v256H0z"/><path fill="#fff" d="M0 256h512v256H0z"/></g>`),
		g.Group(children),
	)
}