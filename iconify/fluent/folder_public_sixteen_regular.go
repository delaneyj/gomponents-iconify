package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPublicSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M3 11V6.499l2.755.001a.5.5 0 0 0 .385-.18L7.235 5h4.763a1 1 0 0 1 1 1v2.09c.205.071.398.19.562.354l.438.437V6a2 2 0 0 0-2-2H7.176l-1.108-.89A.5.5 0 0 0 5.755 3H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2.877l-.436-.434A1.5 1.5 0 0 1 6.085 12H4a1 1 0 0 1-1-1zm1-7h1.579l.712.572l-.77.928L3 5.499V5a1 1 0 0 1 1-1z" fill="currentColor"/><path d="M9.856 9.859a.5.5 0 0 0-.706-.708l-2.003 1.998a.5.5 0 0 0 0 .708l2.003 1.997a.5.5 0 0 0 .706-.708L8.71 12.002h4.584l-1.146 1.144a.5.5 0 1 0 .706.708l2-1.996a.5.5 0 0 0 0-.708l-2-1.999a.5.5 0 0 0-.707.707l1.145 1.144H8.71L9.856 9.86z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}