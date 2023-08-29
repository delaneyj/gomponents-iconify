package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 3.5V6h1.5A1.5 1.5 0 0 1 14 7.5v6a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-6A1.5 1.5 0 0 1 2.5 6H4V3.5a3.5 3.5 0 1 1 7 0Zm-6 0a2.5 2.5 0 0 1 5 0V6H5V3.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}