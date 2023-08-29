package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 15V0h1v15H1Zm4 0V0h1v15H5Zm4 0V0h1v15H9Zm4 0V0h1v15h-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}