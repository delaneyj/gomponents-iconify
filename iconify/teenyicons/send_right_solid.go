package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendRightSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 14V1h1v13H1ZM9.854 3.146L14.207 7.5l-4.353 4.354l-.708-.708L12.293 8H3V7h9.293L9.146 3.854l.708-.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}