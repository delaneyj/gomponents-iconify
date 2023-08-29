package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.535 12.832L12 3L4.03 20.275c-.07.2-.017.424.135.572c.15.148.374.193.57.116L12 18.5c1.38.468 2.416.82 3.107 1.053M19 16v3m0 3v.01"/>`),
		g.Group(children),
	)
}