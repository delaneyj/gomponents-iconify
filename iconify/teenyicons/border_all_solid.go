package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAllSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v13H1V1Zm1 1v5h5V2H2Zm6 0v5h5V2H8Zm5 6H8v5h5V8Zm-6 5V8H2v5h5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}