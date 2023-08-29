package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleBathroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h14a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m7 2a4 4 0 0 0-4 4h3.26l-.41-2.77L12.9 8H16a4 4 0 0 0-4-4m-7 6v10h14V10H5Z"/>`),
		g.Group(children),
	)
}