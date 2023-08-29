package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sigma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 5H7v2.414L15.586 16L7 24.586V27h17v-2H9.414l9-9l-9-9H24V5z"/>`),
		g.Group(children),
	)
}