package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EveryDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.117 4.5h23.767v39H12.117z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.838 28.415h18.289v12.748H14.838zm.012-19.183c2.894-1.368 5.909-1.954 9.133-1.954c3.348 0 6.15.533 9.13 2.004m.005.004V22.01H14.84V9.265h.03"/><circle cx="32.914" cy="25.281" r="1.24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}