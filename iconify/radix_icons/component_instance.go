package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentInstance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.146 1.49a.5.5 0 0 1 .708 0l5.657 5.656a.5.5 0 0 1 0 .708L7.854 13.51a.5.5 0 0 1-.708 0L1.49 7.854a.5.5 0 0 1 0-.708L7.146 1.49ZM7.5 2.55L2.55 7.5l4.95 4.95l4.95-4.95L7.5 2.55Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}