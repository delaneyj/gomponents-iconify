package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BwOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#00cbff" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 192h512v128H0z"/><path d="M0 212.7h512V299H0z"/></g>`),
		g.Group(children),
	)
}