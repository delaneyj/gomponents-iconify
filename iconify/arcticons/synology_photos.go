package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SynologyPhotos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="34.474" cy="30.047" r="9.406" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.474" cy="17.953" r="9.406" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="11.906" r="9.406" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.527" cy="17.953" r="9.406" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13.527" cy="30.047" r="9.406" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="36.094" r="9.406" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}