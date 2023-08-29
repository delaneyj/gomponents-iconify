package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarminJr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.176V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H15.291"/><circle cx="9.5" cy="9.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.176 8.115a2.679 2.679 0 0 0-2.82-2.678a2.782 2.782 0 0 0-2.532 2.83v2.485a2.679 2.679 0 0 0 2.676 2.68h0a2.679 2.679 0 0 0 2.676-2.68H9.5M20.19 37.5h7.62M16.89 20.622h-3.118v3.134a3.999 3.999 0 0 0 3.999 4h1.042"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37.5c0-5.922 7.21-6.878 7.21-19H16.79c0 12.123 7.21 13.078 7.21 19"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.11 20.622h3.118v3.134a3.999 3.999 0 0 1-3.999 4h-1.042M24 20.965l.948 1.92l2.118.307l-1.533 1.495l.362 2.11L24 25.801l-1.895.996l.362-2.11l-1.533-1.495l2.118-.307l.948-1.92z"/>`),
		g.Group(children),
	)
}