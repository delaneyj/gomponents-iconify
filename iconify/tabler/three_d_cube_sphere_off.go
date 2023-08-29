package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDCubeSphereOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 17.6l-2-1.1V14m0-4V7.5l2-1.1m4-2.3L12 3l2 1.1m4 2.3l2 1.1V10m0 4v2m-6 3.9L12 21l-2-1.1m8-11.3l2-1.1M12 12v2.5m0 4V21m0-9l-2-1.12M6 8.6L4 7.5M3 3l18 18"/>`),
		g.Group(children),
	)
}