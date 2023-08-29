package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoldMedalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGoldMedalTwo0"><g fill="none"><path fill="#fff" d="M36 32a11.97 11.97 0 0 0-4-8.944A11.955 11.955 0 0 0 24 20a11.955 11.955 0 0 0-8 3.056A11.97 11.97 0 0 0 12 32c0 6.627 5.373 12 12 12s12-5.373 12-12Z"/><path fill="#fff" d="M16 4h16v19.056A11.955 11.955 0 0 0 24 20a11.955 11.955 0 0 0-8 3.056V4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 4h16M16 4H8v10l8 6m0-16v16M32 4h8v10l-8 6m0-16v16m-16 3.056A11.955 11.955 0 0 1 24 20c3.073 0 5.877 1.155 8 3.056m-16 0A11.97 11.97 0 0 0 12 32c0 6.627 5.373 12 12 12s12-5.373 12-12a11.97 11.97 0 0 0-4-8.944m-16 0V20m16 3.056V20"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 4h8v10l-8 6M16 4H8v10l8 6"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 32a11.97 11.97 0 0 0-4-8.944A11.955 11.955 0 0 0 24 20a11.955 11.955 0 0 0-8 3.056A11.97 11.97 0 0 0 12 32c0 6.627 5.373 12 12 12s12-5.373 12-12Z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 4h16v19.056A11.955 11.955 0 0 0 24 20a11.955 11.955 0 0 0-8 3.056V4Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 37V27l-2 1m2 9h2m-2 0h-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGoldMedalTwo0)"/>`),
		g.Group(children),
	)
}