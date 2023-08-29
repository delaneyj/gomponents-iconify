package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageAltEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h3.5l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5zm4.302 11.987h-1.8v-1.799l4.978-4.97l1.798 1.799l-4.976 4.97zm5.823-5.817l-1.798-1.799L14.698 5l1.8 1.799l-1.373 1.371z"/>`),
		g.Group(children),
	)
}