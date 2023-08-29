package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3v9.586l4-4L18.414 10L12 16.414L5.586 10L7 8.586l4 4V3h2ZM3 18h18v2H3v-2Z"/>`),
		g.Group(children),
	)
}