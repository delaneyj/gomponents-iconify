package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m5.618 9.382l1.255-2.51l2.509-1.254l-1.255 2.51l-2.509 1.254Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm10.947-2.776a.5.5 0 0 0-.67-.671l-4 2a.5.5 0 0 0-.224.223l-2 4a.5.5 0 0 0 .67.671l4-2a.5.5 0 0 0 .224-.223l2-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}