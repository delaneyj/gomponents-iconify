package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrinTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3 3.688a1.815 1.815 0 0 1 1.788 2.115a.312.312 0 0 1-.616 0c-.096-.573-.589-.833-1.171-.833s-1.074.26-1.171.833a.312.312 0 0 1-.616 0a1.815 1.815 0 0 1 1.788-2.115zm-6 0a1.815 1.815 0 0 1 1.788 2.115a.312.312 0 0 1-.616 0c-.096-.573-.589-.833-1.171-.833s-1.074.26-1.171.833a.312.312 0 0 1-.616 0a1.815 1.815 0 0 1 1.788-2.115zM3 9h3v3.873A4.017 4.017 0 0 1 3 9zm4 4V9h2v4H7zm3-.127V9h3a4.017 4.017 0 0 1-3 3.873z"/>`),
		g.Group(children),
	)
}