package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coopvoce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5C21.989 19.145 19.145 21.989 2.5 24c16.645 2.011 19.537 4.917 21.5 21.5c2.011-16.645 4.855-19.489 21.5-21.5C28.855 21.989 26.011 19.145 24 2.5Z"/>`),
		g.Group(children),
	)
}