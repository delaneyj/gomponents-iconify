package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedDoubleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M0 0h15v6h-1V1H1v5H0V0Z"/><path fill="currentColor" d="M6 6H2V5h4v1Zm-6 9h1v-4h13v4h1V7H0v8Zm9-9h4V5H9v1Z"/>`),
		g.Group(children),
	)
}