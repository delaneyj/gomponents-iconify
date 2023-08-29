package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.985 4.044c0 1.671-1.336 4.863-2.983 4.863c-1.648 0-2.986-3.192-2.986-4.863c0-1.669 1.338-3.023 2.986-3.023c1.647 0 2.983 1.354 2.983 3.023zm-1.047 8.938h5.823c-.447-1.584-1.593-3.953-4.058-3.953c-.854 1.159-3.692 1.456-3.692 1.456s-2.903-.286-3.759-1.43c-4.096 0-4.221 5.923-4.221 5.923h9.906v-1.996h.001z"/><path d="M11 14h4.937v.972H11z"/></g>`),
		g.Group(children),
	)
}