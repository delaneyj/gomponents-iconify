package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m5 1.5l-4.5 8l2 4M5 1.5l2.5 4l-5 8M5 1.5h5l4.5 8M5 1.5l5 8h4.5m-12 4l2.5-4h9.5m-12 4h10l2-4"/>`),
		g.Group(children),
	)
}