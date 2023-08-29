package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.985 3.044c0 1.671-1.336 4.863-2.983 4.863c-1.648 0-2.986-3.192-2.986-4.863c0-1.669 1.338-3.023 2.986-3.023c1.647 0 2.983 1.354 2.983 3.023zm4.973 9.977h-1.982v-1.983h-.972v1.983h-1.983v.971h1.983v1.983h.972v-1.983h1.982v-.971z"/><path d="M14.861 9.938c-.663-1.037-1.666-1.908-3.158-1.908c-.854 1.159-3.692 1.456-3.692 1.456S5.108 9.2 4.252 8.056c-4.096 0-4.221 5.923-4.221 5.923h9.906v-1.996h2.031V9.938h2.893z"/></g>`),
		g.Group(children),
	)
}