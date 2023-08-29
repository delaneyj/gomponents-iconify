package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.11 21.46L2.39 1.73L1.11 3l6 6H5l7 7l1.06-1.05L16.11 18H5v2h13.11l2.73 2.73l1.27-1.27M11 5h2v4.8l2.6 2.6L19 9h-4V3H9v2.8l2 2V5Z"/>`),
		g.Group(children),
	)
}