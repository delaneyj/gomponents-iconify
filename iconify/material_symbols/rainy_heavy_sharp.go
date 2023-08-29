package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainyHeavySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.55 19.35L.65 5.55l1.8-.875l6.9 13.75l-1.8.925Zm4.675 0l-6.9-13.8l1.8-.9L14 18.425l-1.775.925Zm9.325-.025L14.675 5.55l1.775-.9l6.9 13.8l-1.8.875Zm-4.65.025L10 5.55l1.775-.9l6.9 13.775l-1.775.925Z"/>`),
		g.Group(children),
	)
}