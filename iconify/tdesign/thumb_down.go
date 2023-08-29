package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.879 22.617l1.279-.213A4 4 0 0 0 15.5 18.46V16h5.32a2 2 0 0 0 1.972-2.329l-1.666-10a2 2 0 0 0-1.973-1.67H7v11.197l3.879 9.42Zm1.234-2.254L9 12.803V4h10.153l1.667 10H13.5v4.459a2 2 0 0 1-1.387 1.904ZM4 14V2H2v12h2Z"/>`),
		g.Group(children),
	)
}