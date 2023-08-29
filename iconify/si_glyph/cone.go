package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.287 13.025h-2.013l-.528-2h-9.57l-.528 2H1.692c-.37 0-.671.271-.671.609v.684c0 .337.301.609.671.609h14.596c.37 0 .671-.272.671-.609v-.684c-.001-.337-.302-.609-.672-.609zm-5.08-11.994h-4.41l-.759 2.921h5.931l-.762-2.921zm1.139 4.973H5.658l-.637 2.951h7.963l-.638-2.951z"/>`),
		g.Group(children),
	)
}