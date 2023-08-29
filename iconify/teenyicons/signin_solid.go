package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SigninSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 1h7v13H7v-1h6V2H7V1Zm.854 3.146l3.34 3.34l-3.327 3.603l-.734-.678L9.358 8H1V7h8.293L7.146 4.854l.708-.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}