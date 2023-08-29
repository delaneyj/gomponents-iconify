package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipelineAddThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 6a2 2 0 0 0-2 2v16a2 2 0 1 0 4 0V8a2 2 0 0 0-2-2Zm24 0a2 2 0 0 0-2 2v6.512a9.018 9.018 0 0 1 4 2.83V8a2 2 0 0 0-2-2Zm-5 8c.338 0 .672.019 1 .055V9H8v14h6a9 9 0 0 1 9-9Zm0 16.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm1-12.25V22h3.75a.75.75 0 0 1 0 1.5H24v3.75a.75.75 0 0 1-1.5 0V23.5h-3.75a.75.75 0 0 1 0-1.5h3.75v-3.75a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}