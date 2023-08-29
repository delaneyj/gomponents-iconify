package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarminGolf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.176V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H15.291"/><circle cx="9.5" cy="9.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.176 8.115a2.679 2.679 0 0 0-2.82-2.678a2.782 2.782 0 0 0-2.532 2.83v2.485a2.679 2.679 0 0 0 2.676 2.68h0a2.679 2.679 0 0 0 2.676-2.68H9.5"/><circle cx="24" cy="24.492" r="7.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.228 33.282a10.015 10.015 0 0 0 9.544 0m-3.468-5.468a1.962 1.962 0 0 0 1.322-1.322m1.477 1.827a1.962 1.962 0 0 0 .956-1.607m-4.085 3.325a1.963 1.963 0 0 0 1.756-.644M24 34.492v4.516"/>`),
		g.Group(children),
	)
}