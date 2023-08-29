package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1 15V0H0v15h1Zm14 0V0h-1v15h1ZM10 1H5v13h5V1Z"/>`),
		g.Group(children),
	)
}