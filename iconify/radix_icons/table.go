package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 2h4.5a.5.5 0 0 1 .5.5V5H8V2ZM7 5V2H2.5a.5.5 0 0 0-.5.5V5h5ZM2 6v3h5V6H2Zm6 0h5v3H8V6Zm0 4h5v2.5a.5.5 0 0 1-.5.5H8v-3Zm-6 2.5V10h5v3H2.5a.5.5 0 0 1-.5-.5Zm-1-10A1.5 1.5 0 0 1 2.5 1h10A1.5 1.5 0 0 1 14 2.5v10a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 12.5v-10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}