package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.879 1.383l1.279.213A4 4 0 0 1 15.5 5.54V8h5.32a2 2 0 0 1 1.972 2.329l-1.666 10a2 2 0 0 1-1.973 1.67H7V10.803l3.879-9.42Zm1.234 2.254L9 11.197V20h10.153l1.667-10H13.5V5.54a2 2 0 0 0-1.387-1.904ZM4 10v12H2V10h2Z"/>`),
		g.Group(children),
	)
}