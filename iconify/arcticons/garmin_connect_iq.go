package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarminConnectIq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.5 37.635V20.002m18.442 17.041a9.695 9.695 0 1 1 3.087-3.088m-4.92-1.83l6.391 6.391"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.176V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H15.291"/><circle cx="9.5" cy="9.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.176 8.115a2.679 2.679 0 0 0-2.82-2.678a2.782 2.782 0 0 0-2.532 2.83v2.485a2.679 2.679 0 0 0 2.676 2.68h0a2.679 2.679 0 0 0 2.676-2.68H9.5"/>`),
		g.Group(children),
	)
}