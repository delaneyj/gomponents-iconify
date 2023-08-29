package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 44H48a20 20 0 0 0-20 20v112a20 20 0 0 0 20 20h76v24H96a4 4 0 0 0 0 8h64a4 4 0 0 0 0-8h-28v-24h76a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20ZM48 52h160a12 12 0 0 1 12 12v84H36V64a12 12 0 0 1 12-12Zm160 136H48a12 12 0 0 1-12-12v-20h184v20a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}