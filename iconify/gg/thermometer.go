package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.95 5.636a1 1 0 1 1 1.414 1.414l-7.071 7.071a1 1 0 1 1-1.414-1.414l7.07-7.07Z"/><path fill-rule="evenodd" d="M7.828 17.586a5.002 5.002 0 0 0 6.293-.636l7.071-7.071a5 5 0 1 0-7.07-7.071L7.05 9.878a5.002 5.002 0 0 0-.636 6.294l-3.606 3.606a1 1 0 1 0 1.414 1.415l3.606-3.607Zm4.88-2.05l7.07-7.071a3 3 0 1 0-4.242-4.243l-7.071 7.071a3 3 0 1 0 4.242 4.243Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}