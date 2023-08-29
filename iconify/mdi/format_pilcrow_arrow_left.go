package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatPilcrowArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17v-3l-4 4l4 4v-3h12v-2m-10-7v5h2V4h2v11h2V4h2V2h-8a4 4 0 0 0-4 4a4 4 0 0 0 4 4Z"/>`),
		g.Group(children),
	)
}