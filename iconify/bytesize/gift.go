package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14v16h24V14M2 9v5h28V9Zm14 0s-2-9-8-6s8 6 8 6s2-9 8-6s-8 6-8 6m0 0v21"/>`),
		g.Group(children),
	)
}