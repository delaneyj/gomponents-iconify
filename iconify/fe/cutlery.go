package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cutlery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCutlery0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCutlery1" fill="currentColor"><path id="feCutlery2" d="M9 13v8a1 1 0 0 1-2 0v-8a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2v6h2V2a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2Zm9-9v17a1 1 0 0 1-2 0v-6h-1a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2Z"/></g></g>`),
		g.Group(children),
	)
}