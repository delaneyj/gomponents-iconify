package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarcodeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 13V2h1v11H0Zm5 0V2h1v11H5Zm2 0V2h1v11H7Zm3 0V2h1v11h-1Zm4 0V2h1v11h-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}