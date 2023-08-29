package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd" stroke-width="1pt"><path fill="#0000b4" d="M0 0h512v512H0z"/><path fill="#fff" d="M0 80.5h512v343.7H0z"/><path fill="#d90000" d="M0 168.2h512v168.2H0z"/></g>`),
		g.Group(children),
	)
}