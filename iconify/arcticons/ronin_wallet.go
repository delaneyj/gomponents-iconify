package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoninWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.497 30.887V10.941h13.255v7.433l-2.446 2.447l-10.809.094m10.809-.094l2.57 2.571v7.495"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 4.5h23.001a2.684 2.684 0 0 1 2.684 2.684v24.86a2.684 2.684 0 0 1-1.046 2.126l-11.391 8.773a2.684 2.684 0 0 1-3.257.013L10.879 34.17a2.684 2.684 0 0 1-1.064-2.14l.001-24.846A2.684 2.684 0 0 1 12.5 4.5Z"/>`),
		g.Group(children),
	)
}