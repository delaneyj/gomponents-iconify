package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM7 4a3 3 0 1 0 1.738 5.445l1.409 1.409l.707-.707l-1.409-1.409A3 3 0 0 0 7 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}