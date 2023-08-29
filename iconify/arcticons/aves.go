package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.692 33.92l-8.789-9.46v16.103a1.106 1.106 0 0 0 1.918.753Zm-.732-18.132l8.789 9.46V6.329l-8.789 9.459z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.904 24.46L5.798 8.201a1.107 1.107 0 0 1 .81-1.859H19.22a2.213 2.213 0 0 1 1.621.707L37.75 25.248l-8.057 8.672m12.644-23.004L37.749 6.33v5.532h4.197a.553.553 0 0 0 .39-.945Z"/>`),
		g.Group(children),
	)
}