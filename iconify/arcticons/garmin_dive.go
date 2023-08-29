package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarminDive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 15.176V40.5a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2H15.291"/><circle cx="9.5" cy="9.433" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.176 8.115a2.679 2.679 0 0 0-2.82-2.678a2.782 2.782 0 0 0-2.532 2.83v2.485a2.679 2.679 0 0 0 2.676 2.68h0a2.679 2.679 0 0 0 2.676-2.68H9.5"/><circle cx="24.365" cy="34.883" r="3.244" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.387 34.883h-2.778m-6.489 0h-1.155a7.178 7.178 0 0 0-7.178 7.178v.446M24.365 25.88h0a3.568 3.568 0 0 0 3.568 3.569h3.71a3.569 3.569 0 0 0 3.57-3.568v-1.825a5.556 5.556 0 0 0-5.556-5.556H19.073a5.556 5.556 0 0 0-5.556 5.556v1.824a3.569 3.569 0 0 0 3.569 3.569h3.71a3.569 3.569 0 0 0 3.569-3.569h0"/>`),
		g.Group(children),
	)
}