package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M205.688 188.671L0 185.711l308.532 96.185l-31.815-71.769l58.451 2.22l-22.937 223.445h168.694l9.618-88.046h-54.751l13.318-130.22h54.751L512 129.48h-54.751l5.179-53.272H347.746l-5.919 55.491H183.491z"/>`),
		g.Group(children),
	)
}