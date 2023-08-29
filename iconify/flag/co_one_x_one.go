package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#ffe800" d="M0 0h512v512H0z"/><path fill="#00148e" d="M0 256h512v256H0z"/><path fill="#da0010" d="M0 384h512v128H0z"/></g>`),
		g.Group(children),
	)
}