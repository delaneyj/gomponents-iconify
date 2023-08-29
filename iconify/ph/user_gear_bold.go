package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserGearBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152.5 156.53a72 72 0 1 0-89 0a124 124 0 0 0-48.69 35.75a12 12 0 0 0 18.38 15.44C52.54 184.69 79.11 172 108 172c37 0 61.12 19.42 74.81 35.72a12 12 0 1 0 18.38-15.44a124 124 0 0 0-48.69-35.75ZM60 100a48 48 0 1 1 48 48a48.05 48.05 0 0 1-48-48Zm182.73 41.23l-3.31 1.08l2 2.82a12 12 0 1 1-19.41 14.1l-2.01-2.82l-2 2.82a12 12 0 1 1-19.41-14.1l2-2.82l-3.31-1.08a12 12 0 0 1 7.42-22.82l3.31 1.07V116a12 12 0 0 1 24 0v3.48l3.31-1.07a12 12 0 0 1 7.42 22.82Z"/>`),
		g.Group(children),
	)
}