package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThisSideUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 28h28v2H2zM7 5.828V24h2V5.828l3.586 3.586L14 8L8 2L2 8l1.414 1.414L7 5.828zm16 0V24h2V5.828l3.586 3.586L30 8l-6-6l-6 6l1.414 1.414L23 5.828z"/>`),
		g.Group(children),
	)
}