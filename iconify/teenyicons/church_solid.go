package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChurchSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 2V0h1v2h2v1H8v2.21l6.748 3.856l-.496.868L13 9.22V14h2v1h-5v-5H5v5H0v-1h2V9.219l-1.252.715l-.496-.868L7 5.21V3H5V2h2Z"/><path fill="currentColor" d="M6 15h3v-4H6v4Z"/>`),
		g.Group(children),
	)
}