package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipelineOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20v-8a2 2 0 0 0-2-2h-5V9h-2v6h2v-1h3v6h-1v2h6v-2Zm-1 0h-2v-7h-4v-2h5a1 1 0 0 1 1 1ZM9 9v1H6V4h1V2H1v2h1v8a2 2 0 0 0 2 2h5v1h2V9Zm0 4H4a1 1 0 0 1-1-1V4h2v7h4Z"/>`),
		g.Group(children),
	)
}