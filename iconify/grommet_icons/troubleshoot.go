package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Troubleshoot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 5c0-2 1-4 2-4l2 4h2l2-4c1 0 2 2 2 4c0 2.254-1 4-3 5v11c0 1 0 2-2 2s-2-1-2-2V10C2 9 1 7.254 1 5Zm18 7v6m-2 0l1 5h2l1-5h-4Zm-3-6h10h-10Zm7 0V3a2 2 0 1 0-4 0v9"/>`),
		g.Group(children),
	)
}