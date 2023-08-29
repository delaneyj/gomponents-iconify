package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowMinimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1792 1056v192q0 66-47 113t-113 47H160q-66 0-113-47T0 1248v-192q0-66 47-113t113-47h1472q66 0 113 47t47 113z"/>`),
		g.Group(children),
	)
}