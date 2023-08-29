package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Insights(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 8h4v12H5V10h4V4h6v4Zm-2-2h-2v12h2V6Zm2 4v8h2v-8h-2Zm-6 2v6H7v-6h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}