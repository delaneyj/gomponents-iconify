package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UaOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="gold" d="M0 0h512v512H0z"/><path fill="#0057b8" d="M0 0h512v256H0z"/></g>`),
		g.Group(children),
	)
}