package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImageBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M114 145.34a12 12 0 0 0-20 0l-15 22.48l-6.7-10.09a12 12 0 0 0-20.07.17l-38.32 59.61A12 12 0 0 0 24 236h128a12 12 0 0 0 10-18.66ZM46 212l16.5-25.7l6.55 9.81a12 12 0 0 0 20 0l15-22.48L129.58 212ZM216.49 79.51l-56-56A12 12 0 0 0 152 20H56a20 20 0 0 0-20 20v84a12 12 0 0 0 24 0V44h76v48a12 12 0 0 0 12 12h48v108a12 12 0 0 0 0 24h4a20 20 0 0 0 20-20V88a12 12 0 0 0-3.51-8.49ZM160 57l23 23h-23Z"/>`),
		g.Group(children),
	)
}