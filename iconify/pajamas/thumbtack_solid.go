package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbtackSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.28 1.22a.75.75 0 0 0-1.26.7L6.69 5.25H4.206c-1.114 0-1.671 1.346-.884 2.134l1.911 1.911l-3.72 4.135A2 2 0 0 0 1 14.768V15h.233a2 2 0 0 0 1.337-.513l4.135-3.721l1.911 1.91c.788.788 2.134.23 2.134-.883V9.31l3.33-3.33a.75.75 0 0 0 .7-1.261l-.603-.604l-2.293-2.293l-.604-.603Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}