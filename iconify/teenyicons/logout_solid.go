package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoutSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h7v1H2v11h6v1H1V1Zm9.854 3.146l3.34 3.34l-3.327 3.603l-.734-.678L12.358 8H4V7h8.293l-2.147-2.146l.708-.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}