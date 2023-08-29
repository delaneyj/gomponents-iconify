package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectSource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24 10l-1.414 1.414L26.172 15H11.899A5.014 5.014 0 0 0 8 11.101V2H6v9.101A5 5 0 0 0 6 20.9V30h2v-9.101A5.014 5.014 0 0 0 11.899 17h14.273l-3.586 3.586L24 22l6-6ZM7 19a3 3 0 1 1 3-3a3.003 3.003 0 0 1-3 3Z"/>`),
		g.Group(children),
	)
}